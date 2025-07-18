package controllers

import (
	"encoding/json"
	"errors"
	"musicfy/internal/auth/domain"
	"musicfy/internal/auth/domain/entities"
	"musicfy/internal/auth/domain/usecases"
	"musicfy/internal/auth/presentation/dtos"
	"musicfy/internal/shared"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// AuthController handles authentication-related HTTP requests
type AuthController struct {
	authUseCase *usecases.AuthUseCase
	validate    *validator.Validate
}

// NewAuthController creates a new auth controller
func NewAuthController(authUseCase *usecases.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
		validate:    validator.New(),
	}
}

// Register handles user registration requests
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	// Parse and validate request body
	var req dtos.RegisterRequest
	if err := c.decodeAndValidateRequest(w, r, &req); err != nil {
		return
	}

	// Convert age string to int
	age, err := strconv.Atoi(req.Age)
	if err != nil {
		shared.Error(w, http.StatusBadRequest, "Invalid age format", err.Error())
		return
	}

	// Register user through use case
	if err := c.authUseCase.RegisterUser(
		req.FirstName,
		req.LastName,
		req.Username,
		req.Email,
		req.Password,
		age,
	); err != nil {
		c.handleUseCaseError(w, err)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	shared.Success(w, "User created successfully", nil)
}

// Login handles user login requests
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// Parse and validate request body
	var req dtos.LoginRequest
	if err := c.decodeAndValidateRequest(w, r, &req); err != nil {
		return
	}

	// Authenticate user through use case
	token, err := c.authUseCase.LoginUser(req.UsernameOrEmail, req.Password)
	if err != nil {
		c.handleUseCaseError(w, err)
		return
	}

	// Return success response with token
	w.WriteHeader(http.StatusOK)
	shared.Success(w, "Login successful", dtos.LoginResponse{
		Token: token,
	})
}

// GetProfile retrieves the profile of the authenticated user
func (c *AuthController) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID, err := c.getUserIDFromContext(r)
	if err != nil {
		shared.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Get user from use case
	user, err := c.authUseCase.GetUserByID(userID)
	if err != nil {
		c.handleUseCaseError(w, err)
		return
	}

	// Map user entity to response DTO
	response := c.mapUserToProfileResponse(user)

	// Return success response
	w.WriteHeader(http.StatusOK)
	shared.Success(w, "User retrieved successfully", response)
}

// Helper functions

// decodeAndValidateRequest decodes and validates the request body
func (c *AuthController) decodeAndValidateRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		shared.Error(w, http.StatusBadRequest, "Invalid JSON body", err.Error())
		return err
	}

	if err := c.validate.Struct(req); err != nil {
		shared.Error(w, http.StatusBadRequest, "Validation failed", err.Error())
		return err
	}

	return nil
}

// getUserIDFromContext extracts and validates the user ID from the request context
func (c *AuthController) getUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	userID := r.Context().Value("userID")
	if userID == nil {
		return uuid.Nil, errors.New("user not found in context")
	}

	userIDStr, ok := userID.(string)
	if !ok {
		return uuid.Nil, errors.New("user ID in context is not a string")
	}

	uuidValue, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid user ID format")
	}

	return uuidValue, nil
}

// mapUserToProfileResponse maps a user entity to a profile response DTO
func (c *AuthController) mapUserToProfileResponse(user *entities.User) dtos.UserProfileResponse {
	return dtos.UserProfileResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// handleUseCaseError maps use case errors to appropriate HTTP responses
func (c *AuthController) handleUseCaseError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		shared.Error(w, http.StatusNotFound, "User not found", err.Error())
	case errors.Is(err, domain.ErrUsernameExists), errors.Is(err, domain.ErrEmailExists):
		shared.Error(w, http.StatusConflict, err.Error(), nil)
	case errors.Is(err, domain.ErrInvalidPassword):
		shared.Error(w, http.StatusUnauthorized, "Invalid credentials", nil)
	case errors.Is(err, domain.ErrJWTGeneration):
		shared.Error(w, http.StatusInternalServerError, "Authentication error", nil)
	default:
		shared.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
	}
}

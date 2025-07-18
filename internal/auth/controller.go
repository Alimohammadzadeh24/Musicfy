package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"musicfy/internal/auth/dtos"
	"musicfy/internal/auth/models"
	"musicfy/internal/shared"
)

// validator instance is created once and reused
var validate = validator.New()

// RegisterUserController handles user registration requests
func RegisterUserController(w http.ResponseWriter, r *http.Request) {
	// Parse and validate request body
	var req dtos.RegisterRequest
	if err := decodeAndValidateRequest(w, r, &req); err != nil {
		return
	}

	// Convert age string to int
	age, err := strconv.Atoi(req.Age)
	if err != nil {
		shared.Error(w, http.StatusBadRequest, "Invalid age format", err.Error())
		return
	}

	// Create user model from request
	newUser := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.UserName,
		Email:     req.Email,
		Age:       age,
	}

	// Register user through service layer
	if err := RegisterUserService(newUser, req.Password); err != nil {
		handleServiceError(w, err)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	shared.Success(w, "User created successfully", nil)
}

// LoginUserController handles user login requests
func LoginUserController(w http.ResponseWriter, r *http.Request) {
	// Parse and validate request body
	var req dtos.LoginRequestDto
	if err := decodeAndValidateRequest(w, r, &req); err != nil {
		return
	}

	// Authenticate user through service layer
	token, err := LoginUserService(req.UsernameOrEmail, req.Password)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	// Return success response with token
	w.WriteHeader(http.StatusOK)
	shared.Success(w, "Login successful", map[string]string{
		"token": token,
	})
}

// GetUserProfileController retrieves the profile of the authenticated user
func GetUserProfileController(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID, err := getUserIDFromContext(r)
	if err != nil {
		shared.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Get user from service layer
	user, err := GetUserByIDService(userID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	// Map user model to response DTO
	response := mapUserToProfileResponse(user)

	// Return success response
	w.WriteHeader(http.StatusOK)
	shared.Success(w, "User retrieved successfully", response)
}

// Helper functions

// decodeAndValidateRequest decodes and validates the request body
func decodeAndValidateRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		shared.Error(w, http.StatusBadRequest, "Invalid JSON body", err.Error())
		return err
	}

	if err := validate.Struct(req); err != nil {
		shared.Error(w, http.StatusBadRequest, "Validation failed", err.Error())
		return err
	}

	return nil
}

// getUserIDFromContext extracts and validates the user ID from the request context
func getUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	userID := r.Context().Value(userIDKey)
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

// mapUserToProfileResponse maps a user model to a profile response DTO
func mapUserToProfileResponse(user *models.User) dtos.UserProfileResponseDto {
	return dtos.UserProfileResponseDto{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// handleServiceError maps service errors to appropriate HTTP responses
func handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrUserNotFound):
		shared.Error(w, http.StatusNotFound, "User not found", err.Error())
	case errors.Is(err, ErrUsernameExists), errors.Is(err, ErrEmailExists):
		shared.Error(w, http.StatusConflict, err.Error(), nil)
	case errors.Is(err, ErrInvalidPassword):
		shared.Error(w, http.StatusUnauthorized, "Invalid credentials", nil)
	case errors.Is(err, ErrJWTGeneration):
		shared.Error(w, http.StatusInternalServerError, "Authentication error", nil)
	default:
		shared.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
	}
}

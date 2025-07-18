package usecases

import (
	"errors"
	"musicfy/internal/auth/domain"
	"musicfy/internal/auth/domain/entities"
	"musicfy/internal/auth/domain/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthUseCase handles authentication business logic
type AuthUseCase struct {
	userRepository repositories.UserRepository
	jwtService     JWTService
}

// NewAuthUseCase creates a new auth use case
func NewAuthUseCase(userRepo repositories.UserRepository, jwtService JWTService) *AuthUseCase {
	return &AuthUseCase{
		userRepository: userRepo,
		jwtService:     jwtService,
	}
}

// RegisterUser handles user registration
func (uc *AuthUseCase) RegisterUser(firstName, lastName, username, email, password string, age int) error {
	// Check for duplicate username
	existingByUsername, err := uc.userRepository.FindByUsername(username)
	if err != nil {
		return err
	}
	if existingByUsername != nil {
		return domain.ErrUsernameExists
	}

	// Check for duplicate email
	existingByEmail, err := uc.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	if existingByEmail != nil {
		return domain.ErrEmailExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Create user
	newUser := entities.NewUser(firstName, lastName, username, email, age, string(hashedPassword))
	return uc.userRepository.Create(newUser)
}

// LoginUser handles user login
func (uc *AuthUseCase) LoginUser(usernameOrEmail, password string) (string, error) {
	// Find user
	user, err := uc.userRepository.FindByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", domain.ErrUserNotFound
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", domain.ErrInvalidPassword
	}

	// Generate token
	token, err := uc.jwtService.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", domain.ErrJWTGeneration
	}

	return token, nil
}

// GetUserByID retrieves a user by ID
func (uc *AuthUseCase) GetUserByID(id uuid.UUID) (*entities.User, error) {
	user, err := uc.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

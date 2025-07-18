package auth

import (
	"errors"
	"musicfy/internal/auth/models"
	"musicfy/internal/auth/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var userRepo repository.UserRepository

// init initializes the repository
func init() {
	userRepo = repository.NewUserRepository()
}

func RegisterUserService(newUser models.User, rawPassword string) error {
	// 1. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(rawPassword), bcrypt.DefaultCost,
	)
	if err != nil {
		return errors.New("failed to hash password")
	}
	newUser.PasswordHash = string(hashedPassword)

	// 2. Check for duplicate username/email
	existingByUsername, err := userRepo.FindByUsername(newUser.Username)
	if err != nil {
		return err
	}
	if existingByUsername != nil {
		return ErrUsernameExists
	}

	existingByEmail, err := userRepo.FindByEmail(newUser.Email)
	if err != nil {
		return err
	}
	if existingByEmail != nil {
		return ErrEmailExists
	}

	// 3. Create user
	newUser.ID = uuid.New()
	if err := userRepo.Create(&newUser); err != nil {
		return err
	}

	return nil
}

func LoginUserService(usernameOrEmail, password string) (string, error) {
	// 1. Find user by username or email
	existingUser, err := userRepo.FindByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		return "", err
	}
	if existingUser == nil {
		return "", ErrUserNotFound
	}

	// 2. Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(password)); err != nil {
		return "", ErrInvalidPassword
	}

	// 3. Generate JWT token
	token, err := GenerateJWT(existingUser.ID, existingUser.Username)
	if err != nil {
		return "", ErrJWTGeneration
	}
	return token, nil
}

func GetUserByIDService(id uuid.UUID) (*models.User, error) {
	user, err := userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// GetUserByEmail finds a user by email
func GetUserByEmail(email string) (*models.User, error) {
	user, err := userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

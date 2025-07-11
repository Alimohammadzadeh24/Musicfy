package auth

import (
	"errors"
	"musicfy/internal/auth/models"
	"musicfy/internal/db"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserService(newUser models.User, rawPassword string) error {
	// 1. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(rawPassword), bcrypt.DefaultCost,
	)
	if err != nil {
		return errors.New("failed to hash password")
	}
	newUser.PasswordHash = string(hashedPassword)

	// 2. Optionally: check for duplicate username/email in DB
	dbInstance := db.GetDatabase()
	var existingUser models.User
	if err := dbInstance.Where("username = ?", newUser.Username).Or("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		if existingUser.Username == newUser.Username {
			return errors.New(ErrUsernameExists.Error())
		}
		if existingUser.Email == newUser.Email {
			return errors.New(ErrEmailExists.Error())
		}
	}
	newUser.ID = uuid.New()

	// 3. Save to DB
	if err := dbInstance.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func LoginUserService(usernameOrEmail, password string) (string, error) {
	// 1. Find user by username or email
	dbInstance := db.GetDatabase()
	var existingUser models.User
	if err := dbInstance.Where("username = ?", usernameOrEmail).Or("email = ?", usernameOrEmail).First(&existingUser).Error; err != nil {
		return "", ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(password)); err != nil {
		return "", ErrInvalidPassword
	}

	token, err := GenerateJWT(existingUser.ID, existingUser.Username)
	if err != nil {
		return "", ErrJWTGeneration
	}
	return token, nil
}

func GetUserByIDService(id uuid.UUID) (*models.User, error) {
	// 1. Find user by ID
	dbInstance := db.GetDatabase()
	var existingUser models.User
	if err := dbInstance.Where("ID = ?", id).First(&existingUser).Error; err != nil {
		return nil, ErrUserNotFound
	}
	return &existingUser, nil
}

// Example: Get user by email using GORM
func GetUserByEmail(email string) (*models.User, error) {
	dbInstance := db.GetDatabase()
	var user models.User
	if err := dbInstance.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

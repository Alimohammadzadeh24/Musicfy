package auth

import (
	"errors"
	"musicfy/internal/auth/models"

	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

// Fake in-memory DB for now
var userStore = []models.User{} // In a real app, this would be a database or repo layer

func RegisterUserService(newUser models.User, rawPassword string) error {
	// 1. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(rawPassword), bcrypt.DefaultCost,
	)
	if err != nil {
		return errors.New("failed to hash password")
	}
	newUser.PasswordHash = string(hashedPassword)

	// 2. Optionally: check for duplicate username/email in userStore here...
	if newUser.Username == "" {
		newUser.Username = newUser.FirstName + newUser.LastName
	} else {
		if userIsExsit(userStore, newUser.Username) {
			return errors.New(ErrUsernameExists.Error())
		}
	}

	if emailIsExsit(userStore, newUser.Email) {
		return errors.New(ErrEmailExists.Error())
	}

	newUser.ID = uuid.New()

	// 3. Save to store
	userStore = append(userStore, newUser)

	return nil
}

func LoginUserService(usernameOrEmail, password string) (string, error) {
	// 1. Find user by username or email
	for _, user := range userStore {
		if user.Username == usernameOrEmail || user.Email == usernameOrEmail {
			if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) == nil {
				// 2. If found, check password
				// 3. Generate JWT token
				token, err := GenerateJWT(user.ID, user.Username)
				if err != nil {
					return "", errors.New(ErrJWTGeneration.Error())
				}
				return token, nil
			} else {
				return "", errors.New(ErrInvalidPassword.Error())
			}
		}
	}

	return "", errors.New(ErrUserNotFound.Error())
}

func GetUserByIDService(id uuid.UUID) (*models.User, error) {
	// 1. Find user by ID
	for _, user := range userStore {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, errors.New(ErrUserNotFound.Error())
}

// findById
func userIsExsit(users []models.User, username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}

func emailIsExsit(users []models.User, email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}
	return false
}

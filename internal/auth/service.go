package auth

import (
	"errors"
	"musicfy/internal/auth/models"

	"golang.org/x/crypto/bcrypt"
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
	}else{
		if(userIsExsit(userStore, newUser.Username)){
			return errors.New("username already exists")
		}
	}

	// 3. Save to store
	userStore = append(userStore, newUser)

	return nil
}


//findById
func userIsExsit(users []models.User, username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}
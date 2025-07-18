package entities

import (
	"time"

	"github.com/google/uuid"
)

// User represents the core user entity in the domain
type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	Username     string
	Email        string
	Age          int
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewUser creates a new user with default values
func NewUser(firstName, lastName, username, email string, age int, passwordHash string) *User {
	now := time.Now()
	return &User{
		ID:           uuid.New(),
		FirstName:    firstName,
		LastName:     lastName,
		Username:     username,
		Email:        email,
		Age:          age,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

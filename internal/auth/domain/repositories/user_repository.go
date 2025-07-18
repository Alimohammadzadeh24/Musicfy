package repositories

import (
	"musicfy/internal/auth/domain/entities"

	"github.com/google/uuid"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	// Create inserts a new user into the database
	Create(user *entities.User) error

	// FindByUsername finds a user by username
	FindByUsername(username string) (*entities.User, error)

	// FindByEmail finds a user by email
	FindByEmail(email string) (*entities.User, error)

	// FindByUsernameOrEmail finds a user by username or email
	FindByUsernameOrEmail(usernameOrEmail string) (*entities.User, error)

	// FindByID finds a user by ID
	FindByID(id uuid.UUID) (*entities.User, error)

	// Update updates an existing user in the database
	Update(user *entities.User) error
}

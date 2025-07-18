package repositories

import (
	"database/sql"
	"errors"
	"musicfy/internal/auth/domain/entities"
	"musicfy/internal/auth/domain/repositories"
	"musicfy/internal/db"
	"time"

	"github.com/google/uuid"
)

// UserRepositoryImpl implements the UserRepository interface for PostgreSQL
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository creates a new PostgreSQL user repository
func NewUserRepository() repositories.UserRepository {
	return &UserRepositoryImpl{
		db: db.GetDB(),
	}
}

// Create inserts a new user into the database
func (r *UserRepositoryImpl) Create(user *entities.User) error {
	query := `
		INSERT INTO users (id, first_name, last_name, username, email, age, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Age,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// FindByUsername finds a user by username
func (r *UserRepositoryImpl) FindByUsername(username string) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	return r.findOneByQuery(query, username)
}

// FindByEmail finds a user by email
func (r *UserRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	return r.findOneByQuery(query, email)
}

// FindByUsernameOrEmail finds a user by username or email
func (r *UserRepositoryImpl) FindByUsernameOrEmail(usernameOrEmail string) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
		FROM users
		WHERE username = $1 OR email = $1
	`

	return r.findOneByQuery(query, usernameOrEmail)
}

// FindByID finds a user by ID
func (r *UserRepositoryImpl) FindByID(id uuid.UUID) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	return r.findOneByQuery(query, id)
}

// Update updates an existing user in the database
func (r *UserRepositoryImpl) Update(user *entities.User) error {
	query := `
		UPDATE users
		SET first_name = $1, last_name = $2, username = $3, email = $4, 
		    age = $5, password_hash = $6, updated_at = $7
		WHERE id = $8
	`

	user.UpdatedAt = time.Now()

	_, err := r.db.Exec(
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Age,
		user.PasswordHash,
		user.UpdatedAt,
		user.ID,
	)

	return err
}

// Helper function to find one user by a query
func (r *UserRepositoryImpl) findOneByQuery(query string, args ...interface{}) (*entities.User, error) {
	var user entities.User

	row := r.db.QueryRow(query, args...)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Age,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

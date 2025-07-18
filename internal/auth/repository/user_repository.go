package repository

import (
	"database/sql"
	"errors"
	"musicfy/internal/auth/models"
	"musicfy/internal/db"
	"time"

	"github.com/google/uuid"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
	FindByID(id uuid.UUID) (*models.User, error)
	Update(user *models.User) error
}

// PostgresUserRepository implements UserRepository for PostgreSQL
type PostgresUserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new PostgreSQL user repository
func NewUserRepository() UserRepository {
	return &PostgresUserRepository{
		db: db.GetDB(),
	}
}

// Create inserts a new user into the database
func (r *PostgresUserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (id, first_name, last_name, username, email, age, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

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
func (r *PostgresUserRepository) FindByUsername(username string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	return r.findOneByQuery(query, username)
}

// FindByEmail finds a user by email
func (r *PostgresUserRepository) FindByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	return r.findOneByQuery(query, email)
}

// FindByUsernameOrEmail finds a user by username or email
func (r *PostgresUserRepository) FindByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
		FROM users
		WHERE username = $1 OR email = $1
	`

	return r.findOneByQuery(query, usernameOrEmail)
}

// FindByID finds a user by ID
func (r *PostgresUserRepository) FindByID(id uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	return r.findOneByQuery(query, id)
}

// Update updates an existing user in the database
func (r *PostgresUserRepository) Update(user *models.User) error {
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
func (r *PostgresUserRepository) findOneByQuery(query string, args ...interface{}) (*models.User, error) {
	var user models.User

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

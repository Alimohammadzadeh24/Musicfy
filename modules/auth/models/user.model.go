package models

import "time"

type User struct {
	ID           int64 // Unique identifier (can be UUID or int64)
	FirstName    string
	LastName     string
	Username     string
	Email        string
	Age          int
	PasswordHash string // Store hashed password, not raw
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

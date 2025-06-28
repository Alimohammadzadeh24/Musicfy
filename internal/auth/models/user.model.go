package models

import "time"

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Username     string
	Email        string
	Age          int
	Token        string
	PasswordHash string // Store hashed password, not raw
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

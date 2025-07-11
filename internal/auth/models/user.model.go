package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName    string    `gorm:"size:100;not null"`
	LastName     string    `gorm:"size:100;not null"`
	Username     string    `gorm:"size:100;unique;not null"`
	Email        string    `gorm:"size:100;unique;not null"`
	Age          int       `gorm:"not null"`
	Token        string    `gorm:"size:255"`
	PasswordHash string    `gorm:"size:255;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

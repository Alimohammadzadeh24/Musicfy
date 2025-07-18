package usecases

import "github.com/google/uuid"

// JWTService defines the interface for JWT operations
type JWTService interface {
	// GenerateToken creates a new JWT token for a user
	GenerateToken(userID uuid.UUID, username string) (string, error)

	// ValidateToken validates a JWT token and returns the claims
	ValidateToken(tokenString string) (*JWTClaims, error)
}

// JWTClaims represents the claims in a JWT token
type JWTClaims struct {
	UserID   uuid.UUID
	Username string
}

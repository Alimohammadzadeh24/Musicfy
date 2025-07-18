package services

import (
	"log"
	"musicfy/internal/auth/domain/usecases"
	"musicfy/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTServiceImpl implements the JWTService interface
type JWTServiceImpl struct {
	jwtKey      []byte
	expiryHours int
}

// jwtClaims is the internal claims structure for JWT
type jwtClaims struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

// NewJWTService creates a new JWT service
func NewJWTService() *JWTServiceImpl {
	// Ensure config is loaded
	if config.AppConfig.JWTConfig.Secret == "" {
		config.LoadConfig()
	}

	jwtSecret := config.AppConfig.JWTConfig.Secret
	if jwtSecret == "" {
		log.Fatalf("JWT secret not configured")
	}

	expiryHours := config.AppConfig.JWTConfig.ExpiryHours
	if expiryHours <= 0 {
		expiryHours = 24 // Default to 24 hours
	}

	return &JWTServiceImpl{
		jwtKey:      []byte(jwtSecret),
		expiryHours: expiryHours,
	}
}

// GenerateToken creates a new JWT token for a user
func (s *JWTServiceImpl) GenerateToken(userID uuid.UUID, username string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(s.expiryHours) * time.Hour)
	claims := &jwtClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtKey)
}

// ValidateToken validates a JWT token and returns the claims
func (s *JWTServiceImpl) ValidateToken(tokenString string) (*usecases.JWTClaims, error) {
	claims := &jwtClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return &usecases.JWTClaims{
		UserID:   claims.UserID,
		Username: claims.Username,
	}, nil
}

package auth

import (
	"log"
	"musicfy/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey []byte

// InitJWT initializes the JWT key from configuration
func InitJWT() {
	// Ensure config is loaded
	if config.AppConfig.JWTConfig.Secret == "" {
		config.LoadConfig()
	}

	jwtSecret := config.AppConfig.JWTConfig.Secret
	if jwtSecret == "" {
		log.Fatalf("JWT secret not configured")
	}

	jwtKey = []byte(jwtSecret)
	log.Println("JWT configuration loaded")
}

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uuid.UUID, username string) (string, error) {
	// Get expiry hours from config
	expiryHours := config.AppConfig.JWTConfig.ExpiryHours
	if expiryHours <= 0 {
		expiryHours = 24 // Default to 24 hours
	}

	expirationTime := time.Now().Add(time.Duration(expiryHours) * time.Hour)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

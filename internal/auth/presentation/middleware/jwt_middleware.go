package middleware

import (
	"context"
	"musicfy/internal/auth/domain/usecases"
	"musicfy/internal/shared"
	"net/http"
	"strings"
)

// JWTMiddleware handles JWT authentication
type JWTMiddleware struct {
	jwtService usecases.JWTService
}

// NewJWTMiddleware creates a new JWT middleware
func NewJWTMiddleware(jwtService usecases.JWTService) *JWTMiddleware {
	return &JWTMiddleware{
		jwtService: jwtService,
	}
}

// Middleware returns a middleware function that validates JWT tokens
func (m *JWTMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			shared.Error(w, http.StatusUnauthorized, "Unauthorized: missing or invalid token", nil)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token
		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			shared.Error(w, http.StatusUnauthorized, "Unauthorized: invalid token", nil)
			return
		}

		// Add user ID to context
		ctx := context.WithValue(r.Context(), "userID", claims.UserID.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

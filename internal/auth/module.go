package auth

import (
	"musicfy/internal/auth/presentation/routes"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all auth routes with the given router
func RegisterRoutes(router *mux.Router) {
	routes.RegisterAuthRoutes(router)
}

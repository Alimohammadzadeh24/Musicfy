package auth

import (
	"github.com/gorilla/mux"
	//locals
)

// RegisterAuthRoutes mounts the auth-related routes on the given router
func RegisterAuthRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	// Public routes
	authRouter.HandleFunc("/register", RegisterUserController).Methods("POST")
	authRouter.HandleFunc("/login", LoginUserController).Methods("POST")

	// Protected routes
	protected := authRouter.PathPrefix("").Subrouter()
	protected.Use(JWTMiddleware)
	protected.HandleFunc("/user/profile", GetUserProfileController).Methods("GET")
}
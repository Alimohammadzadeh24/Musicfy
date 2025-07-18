package routes

import (
	"musicfy/internal/auth/data/repositories"
	"musicfy/internal/auth/data/services"
	"musicfy/internal/auth/domain/usecases"
	"musicfy/internal/auth/presentation/controllers"
	"musicfy/internal/auth/presentation/middleware"

	"github.com/gorilla/mux"
)

// RegisterAuthRoutes sets up authentication routes
func RegisterAuthRoutes(router *mux.Router) {
	// Initialize dependencies
	userRepository := repositories.NewUserRepository()
	jwtService := services.NewJWTService()
	authUseCase := usecases.NewAuthUseCase(userRepository, jwtService)
	authController := controllers.NewAuthController(authUseCase)
	jwtMiddleware := middleware.NewJWTMiddleware(jwtService)

	// Create subrouter for auth routes
	authRouter := router.PathPrefix("/auth").Subrouter()

	// Public routes
	authRouter.HandleFunc("/register", authController.Register).Methods("POST")
	authRouter.HandleFunc("/login", authController.Login).Methods("POST")

	// Protected routes
	protected := authRouter.PathPrefix("").Subrouter()
	protected.Use(jwtMiddleware.Middleware)
	protected.HandleFunc("/profile", authController.GetProfile).Methods("GET")
}

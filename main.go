package main

import (
	"log"
	"musicfy/internal/auth"
	"musicfy/internal/config"
	"musicfy/internal/db"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration from .env file
	config.LoadConfig()

	// Initialize database connection
	db.InitializeDatabase()

	// Run database migrations
	if err := db.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Set up router with environment-specific settings
	router := setupRouter()

	// Get server configuration
	host := config.AppConfig.ServerConfig.Host
	port := config.AppConfig.ServerConfig.Port

	// Log startup information
	log.Printf("Server starting on %s:%s in %s mode", host, port, config.AppConfig.Environment)

	// Start server
	serverAddr := host + ":" + port
	log.Fatal(http.ListenAndServe(serverAddr, router))
}

// setupRouter configures the HTTP router with routes
func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Add environment indicator to response headers if in development mode
	if config.IsDevelopment() {
		router.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Environment", string(config.AppConfig.Environment))
				next.ServeHTTP(w, r)
			})
		})
	}

	// API router
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// Root endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Musicfy!"))
	}).Methods("GET")

	// Add environment info endpoint in non-production environments
	if !config.IsProduction() {
		router.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Current environment: " + string(config.AppConfig.Environment)))
		}).Methods("GET")
	}

	// Register auth routes
	auth.RegisterRoutes(apiRouter)

	return router
}

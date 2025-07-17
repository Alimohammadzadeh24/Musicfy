package main

import (
	"log"
	"musicfy/internal/auth"
	"musicfy/internal/auth/models"
	"musicfy/internal/db"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	auth.InitJWT()
	db.InitializeDatabase()
	database := db.GetDatabase()
	database.AutoMigrate(&models.User{})

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Musicfy!"))
	}).Methods("GET")

	// Register auth routes
	auth.RegisterAuthRoutes(apiRouter)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server starting on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

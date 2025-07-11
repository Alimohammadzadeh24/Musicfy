package main

import (
	"log"
	"musicfy/internal/auth"
	"musicfy/internal/auth/models"
	"musicfy/internal/db"
	"net/http"
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

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

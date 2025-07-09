package main

import (
	"log"
	"net/http"
	"musicfy/internal/auth"
	"github.com/gorilla/mux"
)

func main() {
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

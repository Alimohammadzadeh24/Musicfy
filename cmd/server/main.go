package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"musicfy/internal/auth"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Musicfy!"))
	}).Methods("GET")

	// Register auth routes
	auth.RegisterAuthRoutes(router)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"musicfy/modules/auth"
)

func main() {
	router := mux.NewRouter()

	// Register auth routes
	auth.RegisterAuthRoutes(router)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

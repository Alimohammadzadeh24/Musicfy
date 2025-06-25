package auth

import (
	"github.com/gorilla/mux"
	
	//locals 
)

// RegisterAuthRoutes mounts the auth-related routes on the given router
func RegisterAuthRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/register", RegisterController).Methods("POST")
}

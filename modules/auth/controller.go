package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"musicfy/modules/auth/dtos"
	"musicfy/modules/auth/models"
)
import "strconv"

func RegisterController(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var req dtos.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	
	var newUser models.User
	age, err := strconv.Atoi(req.Age)

	if err != nil {
		if err := validate.Struct(req); err != nil {
			http.Error(w, "You age is not valid", http.StatusBadRequest)
			return
		}
	} else {
		newUser = models.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Username:  req.UserName,
			Email:     req.Email,
			Age:       age,
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

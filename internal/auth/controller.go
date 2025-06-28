package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"musicfy/internal/auth/dtos"
	"musicfy/internal/auth/models"
	"musicfy/internal/shared"
	"strconv"
)

func RegisterUserController(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var req dtos.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.Error(w, http.StatusBadRequest, "Invalid JSON body", err.Error())
		return
	}
	if err := validate.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(req.Age)

	if err != nil {
		http.Error(w, "Invalid age format", http.StatusBadRequest)
		return
	} else {
		newUser := models.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Username:  req.UserName,
			Email:     req.Email,
			Age:       age,
		}
		if err := RegisterUserService(newUser, req.Password); err != nil {
			shared.Error(w, http.StatusInternalServerError, "Failed to register user: "+err.Error(), err.Error())
			return
		}
		fmt.Println("new user created : ", newUser)
	}

	w.WriteHeader(http.StatusCreated)
	shared.Success(w, "User Created Successfully", nil)
}

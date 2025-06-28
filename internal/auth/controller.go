package auth

import (
	"encoding/json"
	"errors"
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
			shared.Error(w, http.StatusBadRequest, "Failed to register user", err.Error())
			return
		}
		fmt.Println("new user created : ", newUser)
	}

	w.WriteHeader(http.StatusCreated)
	shared.Success(w, "User Created Successfully", nil)
}

func LoginUserController(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var req dtos.LoginRequestDto

	if error := json.NewDecoder(r.Body).Decode(&req); error != nil {
		shared.Error(w, http.StatusBadRequest, "Invalid JSON body", error.Error())
		return
	}
	if error := validate.Struct(req); error != nil {
		http.Error(w, "Validation failed: "+error.Error(), http.StatusBadRequest)
		return
	}

	token, err := LoginUserService(req.UsernameOrEmail, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserNotFound):
			shared.Error(w, http.StatusNotFound, err.Error(), err.Error())
		case errors.Is(err, ErrUsernameExists), errors.Is(err, ErrEmailExists):
			shared.Error(w, http.StatusConflict, err.Error(), err.Error())
		case errors.Is(err, ErrInvalidPassword):
			shared.Error(w, http.StatusUnauthorized, err.Error(), err.Error())
		case errors.Is(err, ErrJWTGeneration):
			shared.Error(w, http.StatusInternalServerError, err.Error(), err.Error())
		default:
			shared.Error(w, http.StatusInternalServerError, "", err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	shared.Success(w, "Login successful", map[string]string{
		"token": token,
	})
}

func GetUserProfileController(w http.ResponseWriter, r *http.Request) {
	userInfo, err := ValidateJWT(r.Header.Get("authentication"))
	if err != nil {
		shared.Error(w, http.StatusBadRequest, "Invalid or missing token", err.Error())
	}

	user, err := GetUserByIDService(userInfo.ID)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserNotFound):
			shared.Error(w, http.StatusNotFound, err.Error(), err.Error())
		default:
			shared.Error(w, http.StatusInternalServerError, "", err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	shared.Success(w, "User retrieved successfully", user)
}

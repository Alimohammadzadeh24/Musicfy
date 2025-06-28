package auth

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUsernameExists      = errors.New("username already exists")
	ErrEmailExists         = errors.New("email already exists")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrJWTGeneration       = errors.New("failed to generate JWT token")
	ErrInternalServerError = errors.New("internal server error")
)

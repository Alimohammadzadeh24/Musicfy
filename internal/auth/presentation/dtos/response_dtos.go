package dtos

import "time"

// UserProfileResponse represents the user profile data
type UserProfileResponse struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginResponse represents the login response data
type LoginResponse struct {
	Token string `json:"token"`
}

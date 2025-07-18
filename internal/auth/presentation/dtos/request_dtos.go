package dtos

// LoginRequest represents the login request data
type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

// RegisterRequest represents the registration request data
type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name" validate:"required,min=2"`
	Username  string `json:"username" validate:"required,min=3"`
	Password  string `json:"password" validate:"required,min=8,max=20"`
	Email     string `json:"email" validate:"required,email"`
	Age       string `json:"age" validate:"required"`
}

package dtos

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name" validate:"required,min=2"`
	UserName  string `json:"username"`
	Password  string `json:"password" validate:"required,min=8,max=20"`
	Email     string `json:"email" validate:"required,email"`
	Age       string `json:"age"`
}

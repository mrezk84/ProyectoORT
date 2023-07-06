package dtos

type RegisterUser struct {
	Email    string `json:"email" validate:"required,email"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

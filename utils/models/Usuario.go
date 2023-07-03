package models

type Usuario struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserModel struct {
	Users []Usuario
}
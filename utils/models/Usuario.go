package models

type Usuario struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

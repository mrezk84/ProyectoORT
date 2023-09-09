package entity

type Usuario struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"username"`
	Password string `db:"password"`
}

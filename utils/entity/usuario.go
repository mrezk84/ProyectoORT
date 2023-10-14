package entity

type Usuario struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"username"`
	Password string `db:"password"`
}

package entity

type Control struct {
	ID          int    `db:"id"`
	Descripcion string `db:"descripcion"`
	Tipo        string `db:"tipo"`
}

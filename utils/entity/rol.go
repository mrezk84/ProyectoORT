package entity

type Rol struct {
	ID     int    `db:"id"`
	Nombre string `db:"nombre"`
}

package entity

type Obra struct {
	ID     int    `db:"id_obra"`
	Nombre string `db:"nombre"`
}

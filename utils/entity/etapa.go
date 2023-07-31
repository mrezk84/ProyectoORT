package entity

type Etapa struct {
	ID     int    `db:"id_etapa"`
	Nombre string `db:"nombre"`
}

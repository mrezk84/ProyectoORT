package entity

type Etapa struct {
	ID     int    `db:"id_etapa"`
	nombre string `db:"nombre"`
}

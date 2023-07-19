package entity

import "time"

type Formulario struct {
	ID          int       `db:"id_formulario"`
	Nombre      string    `db:"nombre"`
	Informacion string    `db:"informacion"`
	Version     int       `db:"version"`
	Fecha       time.Time `db:"fecha"`
}

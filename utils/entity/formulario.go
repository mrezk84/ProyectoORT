package entity

type Formulario struct {
	ID          int    `db:"id"`
	Nombre      string `db:"nombre"`
	Informacion string `db:"informacion"`
	Version     string `db:"version"`
	IdControl   int    `db:"control_id"`
}

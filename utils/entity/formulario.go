package entity

type Formulario struct {
	ID          int    `db:"id"`
	Nombre      string `db:"nombre"`
	Informacion string `db:"informacion"`
	Version     int    `db:"version"`
	ControlID   int    `db:"control_id"`
	UsuarioID   int    `db:"usuario_id"`
}

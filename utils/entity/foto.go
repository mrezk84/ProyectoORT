package entity

type Foto struct {
	ID           int    `db:"id"`
	Nombre       string `db:"nombre"`
	Notas        string `db:"notas"`
	FormularioID int    `db:"formulario_id"`
}

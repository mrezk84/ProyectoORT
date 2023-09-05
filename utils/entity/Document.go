package entity

type Document struct {
	ID           int64 `db:"id"`
	FormularioID int64 `db:"formulario_id"`
	ObraID       int64 `db:"obra_id"`
	PisoID       int64 `db:"piso_id"`
}

package entity

type ControlForm struct {
	ID           int64 `db:"id"`
	ControlID    int64 `db:"control_id"`
	FormularioID int64 `db:"formulario_id"`
}

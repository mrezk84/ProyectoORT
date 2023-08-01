package entity

type CheckFormulario struct {
	ID           int64 `db:"id"`
	CheckID      int64 `json:"check_id"`
	FormularioID int64 `json:"formulario_id"`
}

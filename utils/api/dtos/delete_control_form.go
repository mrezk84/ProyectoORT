package dtos

type DeleteControlForm struct {
	ControlID    int64 `param:"control_id"`
	FormularioID int64 `json:"formulario_id"`
}

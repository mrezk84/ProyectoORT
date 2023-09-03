package dtos

type DocumentAudit struct {
	ID          int    `json:"id_formulario"`
	Informacion string `json:"informacion" validate:"required"`
	Version     int    `json:"version" validate:"required"`
	Nombre      string `json:"nombre" validate:"required"`
	ControlID   int    `json:"control_id" validate:"required"`
	UsuarioID   int    `json:"usuario_id" validate:"required"`
}

package dtos

type DocumentAudit struct {
	ID          int    `json:"id_formulario"`
	Informacion string `json:"informacion"`
	Version     int    `json:"version" `
	Nombre      string `json:"nombre" `
	ControlID   int    `json:"control_id"`
	UsuarioID   int    `json:"usuario_id"`
}

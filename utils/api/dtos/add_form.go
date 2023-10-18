package dtos

type AddForm struct {
	Nombre      string `json:"nombre"`
	Informacion string `json:"informacion"`
	Version     int    `json:"version"`
	ControlID   int    `json:"control_id"`
	UsuarioID   int    `json:"usuario_id"`
}

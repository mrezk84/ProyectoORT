package dtos

type FotoDTO struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre" `
	Notas        string `json:"notas" `
	FormularioID int    `json:"formulario_id"`
}

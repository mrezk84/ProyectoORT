package dtos

type FotoDTO struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre" validate:"required"`
	Notas        string `json:"notas" validate:"required"`
	FormularioID int    `json:"formulario_id" validate:"required"`
}

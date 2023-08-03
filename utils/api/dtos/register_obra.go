package dtos

type RegisterObra struct {
	Nombre string `json:"name" validate:"required"`
}

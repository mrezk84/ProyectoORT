package dtos

type EliminarObra struct {
	Nombre string `json:"name" validate:"required"`
}

package dtos

type RegisterEtapa struct {
	Nombre string `json:"nombre" validate:"required"`
}

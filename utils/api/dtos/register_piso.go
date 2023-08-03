package dtos

type RegisterPiso struct {
	Numero int `json:"Numero" validate:"required"`
}

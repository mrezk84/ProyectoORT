package dtos

type RegisterPiso struct {
	ID     int `json:"id"`
	Numero int `json:"numero" validate:"required"`
}

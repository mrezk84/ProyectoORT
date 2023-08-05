package dtos

type RegisterCheck struct {
	Estado        string `json:"estado" validate:"required"`
	Observaciones string `json:"observaciones" validate:"required"`
	Version       int    `json:"version" validate:"required"`
	Fecha         string `json:"fecha_control" validate:"required"`
}

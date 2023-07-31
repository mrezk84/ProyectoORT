package dtos

type RegisterControl struct {
	ID          int    `json:"id_control"`
	Descripcion string `json:"descripcion"`
	Tipo        string `json:"tipo"`
}

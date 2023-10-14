package dtos

type UpdateControl struct {
	ControlID   int64  `param:"id"`
	Descripcion string `json:"descripcion"`
	Tipo        string `json:"tipo"`
}

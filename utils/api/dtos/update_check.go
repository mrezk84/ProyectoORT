package dtos

type UpdateCheck struct {
	CheckID       int64  `param:"id"`
	Estado        string `json:"estado"`
	Observaciones string `json:"observaciones"`
}

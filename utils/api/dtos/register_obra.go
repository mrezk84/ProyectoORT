package dtos

type RegisterObra struct {
	ID     int    `json:"obra_id" param:"id"`
	Nombre string `json:"nombre"`
}

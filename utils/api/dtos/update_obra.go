package dtos

type UpdateObra struct {
	ObraID int64  `param:"id"`
	Nombre string `json:"nombre"`
}

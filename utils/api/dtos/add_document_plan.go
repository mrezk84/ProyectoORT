package dtos

type AddDocumentPlan struct {
	FormularioID int64 `json:"formulario_id"`
	ObraID       int64 `json:"obra_id"`
	PisoID       int64 `json:"piso_id"`
}

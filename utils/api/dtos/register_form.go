package dtos

type DocumentAudit struct {
	ID          int    `json:"id_formulario" param:"id"`
	Informacion string `json:"informacion"`
	Nombre      string `json:"nombre"`
	Version     string `json:"version"`
	Fecha       string `json:"fecha"`
}

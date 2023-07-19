package dtos

import "time"

type DocumentAudit struct {
	ID          int        `json:"id_formulario"`
	Informacion string     `json:"informacion"`
	Nombre      string     `json:"nombre"`
	Version     int        `json:"version"`
	Fecha       *time.Time `json:"fecha"`
}

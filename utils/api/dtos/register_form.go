package dtos

import "time"

type DocumentAudit struct {
	ID          int64     `json:"id_formulario"`
	Informacion string    `json:"informacion"`
	Nombre      string    `json:"nombre"`
	Version     string    `json:"version"`
	Fecha       time.Time `json:"fecha"`
	EtapaID     int64     `json:"etapa_id"`
	UsuarioID   int64     `json:"usuario_id"`
}

package dtos

import "time"

type Auditoria struct {
	ID           int       `json:"id"`
	FormularioID int       `json:"fomulario_id"`
	Version      string    `json:"version"`
	Fecha        time.Time `json:"fechaAuditoria"`
}

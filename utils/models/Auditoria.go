package models

import (
	"time"
)

type Auditoria struct {
	ID           int       `json:"id"`
	FormularioID int       `json:"formulario_id"`
	Version      string    `json:"version"`
	Fecha        time.Time `json:"fechaAuditoria"`
	Accion       string    `json:"id"`
}

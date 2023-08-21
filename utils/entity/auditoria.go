package entity

import "time"

type Auditoria struct {
	ID           int       `db:"id"`
	FormularioID int       `db:"formulario_id"`
	Version      string    `db:"version"`
	Fecha        time.Time `db:"fecha"`
	Accion       string
}

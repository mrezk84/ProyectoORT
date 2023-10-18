package models

import "time"

type Check struct {
	ID            int    `json:"id"`
	Estado        string `json:"estado"`
	FechaControl  *time.Time
	Responsable   Usuario
	Control       Control
	Document      Document
	Observaciones string `json:"observaciones"`
	Version       int
}

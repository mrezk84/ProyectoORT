package models

import "time"

type Check struct {
	ID            int
	Estado        string
	FechaControl  *time.Time
	Responsable   Usuario
	Control       Control
	Document      Document
	Observaciones string
	Version       int
}

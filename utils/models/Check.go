package models

type Check struct {
	ID            int
	Estado        string
	Responsable   Usuario
	Control       Control
	Document      Document
	Observaciones string
	Version       int
}

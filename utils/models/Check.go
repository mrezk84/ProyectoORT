package models

import "time"

type Check struct {
	ID int
	// todo estado deberia ser un enum
	Estado        string
	FechaControl  *time.Time
	Responsable   Usuario
	Obra          Obra
	Piso          Piso
	Formulario    Formulario
	Etapa         Etapa
	Control       Control
	Observaciones string
	Version       int
}

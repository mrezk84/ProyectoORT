package models

type Check struct {
	ID int
	// todo estado deberia ser un enum
	Estado        string
	Responsable   Usuario
	Obra          Obra
	Piso          Piso
	Formulario    Formulario
	Etapa         Etapa
	Control       Control
	Observaciones string
	Version       int
}

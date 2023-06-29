package model

type Formulario struct {
	ID            int
	Nombre        string
	Observaciones string
	Version       int
	Controles     []Control
	Obra          Obra
	Responsable   Usuario
}

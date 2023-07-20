package model

type Formulario struct {
	ID            int
	Nombre        string
	Observaciones string
	Version       int
	Etapas        []Etapa
	Obra          Obra
	Responsable   Usuario
}

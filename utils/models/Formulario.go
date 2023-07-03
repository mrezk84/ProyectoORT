package models

type Formulario struct {
	ID          int
	Informacion string
	Version     int
	Nombre      string
	Controles   []Control
}

package models

import "time"

type Formulario struct {
	ID          int
	Nombre      string
	Informacion string
	Version     string
	Fecha       time.Time
	EtapaID     int
	UsuarioID   int
	Fotos       []Foto
}

package models

import "time"

type Formulario struct {
	ID          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	Informacion string    `json:"informacion"`
	Version     string    `json:"version"`
	Fecha       time.Time `json:"fecha"`
	EtapaID     int       `json:"etapa"`
	UsuarioID   int       `json:"usuario"`
}

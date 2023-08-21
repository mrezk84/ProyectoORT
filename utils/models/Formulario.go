package models

import "time"

type Formulario struct {
	ID          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	Informacion string    `json:"informacion"`
	Version     string    `json:"version"`
	Fecha       time.Time `json:"fecha"`
	Etapa       []Etapa   `json:"etapa"`
	Usuario     []Usuario `json:"usuario"`
}

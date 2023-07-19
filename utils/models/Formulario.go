package models

type Formulario struct {
	ID          int       `json:"id"`
	Informacion string    `json:"informacion"`
	Version     string    `json:"version"`
	Nombre      string    `json:"nombre"`
	Controles   []Control `json:"controles"`
}

package models

type Formulario struct {
	ID          int       `json:"id"`
	Informacion string    `json:"informacion"`
	Version     int       `json:"version"`
	Nombre      string    `json:"nombre"`
	Controles   []Control `json:"controles"`
	Usuarios    []Usuario `json:"usuarios"`
	Foto        []Foto    `json:"fotos"`
}

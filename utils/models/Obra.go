package models

type Obra struct {
	ID      int    `json:"id"`
	Nombre  string `json:"nombre"`
	Capataz Usuario
	Pisos   []Piso
}

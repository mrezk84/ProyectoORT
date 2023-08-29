package model

type Obra struct {
	ID      int
	Nombre  string
	Capataz Usuario
	Pisos   []Piso
}

package models

type Etapa struct {
	ID        int64
	Nombre    string
	Controles []Control
}

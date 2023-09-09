package models

type Etapa struct {
	ID        int64     `json:"id"`
	Nombre    string    `json:"nombre"`
	Controles []Control `json:"controles"`
}

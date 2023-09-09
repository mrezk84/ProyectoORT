package models

type Control struct {
	ID          int    `json:"id"`
	Descripcion string `json:"descripcion"`
	// todo control type deberia ser un enum
	Tipo string `json:"tipo"`
}

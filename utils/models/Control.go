package models

import "time"

type Control struct {
	ID           int       `json:"id"`
	Descripcion  string    `json:"descripcion"`
	Tipo         string    `json:"tipo"`
	FechaControl time.Time `json:"fechaControl"`
	Foto         Foto
}

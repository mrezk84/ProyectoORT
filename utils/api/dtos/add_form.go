package dtos

import "time"

type AddForm struct {
	Nombre      string    `json:"nombre" validate:"required"`
	Informacion string    `json:"informacion"`
	Version     string    `json:"version" validate:"required"`
	Fecha       time.Time `json:"fecha"`
	EtapaID     int64     `json:"etapa_id"`
	UsuarioID   int64     `json:"usuario_id"`
}

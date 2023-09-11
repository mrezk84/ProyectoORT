package entity

import "time"

type Check struct {
	ID            int        `db:"id"`
	Estado        string     `db:"estado"`
	FechaControl  *time.Time `db:"fecha_control"`
	Observaciones string     `db:"observaciones"`
	Version       int        `db:"version"`
}

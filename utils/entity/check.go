package entity

type Check struct {
	ID            int    `db:"id"`
	Estado        string `db:"estado"`
	Observaciones string `db:"observaciones"`
	Version       int    `db:"version"`
	FechaControl  string `db:"fecha_control"`
}

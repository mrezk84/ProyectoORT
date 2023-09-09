package entity

type Check struct {
	ID            int    `db:"id_check"`
	Estado        string `db:"estado"`
	FechaControl  string `db:"fecha_control"`
	Observaciones string `db:"observaciones"`
	Version       int    `db:"version"`
}

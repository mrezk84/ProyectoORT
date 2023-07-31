package entity

type PisoObra struct {
	ID     int64 `db:"id"`
	PisoID int64 `db:"piso_id"`
	ObraID int64 `db:"obra_id"`
}

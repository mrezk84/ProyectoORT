package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertObra = `
		INSERT INTO Obra (nombre)
		VALUES (?, ?, ?, ?);`

	qryGetObrabyName = `
		SELECT
			ID,
			Nombre
		FROM OBRAS
		WHERE Nombre = ?;`
)

func (r *repo) SaveObra(ctx context.Context, nombre string) error {
	_, err := r.db.ExecContext(ctx, qryInsertObra, nombre)
	return err
}

func (r *repo) GetObrabyName(ctx context.Context, name string) (*entity.Obra, error) {
	o := &entity.Obra{}
	err := r.db.GetContext(ctx, o, qryGetObrabyName, name)
	if err != nil {
		return nil, err
	}

	return o, nil
}

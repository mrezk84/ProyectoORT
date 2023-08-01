package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertControl = `
		INSERT INTO CONTROL(descripcion, tipo)
		VALUES (?, ?);`

	qryGetContById = `
		SELECT
		id,
			descripcion,
			tipo,
		FROM CONTROL
		WHERE id = ?;`

	qryGetAllControls = `
		SELECT 
		id,
		descripcion,
		tipo
		FROM CONTROL;`
)

func (r *repo) SaveControl(ctx context.Context, descripcion, tipo string) error {
	_, err := r.db.ExecContext(ctx, qryInsertControl, descripcion, tipo)
	return err
}
func (r *repo) GetControls(ctx context.Context) ([]entity.Control, error) {
	cc := []entity.Control{}

	err := r.db.SelectContext(ctx, &cc, qryGetAllControls)
	if err != nil {
		return nil, err
	}

	return cc, nil
}
func (r *repo) GetFormById(ctx context.Context, id int) (*entity.Control, error) {
	c := &entity.Control{}
	err := r.db.GetContext(ctx, c, qryGetContById, id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

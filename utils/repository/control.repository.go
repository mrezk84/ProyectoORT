package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertControl = `
		INSERT INTO CONTROL(descripcion, tipo)
		VALUES (?, ?);`

	qryGetContBydescripcion = `
		SELECT
		id,
			descripcion,
			tipo,
		FROM CONTROL
		WHERE descripcion = ?;`

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
func (r *repo) GetConByDes(ctx context.Context, des string) (*entity.Control, error) {
	c := &entity.Control{}
	err := r.db.GetContext(ctx, c, qryGetContBydescripcion, des)
	if err != nil {
		return nil, err
	}

	return c, nil
}

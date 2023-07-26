package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertFrom = `
		INSERT INTO  FORMULARIO (nombre,informacion, version, fecha)
		VALUES (?, ?, ?, ?);`

	qryGetFormByDate = `
		SELECT
			id,
			nombre,
			informacion,
			version,
			fecha
		FROM FORMULARIO
		WHERE fecha = ?;`

	qryGetFormByVersion = `
		SELECT
			id,
			nombre,
			informacion,
			version,
			fecha
		FROM FORMULARIO
		WHERE version = ?;`

	qryGetAllForms = `
		SELECT
		id,
		nombre,
		informacion,
		version,
		fecha
		FROM FORMULARIO;`
)

func (r *repo) SaveFrom(ctx context.Context, nombre, informacion string, version string, fecha string) error {
	_, err := r.db.ExecContext(ctx, qryInsertFrom, informacion, nombre, version, fecha)
	return err
}

func (r *repo) GetFormByDate(ctx context.Context, fecha string) (*entity.Formulario, error) {
	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetFormByDate, fecha)
	if err != nil {
		return nil, err
	}

	return f, nil
}
func (r *repo) GetForm(ctx context.Context) ([]entity.Formulario, error) {
	ff := []entity.Formulario{}

	err := r.db.SelectContext(ctx, &ff, qryGetAllForms)
	if err != nil {
		return nil, err
	}

	return ff, nil
}

func (r *repo) GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error) {
	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetFormByVersion, version)
	if err != nil {
		return nil, err
	}

	return f, nil
}

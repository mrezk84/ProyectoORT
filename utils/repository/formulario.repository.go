package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertFrom = `
		INSERT INTO  FORMULARIO (nombre,informacion, version, fecha)
		VALUES (?, ?, ?, ?);`

	qryGetUserByDate = `
		SELECT
			id,
			nombre,
			informacion,
			version,
			fecha
		FROM FORMULARIO
		WHERE fecha = ?;`

	qryGetUserByVersion = `
		SELECT
			id,
			nombre,
			informacion,
			version,
			fecha
		FROM FORMULARIO
		WHERE version = ?;`

	qryAllGetForms = `
		SELECT
		id,
		nombre,
		infromacion,
		version,
		fecha
		FROM FORMULARIO;`
)

func (r *repo) SaveFrom(ctx context.Context, nombre, informacion string, version string, fecha string) error {
	_, err := r.db.ExecContext(ctx, qryInsertFrom, informacion, nombre, version, fecha)
	return err
}

func (r *repo) GetFormByDate(ctx context.Context, fechaIni, fechaFin string) (*entity.Formulario, error) {
	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetUserByDate, fechaIni, fechaFin)
	if err != nil {
		return nil, err
	}

	return f, nil
}
func (r *repo) GetFrom(ctx context.Context) ([]entity.Formulario, error) {
	form := []entity.Formulario{}

	err := r.db.SelectContext(ctx, &form, qryAllGetForms)
	if err != nil {
		return nil, err
	}

	return form, nil
}

func (r *repo) GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error) {
	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetUserByVersion, version)
	if err != nil {
		return nil, err
	}

	return f, nil
}

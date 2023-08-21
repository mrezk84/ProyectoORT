package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertEtapa = `
		INSERT INTO ETAPA (nombre, fechas)
		VALUES (?);`

	qryGetEtapabyName = `
		SELECT
			ID,
			nombre
		FROM ETAPAS
		WHERE nombre = ?;`

	qryGetEtapabyId = `
		SELECT
			ID,
			nombre,
			fecha_inicio AS fechaInicio,
		FROM ETAPAS
		WHERE nombre = ?;`
)

func (r *repo) SaveEtapa(ctx context.Context, nombre string) error {
	_, err := r.db.ExecContext(ctx, qryInsertEtapa, nombre)
	return err
}

func (r *repo) GetEtapaById(ctx context.Context, id int64) (*entity.Etapa, error) {
	e := &entity.Etapa{}
	err := r.db.GetContext(ctx, e, qryGetEtapabyName, id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repo) GetEtapaByName(ctx context.Context, nombre string) (*entity.Etapa, error) {
	e := &entity.Etapa{}
	err := r.db.GetContext(ctx, e, qryGetEtapabyName, nombre)
	if err != nil {
		return nil, err
	}

	return e, nil
}

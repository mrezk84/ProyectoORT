package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertEtapa = `
		INSERT INTO ETAPA (nombre)
		VALUES (?);`

	qryGetEtapabyName = `
		SELECT
			ID,
			nombre
		FROM ETAPAS
		WHERE nombre = ?;`
)

func (r *repo) SaveEtapa(ctx context.Context, nombre string) error {
	_, err := r.db.ExecContext(ctx, qryInsertEtapa, nombre)
	return err
}

func (r *repo) GetEtapabyName(ctx context.Context, nombre string) (*entity.Etapa, error) {
	e := &entity.Etapa{}
	err := r.db.GetContext(ctx, e, qryGetEtapabyName, nombre)
	if err != nil {
		return nil, err
	}

	return e, nil
}

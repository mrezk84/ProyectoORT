package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertEtapa = `
		INSERT INTO Etapa (nombre)
		VALUES (?);`

	qryGetEtapabyName = `
		SELECT
			ID,
			Nombre
		FROM ETAPAS
		WHERE Nombre = ?;`
)

func (r *repo) SaveEtapa(ctx context.Context, nombre string) error {
	_, err := r.db.ExecContext(ctx, qryInsertEtapa, nombre)
	return err
}

func (r *repo) GetEtapabyName(ctx context.Context, name string) (*entity.Etapa, error) {
	e := &entity.Etapa{}
	err := r.db.GetContext(ctx, e, qryGetEtapabyName, name)
	if err != nil {
		return nil, err
	}

	return e, nil
}

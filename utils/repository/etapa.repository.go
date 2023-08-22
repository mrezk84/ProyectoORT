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


	qryGetAllEtapas = `
		SELECT id,
		nombre,
		informacion,
		version,
		fecha
		FROM FORMULARIO;`
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
func (r *repo) GetEtapa(ctx context.Context) ([]entity.Etapa, error) {
	ff := []entity.Etapa{}

	err := r.db.SelectContext(ctx, &ff, qryGetAllForms)
	if err != nil {
		return nil, err
	}
	err1 := r.db.SelectContext(ctx, &ff, qryGetFormEtapas)

	if err1 != nil {
		return nil, err
	}
	err2 := r.db.SelectContext(ctx, &ff, qryGetFormUsers)
	if err2 != nil {
		return nil, err
	}

	return ff, nil
}
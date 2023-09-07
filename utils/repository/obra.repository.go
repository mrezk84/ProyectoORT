package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertObra = `
		INSERT INTO OBRA (Nombre)
		VALUES (?);`

	qryGetObrabyID = `
		SELECT
			ID,
			Nombre
		FROM OBRA
		WHERE ID = ?;`

	qryGetObrabyName = `
		SELECT
			ID,
			Nombre
		FROM OBRA
		WHERE Nombre = ?;`

	qryGetObras = `
		SELECT
			ID,
			Nombre
		FROM OBRA`

	qryEliminateObra = `DELETE * FROM OBRA WHERE Nombre = ?;`
)

func (r *repo) SaveObra(ctx context.Context, nombre string) error {
	_, err := r.db.ExecContext(ctx, qryInsertObra, nombre)
	return err
}

func (r *repo) GetObras(ctx context.Context) ([]entity.Obra, error) {
	oo := []entity.Obra{}

	err := r.db.SelectContext(ctx, &oo, qryGetObras)
	if err != nil {
		return nil, err
	}

	return oo, nil
}

func (r *repo) GetObrabyName(ctx context.Context, name string) (*entity.Obra, error) {
	o := &entity.Obra{}
	err := r.db.GetContext(ctx, o, qryGetObrabyName, name)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (r *repo) GetObrabyID(ctx context.Context, obraID int64) (*entity.Obra, error) {
	o := &entity.Obra{}
	err := r.db.GetContext(ctx, o, qryGetObrabyID, obraID)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (r *repo) GetobraP(ctx context.Context, pisoID int64) (*entity.Obra, error) {
	obrap := &entity.Obra{}

	err := r.db.SelectContext(ctx, &obrap, "SELECT obra_id, piso_id FROM OBRA_PISOS WHERE piso_id = ?", pisoID)
	if err != nil {
		return nil, err
	}

	return obrap, nil

}

func (r *repo) DeleteObra(ctx context.Context, nombre string) error {
	_, err := r.db.ExecContext(ctx, qryEliminateObra, nombre)
	return err
}

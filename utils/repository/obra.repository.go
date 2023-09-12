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

	qryGetObraDePiso = `
		SELECT
		o.ID ,
		o.Nombre
		FROM OBRA o 
		inner join OBRA_PISOS op on o.ID = op.obra_id
		WHERE op.piso_id = %v;`

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

	err := r.db.GetContext(ctx, &obrap, qryGetObraDePiso, pisoID)
	if err != nil {
		return nil, err
	}

	return obrap, nil

}

func (r *repo) DeleteObra(ctx context.Context, nombre string) error {
	_, err := r.db.ExecContext(ctx, qryEliminateObra, nombre)
	return err
}

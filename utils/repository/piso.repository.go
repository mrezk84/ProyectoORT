package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertPiso = `
		INSERT INTO PISO (Numero)
		VALUES (?);`

	qryGetPisobyNumber = `
		SELECT
			ID
			Numero
		FROM PISO
		WHERE Numero = ?;`

	qryGetPisos = `
		SELECT
			id,
			numero
		FROM PISO;`

	qryGetPisobyID = `
		SELECT
			ID
			Numero
		FROM PISO
		WHERE ID = ?;`

	qryInsertPisoObra = `
		INSERT INTO OBRA_PISOS (obra_id, piso_id) VALUES (:obra_id, :piso_id);`
)

func (r *repo) SavePiso(ctx context.Context, number int) error {
	_, err := r.db.ExecContext(ctx, qryInsertPiso, number)
	return err
}

func (r *repo) GetPisobyNumber(ctx context.Context, number int) (*entity.Piso, error) {
	p := &entity.Piso{}
	err := r.db.GetContext(ctx, p, qryGetPisobyNumber, number)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *repo) GetPisobyID(ctx context.Context, ID int64) (*entity.Piso, error) {
	p := &entity.Piso{}
	err := r.db.GetContext(ctx, p, qryGetPisobyID, ID)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *repo) GetPisos(ctx context.Context) ([]entity.Piso, error) {
	p := []entity.Piso{}

	err := r.db.GetContext(ctx, &p, qryGetPisos)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *repo) GetObraPisos(ctx context.Context, obraID int64) ([]entity.PisoObra, error) {
	obrap := []entity.PisoObra{}

	err := r.db.SelectContext(ctx, &obrap, "SELECT obra_id, piso_id FROM OBRA_PISOS WHERE obra_id = ?", obraID)
	if err != nil {
		return nil, err
	}

	return obrap, nil

}

func (r *repo) SaveObraPiso(ctx context.Context, obraID, pisoID int64) error {
	data := entity.PisoObra{
		ObraID: obraID,
		PisoID: pisoID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertPisoObra, data)
	return err
}

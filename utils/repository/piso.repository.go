package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertPiso = `
		INSERT INTO Piso (number)
		VALUES ();`

	qryGetPisobyNumber = `
		SELECT
			ID
			numero
		FROM PISOS
		WHERE numero = ?;`

	qryInsertPisoObra = `
		INSERT INTO OBRA_PISOS (obra_id, piso_id) VALUES (:obra_id, :piso_id);`
)

func (r *repo) SavePiso(ctx context.Context, number int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertPiso, number)
	return err
}

func (r *repo) GetPisobyNumber(ctx context.Context, number int64) (*entity.Piso, error) {
	p := &entity.Piso{}
	err := r.db.GetContext(ctx, p, qryGetPisobyNumber)
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

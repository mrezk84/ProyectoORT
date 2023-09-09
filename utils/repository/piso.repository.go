package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertPiso = `
		INSERT INTO PISO (id, numero)
		VALUES (?,?);`

	qryGetPisobyNumber = `
		SELECT
			id,
			numero,
		FROM PISO
		WHERE numero = ?;`

	qryGetPisobyID = `
		SELECT
			id,
			numero,
		FROM PISO
		WHERE numero = ?;`

	qryGetPisos = `
		SELECT 
		 id,
		 numero
		FROM PISO;`

	qryInsertPisoObra = `
		INSERT INTO OBRA_PISOS (obra_id, piso_id) VALUES (:obra_id, :piso_id);`
)

func (r *repo) SavePiso(ctx context.Context, id, numero int) error {
	_, err := r.db.ExecContext(ctx, qryInsertPiso, id, numero)
	return err
}

func (r *repo) GetPisobyNumber(ctx context.Context, numero int) (*entity.Piso, error) {
	p := &entity.Piso{}
	err := r.db.GetContext(ctx, p, qryGetPisobyNumber, numero)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *repo) GetPisobyID(ctx context.Context, pisoID int64) (*entity.Piso, error) {
	o := &entity.Piso{}
	err := r.db.GetContext(ctx, o, qryGetObrabyID, pisoID)
	if err != nil {
		return nil, err
	}

	return o, nil
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
func (r *repo) GetPisos(ctx context.Context) ([]entity.Piso, error) {
	p := []entity.Piso{}

	err := r.db.GetContext(ctx, &p, qryGetPisos)
	if err != nil {
		return nil, err
	}

	return p, nil
}

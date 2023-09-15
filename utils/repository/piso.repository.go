package repository

import (
	"context"
	"fmt"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"

	"github.com/labstack/gommon/log"
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
			id,
			numero
		FROM PISO
		WHERE id = ?;`

	qryInsertPisoObra = `
		INSERT INTO OBRA_PISOS (obra_id, piso_id) VALUES (:obra_id, :piso_id);`

	qryUpdatePiso = `
		update PISO
	set numero = '%v'
	where id = %v	
	`

	qryDeletePisosObra = `
		DELETE FROM OBRA_PISOS where piso_id = ?`

	qryDeletePiso = `
		DELETE FROM PISO where id = ?`
)

func (r *repo) SavePiso(ctx context.Context, number int) (models.Piso, error) {
	result, _ := r.db.ExecContext(ctx, qryInsertPiso, number)
	id, err := result.LastInsertId()
	return models.Piso{
		ID: int(id),
	}, err
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

	err := r.db.SelectContext(ctx, &p, qryGetPisos)
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

func (r *repo) UpdatePiso(ctx context.Context, pisoID int64, numero int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Println(err)
		log.Error(err.Error())
		return err
	}
	_, err = tx.ExecContext(ctx, fmt.Sprintf(qryUpdatePiso, numero, pisoID))
	if err != nil {
		fmt.Println(err)
		fmt.Println("qdas")
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (r *repo) DeletePiso(ctx context.Context, pisoID int64) error {
	_, err := r.db.ExecContext(ctx, qryDeletePisosObra, pisoID)
	if err != nil {
		return err
	}
	_, err2 := r.db.ExecContext(ctx, qryDeletePiso, pisoID)
	return err2
}

package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertCheck = `
		INSERT INTO Check (estado, fecha_control, observaciones, version)
		VALUES (?,?,?,?);`

	qryGetCheckByVersion = `
		SELECT
			ID
			estado
			fecha_control
			observaciones
			version
		FROM CHECKS
		WHERE version = ?;`

	qryInsertCheckForm = `
		INSERT INTO CHECK_FORMULARIO (check_id, formulario_id) VALUES (:check_id, :formulario_id);`
)

func (r *repo) SaveCheck(ctx context.Context, estado, fecha, observaciones string, version int) error {
	_, err := r.db.ExecContext(ctx, qryInsertCheck, estado, fecha, observaciones, version)
	return err
}

func (r *repo) GetCheckByVersion(ctx context.Context, version int) (*entity.Check, error) {
	c := &entity.Check{}
	err := r.db.GetContext(ctx, c, qryGetCheckByVersion, version)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *repo) GetCheckForm(ctx context.Context, FormularioID int64) ([]entity.CheckFormulario, error) {
	checkf := []entity.CheckFormulario{}

	err := r.db.SelectContext(ctx, &checkf, "SELECT check_id, formulario_id FROM CHECK_FORMULARIO WHERE formulario_id = ?", FormularioID)
	if err != nil {
		return nil, err
	}

	return checkf, nil

}

func (r *repo) SaveCheckForm(ctx context.Context, checkID, formularioID int64) error {
	data := entity.CheckFormulario{
		CheckID:      checkID,
		FormularioID: formularioID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertCheckForm, data)
	return err
}

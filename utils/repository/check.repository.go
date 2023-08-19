package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertCheck = `
		INSERT INTO CHECK (estado, observaciones, version, fecha_control)
		VALUES (?,?,?,?);`

	qryGetCheckByVersion = `
		SELECT
			id,
			estado,
			observaciones,
			version,
			fecha_control
		FROM CHECK
		WHERE version = ?;`

	qryInsertCheckForm = `
		INSERT INTO CHECK_FORMULARIO (check_id, formulario_id) VALUES (:check_id, :formulario_id);`

	qryGetChecks = `
		SELECT
			id,
			estado,
			observaciones,
			version,
			fecha_control
		FROM CHECK`
)

func (r *repo) SaveCheck(ctx context.Context, estado, observaciones string, version int, fecha string) error {
	_, err := r.db.ExecContext(ctx, qryInsertCheck, estado, observaciones, version, fecha)

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
func (r *repo) GetChecks(ctx context.Context) ([]entity.Check, error) {
	cc := []entity.Check{}
	err := r.db.SelectContext(ctx, &cc, qryGetChecks)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

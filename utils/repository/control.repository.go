package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertControl = `
		INSERT INTO CONTROL(descripcion, tipo)
		VALUES (?, ?);`

	qryGetContBydescripcion = `
		SELECT
		id,
			descripcion,
			tipo,
		FROM CONTROL
		WHERE descripcion = ?;`

	qryGetContByid = `
		SELECT
		id,
			descripcion,
			tipo,
		FROM CONTROL
		WHERE id = ?;`

	qryGetAllControls = `
		SELECT 
		id,
		descripcion,
		tipo
		FROM CONTROL;`

	qryGetControlsByForm = `
		SELECT 
		control_id,
		formulario_id
		FROM CONTROL_FORMULARIO
		WHERE formulario_id = ?;`

	qryInsertControlForm = `INSERT INTO CONTROL_FORMULARIO (control_id, formulario_id) VALUES (:control_id, :formulario_id);`
)

func (r *repo) SaveControl(ctx context.Context, descripcion, tipo string) error {
	_, err := r.db.ExecContext(ctx, qryInsertControl, descripcion, tipo)
	return err
}
func (r *repo) GetControls(ctx context.Context) ([]entity.Control, error) {
	cc := []entity.Control{}

	err := r.db.SelectContext(ctx, &cc, qryGetAllControls)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (r *repo) GetControlsByForm(ctx context.Context, formID int64) ([]entity.ControlForm, error) {
	cc := []entity.ControlForm{}

	err := r.db.SelectContext(ctx, &cc, qryGetControlsByForm)
	if err != nil {
		return nil, err
	}

	return cc, nil
}
func (r *repo) GetConByDes(ctx context.Context, des string) (*entity.Control, error) {
	c := &entity.Control{}
	err := r.db.GetContext(ctx, c, qryGetContBydescripcion, des)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *repo) GetConByid(ctx context.Context, id int) (*entity.Control, error) {
	c := &entity.Control{}
	err := r.db.GetContext(ctx, c, qryGetContByid, id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *repo) GetControlForm(ctx context.Context, controlID int64) ([]entity.ControlForm, error) {
	controlf := []entity.ControlForm{}

	err := r.db.SelectContext(ctx, &controlf, "SELECT control_id, formulario_id FROM CONTROL_FORMULARIO WHERE control_id = ?", controlID)
	if err != nil {
		return nil, err
	}

	return controlf, nil

}

func (r *repo) SaveControlForm(ctx context.Context, controlID, formularioID int64) error {
	data := entity.ControlForm{
		ControlID:    controlID,
		FormularioID: formularioID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertControlForm, data)
	return err
}

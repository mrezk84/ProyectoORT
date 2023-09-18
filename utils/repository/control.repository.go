package repository

import (
	"context"
	"fmt"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"
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

	qryGetAllControls = `
		SELECT 
		id,
		descripcion,
		tipo
		FROM CONTROL;`

	qryGetControlsByForm = `
		SELECT
		c.id,
		c.descripcion,
		c.tipo
		FROM CONTROL c
		inner join CONTROL_FORMULARIO CF on c.id = CF.control_id
		WHERE CF.formulario_id = %v;`

	qryGetControlsSinForm = `
		SELECT
		c.id,
		c.descripcion,
		c.tipo
		FROM CONTROL c
		inner join CONTROL_FORMULARIO CF on c.id != CF.control_id`

	qryInsertControlForm = `INSERT INTO CONTROL_FORMULARIO (control_id, formulario_id) VALUES (:control_id, :formulario_id);`

	qryDeleteControlForm = `Delete from CONTROL_FORMULARIO where control_id = %v and formulario_id = %v`
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

func (r *repo) GetControlsByForm(ctx context.Context, formID int64) ([]entity.Control, error) {
	cc := []entity.Control{}

	err := r.db.SelectContext(ctx, &cc, fmt.Sprintf(qryGetControlsByForm, formID))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cc, nil
}

func (r *repo) GetControlsSinForm(ctx context.Context) ([]entity.Control, error) {
	cc := []entity.Control{}

	err := r.db.SelectContext(ctx, &cc, fmt.Sprintf(qryGetControlsSinForm))
	if err != nil {
		fmt.Println(err)
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

func (r *repo) GetControlForm(ctx context.Context, controlID int64) ([]entity.ControlForm, error) {
	controlf := []entity.ControlForm{}

	err := r.db.SelectContext(ctx, &controlf, "SELECT control_id, formulario_id FROM CONTROL_FORMULARIO WHERE control_id = ?", controlID)
	if err != nil {
		return nil, err
	}

	return controlf, nil

}

func (r *repo) DeleteControlForm(ctx context.Context, controlID, formularioID int64) error {
	tx, err := r.db.Beginx()
	data := entity.ControlForm{
		ControlID:    controlID,
		FormularioID: formularioID,
	}
	_, err = tx.ExecContext(ctx, fmt.Sprintf(qryDeleteControlForm, data.ControlID, data.FormularioID))
	if err != nil {
		tx.Rollback()
		return err
	}
	documents, err := r.GetWipOrTodoDocumentsByFormID(ctx, formularioID)
	err = r.DeleteChecks(ctx, formularioID, documents, int(controlID))
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (r *repo) SaveControlForm(ctx context.Context, controlID, formularioID int64) error {
	tx, err := r.db.Beginx()
	data := entity.ControlForm{
		ControlID:    controlID,
		FormularioID: formularioID,
	}
	_, err = tx.NamedExecContext(ctx, qryInsertControlForm, data)
	if err != nil {
		tx.Rollback()
		return err
	}
	documents, err := r.GetWipOrTodoDocumentsByFormID(ctx, formularioID)
	for _, d := range documents {
		err = r.InsertChecks(ctx, formularioID, d.ID, []models.Control{{ID: int(controlID)}})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return err
}

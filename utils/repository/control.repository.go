package repository

import (
	"context"
	"fmt"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"

	"github.com/labstack/gommon/log"
)

const (
	qryInsertControl = `
		INSERT INTO CONTROL(descripcion, tipo)
		VALUES (?, ?);`

	qryGetContBydescripcionandTipo = `
		SELECT
		id,
		descripcion,
		tipo
		FROM CONTROL
		WHERE descripcion = %v and tipo = %v;`

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

	qryUpdateControl = `
		update CONTROL
	set descripcion = '%v',
	tipo = '%v'
	where id = %v	
	`

	qryInsertControlForm = `INSERT INTO CONTROL_FORMULARIO (control_id, formulario_id) VALUES (:control_id, :formulario_id);`

	qryDeleteControlForm = `Delete from CONTROL_FORMULARIO where control_id = %v and formulario_id = %v`

	qryDeleteControlFormularios = `
		DELETE FROM CONTROL_FORMULARIO where control_id = ?`

	qryDeleteControlChecks = `
		DELETE FROM CHECKS where control_id = ?`

	qryDeleteControl = `
		DELETE FROM CONTROL where id = ?`
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

func (r *repo) UpdateControl(ctx context.Context, ControlID int64, descripcion, tipo string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Println(err)
		log.Error(err.Error())
		return err
	}
	_, err = tx.ExecContext(ctx, fmt.Sprintf(qryUpdateControl, descripcion, tipo, ControlID))
	if err != nil {
		fmt.Println(err)
		fmt.Println("qdas")
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (r *repo) GetConByDesAndTipo(ctx context.Context, des, tipo string) (*entity.Control, error) {
	c := &entity.Control{}
	err := r.db.GetContext(ctx, c, qryGetContBydescripcionandTipo, des, tipo)
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

func (r *repo) DeleteControl(ctx context.Context, controlID int64) error {
	_, err := r.db.ExecContext(ctx, qryDeleteControlFormularios, controlID)
	if err != nil {
		return err
	}
	_, err2 := r.db.ExecContext(ctx, qryDeleteControlChecks, controlID)
	if err2 != nil {
		return err2
	}
	_, err3 := r.db.ExecContext(ctx, qryDeleteControl, controlID)
	return err3
}

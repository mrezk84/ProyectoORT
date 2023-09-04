package repository

import (
	"context"
	"github.com/labstack/gommon/log"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"
	"time"
)

const (
	qryInsertCheck = `
		INSERT INTO CHECKS (estado, observaciones, version, fecha_control)
		VALUES (?,?,?,?);`

	qryCreateCheck = `
		INSERT INTO CHECKS (estado, observaciones, version, fecha_control,document_id,formulario_id,control_id)
		VALUES (?,?,?,?,?,?,?);`

	qryGetCheckByVersion = `
		SELECT
			id
			estado
			observaciones
			version
			fecha_control
		FROM CHECKS
		WHERE version = ?;`

	qryGetCheckSByDocument = `
		SELECT
			id
			estado
			observaciones
			version
			fecha_control
		FROM CHECKS
		inner join document d on CHECK.document_id = d.id
where d.id = ?;`

	qryInsertCheckForm = `
		INSERT INTO CHECK_FORMULARIO (check_id, formulario_id) VALUES (:check_id, :formulario_id);`
)

func (r *repo) InsertChecks(ctx context.Context, formularioID int64, documentID int64, controles []models.Control) error {
	tx, err := r.db.Beginx()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	for _, c := range controles {
		_, err = tx.ExecContext(ctx, qryCreateCheck, "", "", 1, nil, documentID, formularioID, c.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return err
}

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

func (r *repo) GetDocumentsChecks(ctx context.Context, documents []models.Document) ([]models.Check, error) {
	var checks []models.Check
	for _, d := range documents {
		c := &entity.Check{}
		err := r.db.GetContext(ctx, c, qryGetCheckByVersion, d.ID)
		if err != nil {
			return nil, err
		}
		fechaControl, err := time.Parse("2006-01-02", c.FechaControl)
		if err != nil {
			return nil, err
		}
		checks = append(checks, models.Check{
			ID:           c.ID,
			Estado:       c.Estado,
			FechaControl: &fechaControl,
		})
	}

	return checks, nil
}

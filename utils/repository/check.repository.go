package repository

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"
	"time"
)

const (
	qryInsertCheck = `
		INSERT INTO CHECKS (estado, observaciones, version, fecha_control)
		VALUES (?,?,?,?);`

	qryDeleteChecks = `
		delete from CHECKS c where
   		c.document_id in %v
    	and c.control_id = %v
		and formulario_id = %v
 `

	qryCreateCheck = `
		INSERT INTO CHECKS (estado, observaciones, version, fecha_control,document_id,formulario_id,control_id)
		VALUES ('','',0,null,%v,%v,%v);`

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
			c.id,
			c.estado,
			c.observaciones,
			c.version,
			c.fecha_control
		FROM CHECKS c
		inner join document d on c.document_id = d.id
where d.id = ?;`
	qryUpdateCheck = `
	update CHECKS 
set estado = '%v',
observaciones = '%v',
fecha_control = '%v'
where id = %v	
`
	qryInsertChecksHistoric = `
	insert into checks_historico (id,estado,observaciones,fecha_control) values(%v,'%s','%s','%s')	

`
	qryDeleteChecksHistoric = `
	insert into checks_historico (id,estado,observaciones,fecha_control) values(%v,'%s','%s','%s')	`
	qryInsertCheckForm = `
		INSERT INTO CHECK_FORMULARIO (check_id, formulario_id) VALUES (:check_id, :formulario_id);`
)

func (r *repo) InsertChecks(ctx context.Context, formularioID int64, documentID int64, controles []models.Control) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Println(err)
		log.Error(err.Error())
		return err
	}
	for _, c := range controles {
		fmt.Println(fmt.Sprintf(qryCreateCheck, documentID, formularioID, c.ID))
		result, err := tx.ExecContext(ctx, fmt.Sprintf(qryCreateCheck, documentID, formularioID, c.ID))
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
		lastInsertId, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
		_, err = tx.ExecContext(ctx, fmt.Sprintf(qryInsertChecksHistoric, lastInsertId, "", "", ""))
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return err
}
func (r *repo) DeleteChecks(ctx context.Context, formularioID int64, documents []models.Document, control int) error {
	d := "("
	for _, doc := range documents {
		d += fmt.Sprint(doc.ID)
	}
	d += ")"

	fmt.Println(fmt.Sprintf(qryDeleteChecks, d, control, formularioID))
	_, err := r.db.ExecContext(ctx, fmt.Sprintf(qryDeleteChecks, d, control, formularioID))
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

func (r *repo) GetDocumentChecks(ctx context.Context, documentID int64) ([]models.Check, error) {
	var checks []entity.Check
	err := r.db.SelectContext(ctx, &checks, qryGetCheckSByDocument, documentID)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	var checksResponse []models.Check
	for _, check := range checks {
		checksResponse = append(checksResponse, models.Check{
			ID:            check.ID,
			Estado:        check.Estado,
			FechaControl:  check.FechaControl,
			Responsable:   models.Usuario{},
			Control:       models.Control{},
			Document:      models.Document{},
			Observaciones: check.Observaciones,
			Version:       check.Version,
		})
	}
	return checksResponse, nil
}

func (r *repo) UpdateCheck(ctx context.Context, checkID int64, estado, observaciones string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Println(err)
		log.Error(err.Error())
		return err
	}
	_, err = tx.ExecContext(ctx, fmt.Sprintf(qryUpdateCheck, estado, observaciones, time.Now().UTC(), checkID))
	if err != nil {
		fmt.Println(err)
		fmt.Println("qdas")
		tx.Rollback()
		return err
	}
	fmt.Println(fmt.Sprintf(qryInsertChecksHistoric, checkID, estado, observaciones, time.Now().UTC()))
	_, err = tx.ExecContext(ctx, fmt.Sprintf(qryInsertChecksHistoric, checkID, estado, observaciones, time.Now().UTC()))
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

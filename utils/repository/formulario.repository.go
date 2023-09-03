package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertFrom = `
		INSERT INTO FORMULARIO (informacion,version, nombre,control_id, usuario_id)
		VALUES (?, ?, ?, ?, ?, ?);`

	qryGetFormByDate = `
		SELECT
			id,
			nombre,
			informacion,
			version,
			fecha
		FROM FORMULARIO
		WHERE fecha = ?;`

	qryGetFormByVersion = `
		SELECT
			id,
			nombre,
			informacion,
			version,
			fecha
		FROM FORMULARIO
		WHERE version = ?;`

	qryGetAllForms = `
		SELECT id,
		nombre,
		informacion,
		version,
		fecha
		FROM FORMULARIO;`

	qryGetFormById = `
		SELECT
			id,
			nombre,
			informacion,
			version,
			fecha
		FROM FORMULARIO
		WHERE id = ?;`

	qryGetFormByName = `
	SELECT
		id,
		nombre,
		informacion,
		version,
		fecha
	FROM FORMULARIO
	WHERE name = ?;`

	qryGetFormControles = `
		SELECT f.id,f.nombre,f.informacion,f.fecha, c.tipo as contoles
		FROM FORMULARIO f INNER JOIN CONTROL c
		ON f.id=c.id
		WHERE f.id=c.id`

	qryGetFormUsers = `
		SELECT f.id,f.nombre,f.informacion,f.fecha, u.username as usuario
		FROM FORMULARIO f INNER JOIN USUARIOS u
		ON f.id=u.id
		WHERE f.id=u.id`
)

func (r *repo) SaveFrom(ctx context.Context, informacion string, version int, nombre string, control_id int, usuario_id int) error {
	_, err := r.db.ExecContext(ctx, qryInsertFrom, informacion, version, nombre, control_id, usuario_id)
	return err
}

func (r *repo) GetFormByDate(ctx context.Context, fecha string) (*entity.Formulario, error) {
	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetFormByDate, fecha)
	if err != nil {
		return nil, err
	}

	return f, nil
}
func (r *repo) GetForms(ctx context.Context) ([]entity.Formulario, error) {
	ff := []entity.Formulario{}

	err := r.db.SelectContext(ctx, &ff, qryGetAllForms)
	if err != nil {
		return nil, err
	}

	return ff, nil
}

func (r *repo) GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error) {
	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetFormByVersion, version)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *repo) GetFormControles(ctx context.Context) (*entity.Formulario, error) {

	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetFormControles)
	if err != nil {
		return nil, err
	}

	return f, nil

}

func (r *repo) GetFormsById(ctx context.Context, ID int64) (*entity.Formulario, error) {
	f := &entity.Formulario{}

	err := r.db.GetContext(ctx, f, qryGetFormById, ID)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *repo) GetFormByName(ctx context.Context, nombre string) (*entity.Formulario, error) {
	f := &entity.Formulario{}

	err := r.db.GetContext(ctx, f, qryGetFormByName, nombre)
	if err != nil {
		return nil, err
	}

	return f, nil
}
func (r *repo) GetFromUsers(ctx context.Context) (*entity.Formulario, error) {

	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetFormUsers)
	if err != nil {
		return nil, err
	}

	return f, nil

}

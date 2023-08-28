package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertFrom = `
		INSERT INTO  FORMULARIO (nombre,informacion, version, fecha)
		VALUES (?, ?, ?, ?);`

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

	qryGetFormByID = `
		SELECT
			id,
			nombre,
			informacion,
			version,
			fecha
		FROM FORMULARIO
		WHERE id = ?;`

	qryGetNForm = `
		SELECT
			nombre
		FROM FORMULARIO
		WHERE id = ?;`

	qryGetAllForms = `
		SELECT
		id,
		nombre,
		informacion,
		version
		FROM FORMULARIO;`

	qryGetFormCategories = `
		SELECT f.id,f.nombre,f.informacion,f.fecha, c.descripcion as controles
		FROM FORMULARIO f INNER JOIN CONTROLES c
		ON f.id=c.id
		WHERE f.id=c.id`

	qryInsertUserForm = `
		INSERT INTO FORMULARIO_RESPONSABLE (formulario_id, usuario_id) VALUES (:formulario_id, :usuario_id);`
)

func (r *repo) SaveFrom(ctx context.Context, nombre, informacion string, version string, fecha string) error {
	_, err := r.db.ExecContext(ctx, qryInsertFrom, informacion, nombre, version, fecha)
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
func (r *repo) GetFormByID(ctx context.Context, formID int64) (*entity.Formulario, error) {
	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetFormByID, formID)
	if err != nil {
		return nil, err
	}

	return f, nil
}
func (r *repo) GetForm(ctx context.Context) ([]entity.Formulario, error) {
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

func (r *repo) GetFromControles(ctx context.Context, controles string) (*entity.Formulario, error) {

	f := &entity.Formulario{}
	err := r.db.GetContext(ctx, f, qryGetFormCategories, controles)
	if err != nil {
		return nil, err
	}

	return f, nil

}

func (r *repo) GetUsuarioForm(ctx context.Context, usuarioID int64) ([]entity.UsuarioForm, error) {
	usuariosf := []entity.UsuarioForm{}

	err := r.db.SelectContext(ctx, &usuariosf, "SELECT formulario_id, usuario_id FROM FORMULARIO_RESPONSABLE WHERE usuario_id = ?", usuarioID)
	if err != nil {
		return nil, err
	}

	return usuariosf, nil

}

func (r *repo) SaveUserForm(ctx context.Context, formID, usuarioID int64) error {
	data := entity.UsuarioForm{
		FormularioID: formID,
		UsuarioID:    usuarioID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertUserForm, data)
	return err
}

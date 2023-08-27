package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertRol = `
		INSERT INTO ROLES (nombre) VALUES (?);`

	qryGetAllRoles = `
		SELECT
			id,
			nombre
		FROM ROLES;`

	qryGetRolByName = `
		SELECT
			id,
			nombre,
		FROM ROLES
		WHERE nombre = ?;`

	qryGetRolById = `
		SELECT
			id,
			nombre
		FROM ROLES
		WHERE id = ?;`
)

func (r *repo) SaveRole(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, qryInsertRol, id)

	return err

}
func (r *repo) GetRolByName(ctx context.Context, nombre string) (*entity.Rol, error) {
	ro := &entity.Rol{}
	err := r.db.GetContext(ctx, ro, qryGetRolByName, nombre)
	if err != nil {
		return nil, err
	}

	return ro, nil
}
func (r *repo) GetAllRoles(ctx context.Context) ([]entity.Rol, error) {
	ro := []entity.Rol{}

	err := r.db.SelectContext(ctx, &ro, qryGetAllRoles)
	if err != nil {
		return nil, err
	}

	return ro, nil
}
func (r *repo) GetRolById(ctx context.Context, id int) (*entity.Rol, error) {
	ro := &entity.Rol{}
	err := r.db.GetContext(ctx, ro, qryGetRolById, id)
	if err != nil {
		return nil, err
	}

	return ro, nil
}

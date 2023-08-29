package repository

import (
	"context"

	"proyectoort/utils/entity"
)

const (
	qryInsertUser = `
		INSERT INTO USUARIOS (email, username, password)
		VALUES (?, ?, ?);`

	qryGetUserByEmail = `
		SELECT
			id,
			email,
			username,
			password
		FROM USUARIOS
		WHERE email = ?;`

	qryGetUserByID = `
		SELECT
			id,
			email,
			username,
			password
		FROM USUARIOS
		WHERE id = ?;`

	qryAllGetUsers = `
		SELECT
			id,
			email,
			username,
			password
		FROM USUARIOS;`

	qryInsertUserRole = `
		INSERT INTO USUARIOS_ROLES (usuario_id, rol_id) VALUES (:usuario_id, :rol_id);`

	qryRemoveUserRole = `
		DELETE FROM USUARIOS_ROLES where usuario_id = :usuario_id and rol_id = :rol_id;`
)

func (r *repo) SaveUser(ctx context.Context, email, username, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, username, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.Usuario, error) {
	u := &entity.Usuario{}
	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) GetUserById(ctx context.Context, id int64) (*entity.Usuario, error) {
	u := &entity.Usuario{}
	err := r.db.GetContext(ctx, u, qryGetUserByID, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UsarioRol{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)
	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UsarioRol{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryRemoveUserRole, data)

	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]entity.UsarioRol, error) {
	roles := []entity.UsarioRol{}

	err := r.db.SelectContext(ctx, &roles, "SELECT usuario_id, rol_id FROM USUARIOS_ROLES WHERE usuario_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return roles, nil

}
func (r *repo) GetUsers(ctx context.Context) ([]entity.Usuario, error) {
	us := []entity.Usuario{}

	err := r.db.SelectContext(ctx, &us, qryAllGetUsers)
	if err != nil {
		return nil, err
	}

	return us, nil
}

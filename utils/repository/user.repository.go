package repository

import (
	"context"

	"proyectoort/utils/entity"
)

const (
	qryInsertUser = `
		INSERT int USUARIOS (email, name, password)
		VALUES (?, ?, ?);`

	qryGetUserByEmail = `
		SELECT
			id,
			email,
			name,
			password
		FROM USUARIOS
		WHERE email = ?;`

	qryInsertUserRole = `
		INSERT INTO USUARIOS_ROLES (user_id, role_id) VALUES (:user_id, :role_id);`

	qryRemoveUserRole = `
		DELETE FROM USUARIOS_ROLES where user_id = :user_id and role_id = :role_id;`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
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

	err := r.db.SelectContext(ctx, &roles, "SELECT user_id, role_id FROM USER_ROLES WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

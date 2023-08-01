package service

import (
	"context"
	"errors"

	"proyectoort/encryption"
	"proyectoort/utils/models"
)

var (
	ErrUserAlreadyExists  = errors.New("El usuario ya existe")
	ErrInvalidCredentials = errors.New("Creedenciales Inválidas")
	ErrRoleAlreadyAdded   = errors.New("El usuario ya cuenta con el rol asignado")
	ErrRoleNotFound       = errors.New("Error al asignar rol")
)

func (s *serv) RegisterUser(ctx context.Context, email, username, password string) error {

	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)
	return s.repo.SaveUser(ctx, email, username, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.Usuario, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.Usuario{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleID == roleID {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	roleFound := false
	for _, r := range roles {
		if r.RoleID == roleID {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, userID, roleID)
}
func (s *serv) GetUsers(ctx context.Context) ([]models.Usuario, error) {
	us, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	usuarios := []models.Usuario{}
	for _, u := range us {
		usuarios = append(usuarios, models.Usuario{
			ID:    u.ID,
			Email: u.Email,
			Name:  u.Name,
		})
	}

	return usuarios, nil
}

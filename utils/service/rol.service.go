package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrRolAlreadyExists = errors.New("El rol ya existe")
)

func (s *serv) RegisterRol(ctx context.Context, nombre string) error {

	r, _ := s.repo.GetRolByName(ctx, nombre)
	if r != nil {
		return ErrRolAlreadyExists
	}

	return s.repo.SaveRole(ctx, nombre)
}
func (s *serv) GetAllRoles(ctx context.Context) ([]models.Rol, error) {
	ro, err := s.repo.GetAllRoles(ctx)
	if err != nil {
		return nil, err
	}
	roles := []models.Rol{}
	for _, r := range ro {
		roles = append(roles, models.Rol{
			ID:     int64(r.ID),
			Nombre: r.Nombre,
		})

	}
	return roles, nil
}

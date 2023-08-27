package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrRolAlreadyExists = errors.New("El rol ya existe")
)

func (s *serv) RegisterRol(ctx context.Context, id int) error {

	r, _ := s.repo.GetRolById(ctx, id)
	if r != nil {
		return ErrRolAlreadyExists
	}

	return s.repo.SaveRole(ctx, id)
}
func (s *serv) GetAllRoles(ctx context.Context) ([]models.Rol, error) {
	ro, err := s.repo.GetAllRoles(ctx)
	if err != nil {
		return nil, err
	}
	roles := []models.Rol{}
	for _, r := range ro {
		roles = append(roles, models.Rol{
			ID:     r.ID,
			Nombre: r.Nombre,
		})

	}
	return roles, nil
}

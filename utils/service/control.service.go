package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrContAlreadyExists = errors.New("El control ya existe")
	ErrInvalidCont       = errors.New("Control Inválido")
	ErrContAlreadyAdded  = errors.New("Ya se cuenta con el control asignado")
	ErrContNotFound      = errors.New("Error al asignar el control")
)

func (s *serv) RegisterControl(ctx context.Context, descripcion, tipo string) error {

	c, _ := s.repo.GetConByDes(ctx, descripcion)
	if c != nil {
		return ErrFormAlreadyExists
	}

	return s.repo.SaveControl(ctx, descripcion, tipo)
}
func (s *serv) GetControls(ctx context.Context) ([]models.Control, error) {
	cc, err := s.repo.GetControls(ctx)
	if err != nil {
		return nil, err
	}

	controles := []models.Control{}

	for _, c := range cc {
		controles = append(controles, models.Control{
			ID:          c.ID,
			Descripcion: c.Descripcion,
			Tipo:        c.Tipo,
		})

	}

	return controles, nil
}

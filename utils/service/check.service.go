package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrCheckAlreadyExists     = errors.New("El Check ya existe")
	ErrCheckFormAlreadyExists = errors.New("La Conexion ya Existe")
)

func (s *serv) RegisterCheck(ctx context.Context, estado string, fecha string, observaciones string, version int) error {

	f, _ := s.repo.GetCheckByVersion(ctx, version)
	if f != nil {
		return ErrCheckAlreadyExists
	}

	return s.repo.SaveCheck(ctx, estado, observaciones, version, fecha)
}

func (s *serv) AddCheckForm(ctx context.Context, checkID, formularioID int64) error {

	checksf, err := s.repo.GetCheckForm(ctx, formularioID)
	if err != nil {
		return err
	}

	for _, r := range checksf {
		if r.CheckID == checkID {
			return ErrCheckFormAlreadyExists
		}
	}

	return s.repo.SaveCheckForm(ctx, checkID, formularioID)
}

func (s *serv) UpdateCheck(ctx context.Context, checkID int64, estado, observaciones string) error {
	return s.repo.UpdateCheck(ctx, checkID, estado, observaciones)
}

func (s *serv) GetDocumentChecks(ctx context.Context, documentID int64) ([]models.Check, error) {
	return s.repo.GetDocumentChecks(ctx, documentID)
}

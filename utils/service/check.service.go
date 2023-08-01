package service

import (
	"context"
	"errors"
)

var (
	ErrCheckAlreadyExists = errors.New("El Check ya existe")
)

func (s *serv) RegisterCheck(ctx context.Context, estado string, fecha string, observaciones string, version int) error {

	f, _ := s.repo.GetCheckByVersion(ctx, version)
	if f != nil {
		return ErrCheckAlreadyExists
	}

	return s.repo.SaveCheck(ctx, estado, fecha, observaciones, version)
}

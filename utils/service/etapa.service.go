package service

import (
	"context"
	"errors"
)

var (
	ErrEtapaAlreadyExists = errors.New("La Etapa ya existe")
)

func (s *serv) RegisterEtapa(ctx context.Context, nombre string) error {

	e, _ := s.repo.GetEtapaByName(ctx, nombre)
	if e != nil {
		return ErrEtapaAlreadyExists
	}

	return s.repo.SaveEtapa(ctx, nombre)
}

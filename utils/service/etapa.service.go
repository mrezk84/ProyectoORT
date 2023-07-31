package service

import (
	"context"
	"errors"
)

var (
	ErrEtapaAlreadyExists = errors.New("La Etapa ya existe")
)

func (s *serv) RegisterEtapa(ctx context.Context, name string) error {

	e, _ := s.repo.GetEtapabyName(ctx, name)
	if e != nil {
		return ErrEtapaAlreadyExists
	}

	return s.repo.SaveObra(ctx, name)
}

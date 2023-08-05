package service

import (
	"context"
	"errors"
)

var (
	ErrPisoAlreadyExists     = errors.New("El Piso ya Existe")
	ErrPisoObraAlreadyExists = errors.New("La Conexion ya Existe")
	ErrObraDoesNotExists     = errors.New("La Obra no Existe")
)

func (s *serv) RegisterPiso(ctx context.Context, number int64) error {

	p, _ := s.repo.GetPisobyNumber(ctx, number)
	if p != nil {
		return ErrPisoAlreadyExists
	}

	return s.repo.SavePiso(ctx, number)
}

func (s *serv) AddObraPiso(ctx context.Context, obraID, pisoID int64) error {

	pisos, err := s.repo.GetObraPisos(ctx, obraID)
	if err != nil {
		return err
	}

	for _, r := range pisos {
		if r.PisoID == pisoID {
			return ErrPisoObraAlreadyExists
		}
	}

	return s.repo.SaveObraPiso(ctx, obraID, pisoID)
}

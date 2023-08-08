package service

import (
	"context"
	"errors"
)

var (
	ErrObraAlreadyExists = errors.New("La Obra ya existe")
	ErrObraDoesntExists  = errors.New("No existe Obra con ese nombre")
)

func (s *serv) RegisterObra(ctx context.Context, name string) error {

	o, _ := s.repo.GetObrabyName(ctx, name)
	if o != nil {
		return ErrObraAlreadyExists
	}

	return s.repo.SaveObra(ctx, name)
}

func (s *serv) DeleteObra(ctx context.Context, name string) error {

	o, _ := s.repo.GetObrabyName(ctx, name)
	if o != nil {
		return s.repo.DeleteObra(ctx, name)
	}

	return ErrObraDoesNotExists
}

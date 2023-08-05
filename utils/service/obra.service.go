package service

import (
	"context"
	"errors"
)

var (
	ErrObraAlreadyExists = errors.New("La Obra ya existe")
)

func (s *serv) RegisterObra(ctx context.Context, name string) error {

	o, _ := s.repo.GetObrabyName(ctx, name)
	if o != nil {
		return ErrObraAlreadyExists
	}

	return s.repo.SaveObra(ctx, name)
}

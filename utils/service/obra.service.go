package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
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

func (s *serv) GetObras(ctx context.Context) ([]models.Obra, error) {
	oo, err := s.repo.GetObras(ctx)
	if err != nil {
		return nil, err
	}

	obras := []models.Obra{}

	for _, o := range oo {
		obras = append(obras, models.Obra{
			ID:     o.ID,
			Nombre: o.Nombre,
		})

	}

	return obras, nil
}

func (s *serv) GetPisosObra(ctx context.Context, ObraID int64) ([]models.Piso, error) {
	po, err := s.repo.GetObraPisos(ctx, ObraID)
	if err != nil {
		return nil, err
	}

	PisosObra := []models.ObraPiso{}

	for _, p := range po {
		PisosObra = append(PisosObra, models.ObraPiso{
			ObraID: p.ObraID,
			PisoID: p.PisoID,
		})
	}

	Pisos := []models.Piso{}

	for _, P := range PisosObra {
		piso, _ := s.repo.GetPisobyID(ctx, P.PisoID)
		Pisos = append(Pisos, models.Piso{
			ID:     piso.ID,
			Numero: piso.Number,
		})
	}

	return Pisos, nil
}

func (s *serv) DeleteObra(ctx context.Context, name string) error {

	o, _ := s.repo.GetObrabyName(ctx, name)
	if o != nil {
		return s.repo.DeleteObra(ctx, name)
	}

	return ErrObraDoesNotExists
}

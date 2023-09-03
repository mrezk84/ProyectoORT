package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrPisoAlreadyExists     = errors.New("El Piso ya Existe")
	ErrPisoObraAlreadyExists = errors.New("La Conexion ya Existe")
	ErrObraDoesNotExists     = errors.New("La Obra no Existe")
)

func (s *serv) RegisterPiso(ctx context.Context, id, numero int) error {

	p, _ := s.repo.GetPisobyNumber(ctx, numero)
	if p != nil {
		return ErrPisoAlreadyExists
	}

	return s.repo.SavePiso(ctx, id, numero)
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
func (s *serv) GetPisos(ctx context.Context) ([]models.Piso, error) {
	pp, err := s.repo.GetPisos(ctx)
	if err != nil {
		return nil, err
	}

	pisos := []models.Piso{}

	for _, p := range pp {
		pisos = append(pisos, models.Piso{
			ID:     p.ID,
			Numero: p.Numero,
		})

	}

	return pisos, nil
}
func (s *serv) GetPhoto(ctx context.Context, id int) (*models.Foto, error) {
	p, err := s.repo.GetPhotoById(ctx, id)
	if err != nil {
		return nil, err
	}

	foto := &models.Foto{
		ID:           p.ID,
		Nombre:       p.Nombre,
		Notas:        p.Notas,
		FormularioID: p.FormularioID,
	}

	return foto, nil
}

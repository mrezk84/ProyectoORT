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

func (s *serv) RegisterPiso(ctx context.Context, number int, obraId int64) (models.Piso, error) {

	op, _ := s.repo.GetObraPisos(ctx, obraId)
	existe := false
	for _, p := range op {

		piso := models.Piso{
			ID:     p.ID,
			Numero: p.Numero,
		}

		if piso.Numero == number {
			existe = true
		}
	}

	if !existe {
		return s.repo.SavePiso(ctx, number)
	}
	return models.Piso{}, ErrPisoAlreadyExists
}

func (s *serv) GetPisos(ctx context.Context) ([]models.Piso, error) {
	pp, err := s.repo.GetPisos(ctx)
	if err != nil {
		return nil, err
	}

	pisos := []models.Piso{}

	for _, p := range pp {

		// obra, _ := s.repo.GetobraP(ctx, int64(p.ID))

		pisos = append(pisos, models.Piso{
			ID:     p.ID,
			Numero: p.Numero,
			// Obra: models.Obra{
			// 	ID:     obra.ID,
			// 	Nombre: obra.Nombre,
			// },
		})

	}

	return pisos, nil
}

func (s *serv) GetPisosByObra(ctx context.Context, obraID int64) ([]models.Piso, error) {
	op, err := s.repo.GetObraPisos(ctx, obraID)
	if err != nil {
		return nil, err
	}

	pisos := []models.Piso{}

	for _, p := range op {

		pisos = append(pisos, models.Piso{
			ID:     p.ID,
			Numero: p.Numero,
		})

	}

	return pisos, nil
}

func (s *serv) AddObraPiso(ctx context.Context, obraID, pisoID int64) error {

	pisos, err := s.repo.GetObraPisos(ctx, obraID)
	if err != nil {
		return err
	}

	for _, r := range pisos {
		if int64(r.ID) == pisoID {
			return ErrPisoObraAlreadyExists
		}
	}

	return s.repo.SaveObraPiso(ctx, obraID, pisoID)
}

func (s *serv) UpdatePiso(ctx context.Context, pisoID int64, numero int) error {
	return s.repo.UpdatePiso(ctx, pisoID, numero)
}

func (s *serv) DeletePiso(ctx context.Context, PisoID int64) error {

	dd, _ := s.repo.GetDocumentsByPiso(ctx, PisoID)

	for _, d := range dd {
		s.repo.DeleteDocument(ctx, d.ID)
	}

	return s.repo.DeletePiso(ctx, PisoID)
}

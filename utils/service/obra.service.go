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

func (s *serv) GetObra(ctx context.Context, obraID int64) (*models.Obra, error) {
	oo, err := s.repo.GetObrabyID(ctx, obraID)
	if err != nil {
		return nil, err
	}

	obra := &models.Obra{
		ID:     oo.ID,
		Nombre: oo.Nombre,
	}

	return obra, nil
}

func (s *serv) GetPisosObra(ctx context.Context, ObraID int64) ([]models.Piso, error) {
	po, err := s.repo.GetPisosDeObra(ctx, ObraID)
	if err != nil {
		return nil, err
	}

	Pisos := []models.Piso{}

	for _, P := range po {
		Pisos = append(Pisos, models.Piso{
			ID:     P.ID,
			Numero: P.Numero,
		})
	}

	return Pisos, nil
}

func (s *serv) DeleteObra(ctx context.Context, ObraID int64) error {

	dd, _ := s.repo.GetDocumentsByObra(ctx, ObraID)

	for _, d := range dd {
		s.repo.DeleteDocument(ctx, d.ID)
	}

	pp, _ := s.repo.GetPisosDeObra(ctx, ObraID)

	for _, p := range pp {
		s.repo.DeletePiso(ctx, int64(p.ID))
	}

	return s.repo.DeleteObra(ctx, ObraID)
}

func (s *serv) UpdateObra(ctx context.Context, obraID int64, nombre string) error {
	return s.repo.UpdateObra(ctx, obraID, nombre)
}

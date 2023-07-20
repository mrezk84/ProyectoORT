package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrFormAlreadyExists = errors.New("El fomrmulario ya existe")
	ErrInvalidForm       = errors.New("Formulario Inválido")
	ErrFomeAlreadyAdded  = errors.New("El usuario ya cuenta con el formulario asignado")
	ErrFormNotFound      = errors.New("Error al asignar formulario")
)

func (s *serv) RegisterFrom(ctx context.Context, nombre string, informacion string, version string, fecha string) error {

	f, _ := s.repo.GetFormByVersion(ctx, version)
	if f != nil {
		return ErrFormAlreadyExists
	}

	return s.repo.SaveFrom(ctx, nombre, informacion, version, fecha)
}

func (s *serv) GetFormByDate(ctx context.Context, fechaIni, fechaFin string) (*models.Formulario, error) {

	form, err := s.repo.GetFormByDate(ctx, fechaIni, fechaFin)

	if form != nil {
		return nil, err
	}

	formulario := &models.Formulario{
		ID:          form.ID,
		Informacion: form.Informacion,
		Version:     form.Version,
		Nombre:      form.Nombre,
		Controles:   []models.Control{},
	}
	return formulario, nil
}

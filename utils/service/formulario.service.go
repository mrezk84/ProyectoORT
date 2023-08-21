package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
	"time"
)

var (
	ErrFormAlreadyExists    = errors.New("el fomrmulario ya existe")
	ErrInvalidForm          = errors.New("formulario Inválido")
	ErrFomUserAlreadyAdded  = errors.New("el usuario ya cuenta con el formulario asignado")
	ErrFormNotFound         = errors.New("error al asignar formulario")
	ErrFomEtapaAlreadyAdded = errors.New("la etapa ya se encuentra realizada")
)

func (s *serv) RegisterFrom(ctx context.Context, informacion string, nombre string, version string, fecha string, etapa_id, usuairo_id int) error {

	f, _ := s.repo.GetForms(ctx)
	if f != nil {
		return ErrFormAlreadyExists
	}

	u, _ := s.repo.GetUserById(ctx, usuairo_id)
	if u != nil {
		return ErrFomUserAlreadyAdded
	}

	e, _ := s.repo.GetEtapaById(ctx, int64(etapa_id))
	if u != nil {
		return ErrFomEtapaAlreadyAdded
	}
	return s.repo.SaveFrom(ctx, informacion, nombre, version, fecha, int64(e.ID), u.ID)
}

func (s *serv) AddForm(ctx context.Context, version string, formulario models.Formulario) error {

	form, err := s.repo.GetFormByVersion(ctx, version)

	if form != nil {
		return err
	}

	etapas, err := s.repo.GetFromEtapas(ctx)
	if etapas != nil {
		return err
	}

	usuarios, err := s.repo.GetFromUsers(ctx)
	if usuarios != nil {
		return err
	}
	return s.repo.SaveFrom(ctx, form.Nombre, form.Informacion, form.Version, form.Fecha, form.IDEtapa, form.IDUsuario)
}

func (s *serv) GetForms(ctx context.Context) ([]models.Formulario, error) {

	ff, err := s.repo.GetForms(ctx)
	if err != nil {
		return nil, err
	}

	formularios := []models.Formulario{}
	for _, f := range ff {
		// Crear formulario
		formularios = append(formularios, models.Formulario{
			ID:          f.ID,
			Nombre:      f.Nombre,
			Informacion: f.Informacion,
			Version:     f.Version,
			Fecha:       time.Now(),
			Etapa:       []models.Etapa{{ID: f.IDEtapa}},
			Usuario:     []models.Usuario{{ID: f.IDUsuario}},
		})
	}
	return formularios, nil
}

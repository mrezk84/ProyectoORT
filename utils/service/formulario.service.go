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

func (s *serv) RegisterFrom(ctx context.Context, informacion string, nombre string, version string, fecha string, etapa_id int, usuario_id int) error {

	f, _ := s.repo.GetForms(ctx)
	if f != nil {
		return ErrFormAlreadyExists
	}

	u, _ := s.repo.GetFromUsers(ctx)
	if u != nil {
		return ErrFomUserAlreadyAdded
	}

	e, _ := s.repo.GetFromEtapas(ctx)
	if e != nil {
		return ErrFomEtapaAlreadyAdded
	}
	return s.repo.SaveFrom(ctx, informacion, nombre, version, fecha, e.ID, u.ID)
}

func (s *serv) AddForm(ctx context.Context, id int, formulario models.Formulario) error {

	form, err := s.repo.GetFormsById(ctx, int64(id))
	if form != nil {
		return err
	}

	etapas, err := s.repo.GetEtapaById(ctx, form.IDEtapa)
	if etapas != nil {
		return err
	}

	usuarios, err := s.repo.GetUserById(ctx, int(form.IDUsuario))
	if usuarios != nil {
		return err
	}

	return s.repo.SaveFrom(ctx, form.Nombre, form.Informacion, form.Version, form.Fecha, etapas.ID, int(usuarios.ID))
}

func (s *serv) GetForms(ctx context.Context) ([]models.Formulario, error) {

	ff, err := s.repo.GetForms(ctx)
	if err != nil {
		return nil, err
	}

	formularios := []models.Formulario{}
	for _, f := range ff {

		formularios = append(formularios, models.Formulario{
			ID:          f.ID,
			Nombre:      f.Nombre,
			Informacion: f.Informacion,
			Version:     f.Version,
			Fecha:       time.Now(),
			EtapaID:     int(f.IDEtapa),
			UsuarioID:   int(f.IDUsuario),
		})
	}

	return formularios, nil
}

package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
	"time"
)

var (
	ErrFormAlreadyExists          = errors.New("el fomrmulario ya existe")
	ErrInvalidForm                = errors.New("formulario Inválido")
	ErrFomUserAlreadyAdded        = errors.New("el usuario ya cuenta con el formulario asignado")
	ErrFormNotFound               = errors.New("error al asignar formulario")
	ErrFomEtapaAlreadyAdded       = errors.New("la etapa ya se encuentra realizada")
	ErrInvalidPermissions         = errors.New("el usuario no tiene  permisos para agregar el formulario")
	validRolesToAddForm     []int = []int{1, 2, 3, 4}
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

func (s *serv) AddForm(ctx context.Context, email string, formulario models.Formulario) error {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	}

	et, err := s.repo.GetFromEtapas(ctx)
	if err != nil {
		return err
	}
	userCanAdd := false
	for _, r := range roles {
		for _, vr := range validRolesToAddForm {

			if vr == r.RoleID {
				userCanAdd = true
			}
		}
	}

	if !userCanAdd {
		return ErrInvalidPermissions
	}

	return s.repo.SaveFrom(ctx, formulario.Nombre, formulario.Informacion, formulario.Version, formulario.Fecha.Format("dd/mm/aaaa"), et.ID, u.ID)
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

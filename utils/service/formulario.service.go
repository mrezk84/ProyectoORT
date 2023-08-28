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

	f, _ := s.repo.GetForm(ctx)
	if f != nil {
		return ErrFormAlreadyExists
	}

	return s.repo.SaveFrom(ctx, nombre, informacion, version, fecha)
}

func (s *serv) GetFormByDate(ctx context.Context, fecha string) (*models.Formulario, error) {

	form, err := s.repo.GetFormByDate(ctx, fecha)

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

func (s *serv) GetForms(ctx context.Context) ([]models.Formulario, error) {
	ff, err := s.repo.GetForm(ctx)
	if err != nil {
		return nil, err
	}
	cc, err := s.repo.GetControls(ctx)
	if err != nil {
		return nil, err
	}

	formularios := []models.Formulario{}
	controles := []models.Control{}

	for _, c := range cc {
		controles = append(controles, models.Control{
			ID:          c.ID,
			Descripcion: c.Descripcion,
			Tipo:        c.Tipo,
		})
	}

	for _, f := range ff {
		formularios = append(formularios, models.Formulario{
			ID:          f.ID,
			Informacion: f.Informacion,
			Version:     f.Version,
			Nombre:      f.Nombre,
			Controles:   controles,
		})

	}

	return formularios, nil
}

func (s *serv) AddUserForm(ctx context.Context, formID, usuarioID int64) error {

	usuariosf, err := s.repo.GetUsuarioForm(ctx, usuarioID)
	if err != nil {
		return err
	}

	for _, r := range usuariosf {
		if r.FormularioID == formID {
			return ErrFomeAlreadyAdded
		}
	}

	return s.repo.SaveUserForm(ctx, formID, usuarioID)
}

func (s *serv) GetUserOfForm(ctx context.Context, formID int64) (*models.Usuario, error) {

	userF, err := s.repo.GetFormUser(ctx, formID)

	if userF == nil {
		return nil, err
	}

	userForm := &models.FormularioUser{
		FormularioID: userF.FormularioID,
		UserID:       userF.UsuarioID,
	}

	user, _ := s.repo.GetUserById(ctx, userForm.UserID)

	usuario := &models.Usuario{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	return usuario, nil
}

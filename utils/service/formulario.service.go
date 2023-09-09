package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrFormAlreadyExists          = errors.New("el fomrmulario ya existe")
	ErrInvalidForm                = errors.New("formulario Inválido")
	ErrFomUserAlreadyAdded        = errors.New("el usuario ya cuenta con el formulario asignado")
	ErrFormNotFound               = errors.New("error al asignar formulario")
	ErrFomEtapaAlreadyAdded       = errors.New("el control para el fomrulario ya se encuentra realizado")
	ErrInvalidPermissions         = errors.New("el usuario no tiene  permisos para agregar el formulario")
	validRolesToAddForm     []int = []int{1, 2, 3, 4}
)

func (s *serv) RegisterFrom(ctx context.Context, nombre string, informacion string, version int, control_id int, usuario_id int) error {

	f, _ := s.repo.GetFormByName(ctx, nombre)
	if f != nil {
		return ErrFormAlreadyExists
	}
	u, _ := s.repo.GetFromUsers(ctx)
	if u != nil {
		return ErrFomUserAlreadyAdded
	}

	c, _ := s.repo.GetFormControles(ctx)
	if c != nil {
		return ErrFomEtapaAlreadyAdded
	}

	return s.repo.SaveFrom(ctx, nombre, informacion, version, control_id, usuario_id)
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

	co, err := s.repo.GetFormControles(ctx)
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

	return s.repo.SaveFrom(ctx, formulario.Nombre, formulario.Informacion, formulario.Version,
		co.ID, u.ID)
}

func (s *serv) GetForms(ctx context.Context) ([]models.Formulario, error) {

	ff, err := s.repo.GetForms(ctx)
	if err != nil {
		return nil, err
	}

	foto, err := s.repo.GetPhotos(ctx)
	if err != nil {
		return nil, err
	}

	cc, err := s.repo.GetControls(ctx)
	if err != nil {
		return nil, err
	}

	uu, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	formularios := []models.Formulario{}
	fotos := []models.Foto{}
	controles := []models.Control{}
	usuarios := []models.Usuario{}

	for _, c := range cc {
		controles = append(controles, models.Control{
			ID:          c.ID,
			Descripcion: c.Descripcion,
			Tipo:        c.Tipo,
		})
	}

	for _, fo := range foto {

		fotos = append(fotos, models.Foto{
			ID:     fo.ID,
			Nombre: fo.Nombre,
			Notas:  fo.Notas,
		})

	}

	for _, u := range uu {

		usuarios = append(usuarios, models.Usuario{

			ID:    u.ID,
			Email: u.Email,
			Name:  u.Name,
		})

	}

	for _, f := range ff {

		formularios = append(formularios, models.Formulario{
			ID:          f.ID,
			Informacion: f.Informacion,
			Version:     f.Version,
			Nombre:      f.Nombre,
			Controles:   controles,
			Usuarios:    usuarios,
			Foto:        fotos,
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

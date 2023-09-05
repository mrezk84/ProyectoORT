package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrContAlreadyExists = errors.New("El control ya existe")
	ErrInvalidCont       = errors.New("Control Inválido")
	ErrContAlreadyAdded  = errors.New("Ya se cuenta con el control asignado")
	ErrContNotFound      = errors.New("Error al asignar el control")
)

func (s *serv) RegisterControl(ctx context.Context, descripcion, tipo string) error {

	c, _ := s.repo.GetConByDes(ctx, descripcion)
	if c != nil {
		return ErrFormAlreadyExists
	}

	return s.repo.SaveControl(ctx, descripcion, tipo)
}
func (s *serv) GetControls(ctx context.Context) ([]models.Control, error) {
	cc, err := s.repo.GetControls(ctx)
	if err != nil {
		return nil, err
	}

	controles := []models.Control{}

	for _, c := range cc {
		controles = append(controles, models.Control{
			ID:          c.ID,
			Descripcion: c.Descripcion,
			Tipo:        c.Tipo,
		})

	}

	return controles, nil
}

func (s *serv) GetControlsByForm(ctx context.Context, formID int64) ([]models.Control, error) {
	cc, err := s.repo.GetControlsByForm(ctx, formID)
	if err != nil {
		return nil, err
	}

	controlesF := []models.FormularioControl{}
	controles := []models.Control{}

	for _, cf := range cc {
		controlesF = append(controlesF, models.FormularioControl{
			ControlID:    cf.ControlID,
			FormularioID: cf.FormularioID,
		})
	}

	for _, c := range controlesF {

		control, _ := s.repo.GetConByid(ctx, int(c.ControlID))

		controles = append(controles, models.Control{
			ID:          control.ID,
			Descripcion: control.Descripcion,
			Tipo:        control.Tipo,
		})

	}

	return controles, nil
}

func (s *serv) AddControlForm(ctx context.Context, controlID, formularioID int64) error {

	_, err := s.repo.GetControlForm(ctx, controlID)
	if err != nil {
		return err
	}

	return s.repo.SaveControlForm(ctx, controlID, formularioID)
}

func (s *serv) GetFormdeControl(ctx context.Context, controlID int64) (*models.Formulario, error) {

	cont, err := s.repo.GetControlForm(ctx, controlID)

	if cont != nil {
		return nil, err
	}

	for _, r := range cont {
		id := r.FormularioID

		form, _ := s.repo.GetFormByID(ctx, id)

		formulario := &models.Formulario{
			ID:          form.ID,
			Informacion: form.Informacion,
			Version:     form.Version,
			Nombre:      form.Nombre,
		}
		return formulario, nil
	}

	return nil, err
}

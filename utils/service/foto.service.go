package service

import (
	"context"
	"errors"
	"proyectoort/utils/models"
)

var (
	ErrFotoAlreadyExists    = errors.New("la foto  ya existe")
	ErrInvalidFoto          = errors.New("foto Inválida")
	ErrFomPhotoAlreadyAdded = errors.New("la foto ya se cuenta en el formulario")
	ErrFotoNotFound         = errors.New("error al asignar foto")
	ErrNoPhotosInForm       = errors.New("no hay fotos para agregar a la solicitud")
)

func (s *serv) RegisterPhoto(ctx context.Context, nombre, notas string, formulario_id int) error {

	fo, _ := s.repo.GetPhotoByForm(ctx, formulario_id)
	if fo != nil {
		return ErrFotoAlreadyExists
	}

	return s.repo.SavePhoto(ctx, nombre, notas, fo.FormularioID)

}

func (s *serv) GetPhotos(ctx context.Context) ([]models.Foto, error) {
	fo, err := s.repo.GetPhotos(ctx)
	if err != nil {
		return nil, err
	}

	fotos := []models.Foto{}

	for _, f := range fo {

		fotos = append(fotos, models.Foto{
			ID:           f.ID,
			Nombre:       f.Nombre,
			Notas:        f.Notas,
			FormularioID: f.FormularioID,
		})

	}

	return fotos, nil
}

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

	fo, _ := s.repo.GetPhotos(ctx)
	if fo != nil {
		return ErrFotoAlreadyExists
	}

	return s.repo.SavePhoto(ctx, nombre, notas, formulario_id)

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
func (s *serv) GetPhoto(ctx context.Context, id int) (*models.Foto, error) {
	p, err := s.repo.GetPhotoById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	foto := &models.Foto{
		ID:           p.ID,
		Nombre:       p.Nombre,
		Notas:        p.Notas,
		FormularioID: p.FormularioID,
	}

	return foto, nil
}
func (s *serv) GetPhotoFilePath(ctx context.Context, fotoID int) (string, error) {
	// Llama a la función correspondiente en el repositorio para obtener la ruta del archivo
	filePath, err := s.repo.GetPhotoFilePath(ctx, fotoID)
	if err != nil {
		// Maneja el error si la foto no existe o si ocurre algún problema
		return "", ErrFotoNotFound
	}
	return filePath, nil
}

package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertFoto = `
		INSERT INTO foto (nombre, notas, formulario_id)
		VALUES (?, ?, ?);`

	qryGetPhotoByForm = `
		SELECT 
			id, 
			nombre, 
			notas 
		FROM fotos 
		WHERE formulario_id = ?"`

	qryGetPhotos = `
	 SELECT 
			id, 
			nombre, 
			notas 
		FROM fotos 
	
	`
	qryGetPhotoById = `
		SELECT 
			id, 
			nombre, 
			notas 
		FROM fotos 
		WHERE id = ?"`
)

func (r *repo) SavePhoto(ctx context.Context, nombre, notas string, formulario_id int) error {
	_, err := r.db.ExecContext(ctx, qryInsertFoto, nombre, notas, formulario_id)
	return err
}

func (r *repo) GetPhotoByForm(ctx context.Context, formulario_id int) (*entity.Foto, error) {
	fo := &entity.Foto{}
	err := r.db.GetContext(ctx, fo, qryGetPhotoByForm, formulario_id)
	if err != nil {
		return nil, err
	}

	return fo, nil
}
func (r *repo) GetPhotos(ctx context.Context) ([]entity.Foto, error) {
	fo := []entity.Foto{}
	err := r.db.SelectContext(ctx, &fo, qryGetPhotos)
	if err != nil {
		return nil, err
	}
	return fo, nil
}
func (r *repo) GetPhotoById(ctx context.Context, id int) (*entity.Foto, error) {
	f := &entity.Foto{}

	err := r.db.GetContext(ctx, f, qryGetPhotoById, id)
	if err != nil {
		return nil, err
	}

	return f, nil
}

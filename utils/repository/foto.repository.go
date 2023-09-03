package repository

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
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
func (r *repo) GetPhotoFilePath(ctx context.Context, id int) (string, error) {

	query := "SELECT nombre FROM fotos WHERE id = ?"
	var nombreFoto string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&nombreFoto)
	if err != nil {
		if err == sql.ErrNoRows {
			// Si no se encuentra la foto, devuelve un error
			return "", fmt.Errorf("Foto no encontrada")
		}
		// Si ocurre un error diferente, maneja el error de acuerdo a tus necesidades
		return "", fmt.Errorf("Error a obtener la foto")
	}
	filePath := filepath.Join("fotos", nombreFoto)
	return filePath, nil
}

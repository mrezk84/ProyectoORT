package repository

import (
	"context"
	"proyectoort/utils/models"
)

const (
	qryInsertDocument = `
		INSERT INTO document(formulario_id, obra_id, piso_id)
		VALUES (?,?,?);`
)

func (r *repo) InsertDocument(ctx context.Context, formularioID int64, obraID int64, pisoID int64) (models.Document, error) {
	result, err := r.db.ExecContext(ctx, qryInsertDocument, formularioID, obraID, pisoID)
	if err != nil {
		return models.Document{}, err
	}
	id, err := result.LastInsertId()
	return models.Document{
		ID:         id,
		Formulario: models.Formulario{ID: int(formularioID)},
		Obra:       models.Obra{ID: int(obraID)},
		Piso:       models.Piso{ID: int(pisoID)},
	}, err
}

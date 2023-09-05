package repository

import (
	"context"
	"fmt"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"
)

const (
	qryInsertDocument = `
		INSERT INTO document (formulario_id,obra_id,piso_id)
		VALUES (%v,%v,%v);`

	getDocumentsByObra = `
		select * from documents where obra_id = ?`
)

func (r *repo) InsertDocument(ctx context.Context, formularioID int64, obraID int64, pisoID int64) (models.Document, error) {
	result, err := r.db.ExecContext(ctx, fmt.Sprintf(qryInsertDocument, formularioID, obraID, pisoID))
	if err != nil {
		fmt.Println(err)
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

func (r *repo) GetDocumentsByObra(ctx context.Context, obraID int64) ([]models.Document, error) {
	e := []entity.Document{}
	err := r.db.SelectContext(ctx, e, getDocumentsByObra, obraID)
	if err != nil {
		return nil, err
	}
	var documents []models.Document
	for _, d := range e {
		documents = append(documents, models.Document{
			ID: d.ID,
			Obra: models.Obra{
				ID: int(d.ObraID),
			},
			Formulario: models.Formulario{
				ID: int(d.FormularioID),
			},
		})
	}
	return documents, nil

}

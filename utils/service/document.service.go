package service

import (
	"context"
	"proyectoort/utils/models"
)

func (s *serv) InsertDocument(ctx context.Context, formularioID int64, obraID int64, pisoID int64) (models.Document, error) {
	return s.repo.InsertDocument(ctx, formularioID, obraID, pisoID)
}

func (s *serv) InsertChecks(ctx context.Context, controles []models.Control, document models.Document, formularioID int64) error {
	return s.repo.InsertChecks(ctx, formularioID, document.ID, controles)
}

func (s *serv) GetObraDocuments(ctx context.Context, obraID int64) ([]models.Document, error) {
	return s.repo.GetDocumentsByObra(ctx, obraID)
}

// func (s *serv) GetDocumentPDF() ([]byte, error) {
// 	return s.repo.ExportDocument()
// }

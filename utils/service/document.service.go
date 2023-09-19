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

func (s *serv) GetResponsableDocuments(ctx context.Context, obraID, usuarioID int64) ([]models.Document, error) {
	do, _ := s.repo.GetDocumentsByObra(ctx, obraID)

	documentos := []models.Document{}

	for _, d := range do {
		user, _ := s.repo.GetUserForm(ctx, int64(d.Formulario.ID))
		if user.ID == usuarioID {
			documentos = append(documentos, d)
		}
	}

	return documentos, nil
}

func (s *serv) GetDocumentPDF(ctx context.Context, documentID int64) ([]byte, error) {
	return s.repo.ExportDocument(ctx, documentID)
}

func (s *serv) GetDocumentsPDFByObra(ctx context.Context, obraID int64) ([]byte, error) {
	return s.repo.ExportDocumentsByObra(ctx, obraID)
}

func (s *serv) DeleteDocument(ctx context.Context, DocID int64) error {
	return s.repo.DeleteDocument(ctx, DocID)
}

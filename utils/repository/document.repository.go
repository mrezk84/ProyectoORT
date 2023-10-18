package repository

import (
	"bytes"
	"context"
	"fmt"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"

	"github.com/jung-kurt/gofpdf"
)

const (
	qryInsertDocument = `
		INSERT INTO document (formulario_id,obra_id,piso_id)
		VALUES (?,?,?);`

	qryGetFormularioByDocumentID = `
	select f.id,f.nombre,f.informacion from FORMULARIO f inner join document d on f.id = d.formulario_id
where d.id = %v
	`
	getDocumentsByObra = `
		select * from document where obra_id = ?`

	getDocumentsByForm = `
		select * from document where formulario_id = ?`

	getDocumentsByPiso = `
		select * from document where piso_id = ?`

	qryDeleteChecksFromDocument = `
		DELETE FROM CHECKS where document_id = ?`

	qryDeleteDocumento = `
		DELETE FROM document where id = ?`

	qryGetWipOrTodoDocuments = `select d.id,d.formulario_id,d.obra_id,d.piso_id FROM document d
LEFT JOIN CHECKS c ON d.id = c.document_id
WHERE c.estado != 'CONFORME' OR c.estado IS NULL
GROUP BY d.id`

	qryGetWipOrTodoDocumentsByFormID = `select d.id,d.formulario_id,d.obra_id,d.piso_id FROM document d
LEFT JOIN CHECKS c ON d.id = c.document_id
WHERE c.estado != 'CONFORME' OR c.estado IS NULL
and d.formulario_id = %v
GROUP BY d.id`
)

func (r *repo) InsertDocument(ctx context.Context, formularioID int64, obraID int64, pisoID int64) (models.Document, error) {
	result, err := r.db.ExecContext(ctx, qryInsertDocument, formularioID, obraID, pisoID)
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
	err := r.db.SelectContext(ctx, &e, getDocumentsByObra, obraID)
	if err != nil {
		return nil, err
	}
	var documents []models.Document
	for _, d := range e {
		formulario, err := r.GetFormByID(ctx, d.FormularioID)
		if err != nil {
			return nil, err
		}
		piso, err := r.GetPisobyID(ctx, d.PisoID)
		if err != nil {
			return nil, err
		}
		documentChecks, err := r.GetDocumentChecks(ctx, d.ID)
		if err != nil {
			return nil, err
		}
		status := getDocumentStatusFromChecks(documentChecks)
		documents = append(documents, models.Document{
			ID: d.ID,
			Obra: models.Obra{
				ID: int(d.ObraID),
			},
			Formulario: models.Formulario{
				ID:          int(d.FormularioID),
				Informacion: formulario.Informacion,
				Version:     formulario.Version,
				Nombre:      formulario.Nombre,
			},
			Piso: models.Piso{
				ID:     int(d.PisoID),
				Numero: piso.Numero,
			},
			Checks: documentChecks,
			Status: status,
		})
	}
	return documents, nil

}

func (r *repo) GetDocumentsByForm(ctx context.Context, formID int64) ([]models.Document, error) {
	e := []entity.Document{}
	err := r.db.SelectContext(ctx, &e, getDocumentsByForm, formID)
	if err != nil {
		return nil, err
	}
	var documents []models.Document
	for _, d := range e {
		formulario, err := r.GetFormByID(ctx, d.FormularioID)
		if err != nil {
			return nil, err
		}
		piso, err := r.GetPisobyID(ctx, d.PisoID)
		if err != nil {
			return nil, err
		}
		documentChecks, err := r.GetDocumentChecks(ctx, d.ID)
		if err != nil {
			return nil, err
		}
		status := getDocumentStatusFromChecks(documentChecks)
		documents = append(documents, models.Document{
			ID: d.ID,
			Obra: models.Obra{
				ID: int(d.ObraID),
			},
			Formulario: models.Formulario{
				ID:          int(d.FormularioID),
				Informacion: formulario.Informacion,
				Version:     formulario.Version,
				Nombre:      formulario.Nombre,
			},
			Piso: models.Piso{
				ID:     int(d.PisoID),
				Numero: piso.Numero,
			},
			Checks: documentChecks,
			Status: status,
		})
	}
	return documents, nil

}

func (r *repo) GetDocumentsByPiso(ctx context.Context, pisoID int64) ([]models.Document, error) {
	e := []entity.Document{}
	err := r.db.SelectContext(ctx, &e, getDocumentsByPiso, pisoID)
	if err != nil {
		return nil, err
	}
	var documents []models.Document
	for _, d := range e {
		formulario, err := r.GetFormByID(ctx, d.FormularioID)
		if err != nil {
			return nil, err
		}
		piso, err := r.GetPisobyID(ctx, d.PisoID)
		if err != nil {
			return nil, err
		}
		documentChecks, err := r.GetDocumentChecks(ctx, d.ID)
		if err != nil {
			return nil, err
		}
		status := getDocumentStatusFromChecks(documentChecks)
		documents = append(documents, models.Document{
			ID: d.ID,
			Obra: models.Obra{
				ID: int(d.ObraID),
			},
			Formulario: models.Formulario{
				ID:          int(d.FormularioID),
				Informacion: formulario.Informacion,
				Version:     formulario.Version,
				Nombre:      formulario.Nombre,
			},
			Piso: models.Piso{
				ID:     int(d.PisoID),
				Numero: piso.Numero,
			},
			Checks: documentChecks,
			Status: status,
		})
	}
	return documents, nil

}

func (r *repo) DeleteDocument(ctx context.Context, DocID int64) error {
	_, err := r.db.ExecContext(ctx, qryDeleteChecksFromDocument, DocID)
	if err != nil {
		return err
	}
	_, err2 := r.db.ExecContext(ctx, qryDeleteDocumento, DocID)
	return err2
}

func (r *repo) ExportDocument(ctx context.Context, documentID int64) ([]byte, error) {
	checks, err := r.GetDocumentChecks(ctx, documentID)
	if err != nil {
		return nil, err
	}
	// Crear un nuevo documento PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	// Configurar la fuente y el tamaño del texto
	pdf.SetFont("Arial", "B", 16)

	form, err := r.GetFormularioByDocumentID(documentID)
	if err != nil {
		return nil, err
	}
	pdf.AddPage()
	pdf.Cell(100, 16, fmt.Sprintf("Formulario: %v", form.Nombre))
	pdf.Ln(15)
	for _, check := range checks {
		// Agregar una página al documento
		pdf.Cell(100, 16, "Estado: "+check.Estado)
		pdf.Ln(15)
		pdf.Cell(100, 16, "Observaciones: "+check.Observaciones)
		pdf.Ln(15)
		pdf.Cell(100, 16, "Fecha control: "+check.FechaControl.String())
		pdf.Ln(15)
		pdf.Cell(100, 16, "Responsable: "+check.Responsable.Name)
		pdf.Ln(30)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buf.Bytes(), err
}

func (r *repo) ExportDocumentsByObra(ctx context.Context, obraID int64) ([]byte, error) {
	documents, err := r.GetDocumentsByObra(ctx, obraID)
	if err != nil {
		return nil, err
	}

	// Crear un nuevo documento PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	// Configurar la fuente y el tamaño del texto
	pdf.SetFont("Arial", "B", 16)
	if err != nil {
		return nil, err
	}

	for _, d := range documents {
		pdf.AddPage()
		pdf.Cell(100, 16, fmt.Sprintf("Formulario: %v", d.Formulario.Nombre))
		pdf.Ln(15)
		for _, check := range d.Checks {
			// Agregar una página al documento
			pdf.Cell(100, 16, "Estado: "+check.Estado)
			pdf.Ln(15)
			pdf.Cell(100, 16, "Observaciones: "+check.Observaciones)
			pdf.Ln(15)
			if check.FechaControl != nil {
				pdf.Cell(100, 16, "Fecha control: "+check.FechaControl.String())
				pdf.Ln(15)
			}
			pdf.Cell(100, 16, "Responsable: "+check.Responsable.Name)
			pdf.Ln(30)
		}
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return buf.Bytes(), err
}

func getDocumentStatusFromChecks(checks []models.Check) string {
	todosVacios := true
	todosConforme := true

	for _, check := range checks {
		if check.Estado != "" {
			todosVacios = false
		}
		if check.Estado != "CONFORME" {
			todosConforme = false
		}
	}

	if todosVacios {
		return "TODO"
	} else if todosConforme {
		return "DONE"
	} else {
		return "WIP"
	}
}

func (r *repo) GetFormularioByDocumentID(documentID int64) (*models.Formulario, error) {
	var formulario entity.Formulario
	err := r.db.Get(&formulario, fmt.Sprintf(qryGetFormularioByDocumentID, documentID))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &models.Formulario{
		ID:          formulario.ID,
		Informacion: formulario.Informacion,
		Version:     formulario.Version,
		Nombre:      formulario.Nombre,
	}, nil
}

func (r *repo) GetWipOrTodoDocuments(ctx context.Context) ([]models.Document, error) {
	e := []entity.Document{}
	err := r.db.SelectContext(ctx, &e, qryGetWipOrTodoDocuments)
	if err != nil {
		return nil, err
	}
	var documents []models.Document
	for _, d := range e {
		documents = append(documents, models.Document{
			ID:         d.ID,
			Formulario: models.Formulario{ID: int(d.FormularioID)},
			Obra:       models.Obra{ID: int(d.ObraID)},
			Piso:       models.Piso{ID: int(d.PisoID)},
		})
	}
	return documents, nil
}

func (r *repo) GetWipOrTodoDocumentsByFormID(ctx context.Context, formID int64) ([]models.Document, error) {
	e := []entity.Document{}
	err := r.db.SelectContext(ctx, &e, fmt.Sprintf(qryGetWipOrTodoDocumentsByFormID, formID))
	if err != nil {
		return nil, err
	}
	var documents []models.Document
	for _, d := range e {
		documents = append(documents, models.Document{
			ID:         d.ID,
			Formulario: models.Formulario{ID: int(d.FormularioID)},
			Obra:       models.Obra{ID: int(d.ObraID)},
			Piso:       models.Piso{ID: int(d.PisoID)},
		})
	}
	return documents, nil
}

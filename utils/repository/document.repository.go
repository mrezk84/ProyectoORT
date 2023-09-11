package repository

import (
	"bytes"
	"context"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"
)

const (
	qryInsertDocument = `
		INSERT INTO document (formulario_id,obra_id,piso_id)
		VALUES (%v,%v,%v);`

	qryGetFormularioByDocumentID = `
	select f.id,f.nombre,f.informacion from FORMULARIO f inner join document d on f.id = d.formulario_id
where d.id = %v
	`
	getDocumentsByObra = `
		select * from document where obra_id = ?`
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

func (r *repo) ExportDocument(ctx context.Context, documentID int64) ([]byte, error) {
	checks, err := r.GetDocumentChecks(ctx, documentID)
	fmt.Println("1")
	if err != nil {
		return nil, err
	}
	fmt.Println("2")
	// Crear un nuevo documento PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	// Configurar la fuente y el tamaño del texto
	pdf.SetFont("Arial", "B", 16)

	form, err := r.GetFormularioByDocumentID(documentID)
	if err != nil {
		return nil, err
	}
	fmt.Println("3")
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
	fmt.Println("4")

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	fmt.Printf(fmt.Sprintf("%v", buf.Bytes()))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("5")
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
			pdf.Cell(100, 16, "Fecha control: "+check.FechaControl.String())
			pdf.Ln(15)
			pdf.Cell(100, 16, "Responsable: "+check.Responsable.Name)
			pdf.Ln(30)
		}
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	fmt.Printf(fmt.Sprintf("%v", buf.Bytes()))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("5")
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

package service

import (
	"context"
	"proyectoort/utils/models"
	"proyectoort/utils/repository"
)

// Service is the business logic of the application.
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	RegisterControl(ctx context.Context, descripcion, tipo string) error
	LoginUser(ctx context.Context, email, password string) (*models.Usuario, error)
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	RegisterFrom(ctx context.Context, nombre string, informacion string) error
	GetFormByDate(ctx context.Context, fecha string) (*models.Formulario, error)
	GetForms(ctx context.Context) ([]models.Formulario, error)
	UpdateFormulario(ctx context.Context, formID int64, nombre, informacion string) error
	AddUserForm(ctx context.Context, formID, usuarioID int64) error
	GetUserOfForm(ctx context.Context, formID int64) (*models.Usuario, error)
	GetControls(ctx context.Context) ([]models.Control, error)
	AddControlForm(ctx context.Context, controlID, formularioID int64) error
	GetFormdeControl(ctx context.Context, controlID int64) (*models.Formulario, error)
	GetControlSinForm(ctx context.Context) ([]models.Control, error)
	GetUsers(ctc context.Context) ([]models.Usuario, error)
	GetControlsByForm(ctx context.Context, formID int64) ([]models.Control, error)
	RegisterObra(ctx context.Context, name string) error
	GetObras(ctx context.Context) ([]models.Obra, error)
	GetObra(ctx context.Context, obraID int64) (*models.Obra, error)
	UpdateObra(ctx context.Context, obraID int64, nombre string) error
	RegisterEtapa(ctx context.Context, nombre string) error
	RegisterPiso(ctx context.Context, number int) (models.Piso, error)
	GetPisos(ctx context.Context) ([]models.Piso, error)
	AddObraPiso(ctx context.Context, obraID, pisoID int64) error
	GetPisosObra(ctx context.Context, ObraID int64) ([]models.Piso, error)
	GetPisosByObra(ctx context.Context, obraID int64) ([]models.Piso, error)
	UpdatePiso(ctx context.Context, pisoID int64, numero int) error
	RegisterCheck(ctx context.Context, estado string, fecha string, observaciones string, version int) error
	AddCheckForm(ctx context.Context, checkID, formularioID int64) error
	InsertDocument(ctx context.Context, formularioID int64, obraID int64, pisoID int64) (models.Document, error)
	InsertChecks(ctx context.Context, controles []models.Control, document models.Document, formularioID int64) error
	DeleteObra(ctx context.Context, ObraID int64) error
	GetObraDocuments(ctx context.Context, obraID int64) ([]models.Document, error)
	// GetDocumentsChecks(ctx context.Context, documents []models.Document) ([]models.Check, error)
	RegisterPhoto(ctx context.Context, nombre, notas string, formulario_id int) error
	GetPhotos(ctx context.Context) ([]models.Foto, error)
	GetPhoto(ctx context.Context, id int) (*models.Foto, error)
	GetPhotoFilePath(ctx context.Context, fotoID int) (string, error)
	GetDocumentChecks(ctx context.Context, documentID int64) ([]models.Check, error)
	UpdateCheck(ctx context.Context, checkID int64, estado, observaciones string) error
	GetDocumentPDF(ctx context.Context, documentID int64) ([]byte, error)
	GetDocumentsPDFByObra(ctx context.Context, obraID int64) ([]byte, error)
}

type serv struct {
	repo repository.Repository
}

// GetFrom implements Service.

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}

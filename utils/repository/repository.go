package repository

import (
	"context"
	"proyectoort/utils/entity"
	"proyectoort/utils/models"

	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps the basic CRUD operations.
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, username, password string) error
	SaveFrom(ctx context.Context, informacion string, nombre string) error
	GetUsuarioForm(ctx context.Context, usuarioID int64) ([]entity.UsuarioForm, error)
	GetFormUser(ctx context.Context, formularioID int64) (*entity.UsuarioForm, error)
	UpdateFormulario(ctx context.Context, FormID int64, nombre, informacion string) error
	SaveUserForm(ctx context.Context, formID, usuarioID int64) error
	SaveUserRole(ctx context.Context, userID, roleID int64) error
	SaveControl(ctx context.Context, descripcion, tipo string) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetUserByEmail(ctx context.Context, email string) (*entity.Usuario, error)
	GetUserById(ctx context.Context, id int64) (*entity.Usuario, error)
	GetUserRoles(ctx context.Context, userID int64) ([]entity.UsarioRol, error)
	GetUsers(ctx context.Context) ([]entity.Usuario, error)
	GetForm(ctx context.Context) ([]entity.Formulario, error)
	GetFormByNombre(ctx context.Context, nombre string) (*entity.Formulario, error)
	GetControls(ctx context.Context) ([]entity.Control, error)
	GetControlsByForm(ctx context.Context, formID int64) ([]entity.Control, error)
	InsertDocument(ctx context.Context, formularioID int64, obraID int64, pisoID int64) (models.Document, error)
	InsertChecks(ctx context.Context, formularioID int64, documentID int64, controles []models.Control) error
	GetConByDes(ctx context.Context, des string) (*entity.Control, error)
	GetControlForm(ctx context.Context, controlID int64) ([]entity.ControlForm, error)
	SaveControlForm(ctx context.Context, controlID, formularioID int64) error
	GetControlsSinForm(ctx context.Context) ([]entity.Control, error)
	GetFormByDate(ctx context.Context, fecha string) (*entity.Formulario, error)
	GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error)
	GetFormByID(ctx context.Context, formID int64) (*entity.Formulario, error)
	GetFromControles(ctx context.Context, controles string) (*entity.Formulario, error)
	SaveObra(ctx context.Context, nombre string) error
	GetObras(ctx context.Context) ([]entity.Obra, error)
	GetObrabyName(ctx context.Context, name string) (*entity.Obra, error)
	GetObrabyID(ctx context.Context, obraID int64) (*entity.Obra, error)
	GetobraP(ctx context.Context, pisoID int64) (*entity.Obra, error)
	UpdateObra(ctx context.Context, obraID int64, nombre string) error
	SaveEtapa(ctx context.Context, nombre string) error
	GetEtapabyName(ctx context.Context, nombre string) (*entity.Etapa, error)
	SavePiso(ctx context.Context, number int) (models.Piso, error)
	GetPisos(ctx context.Context) ([]entity.Piso, error)
	GetPisobyNumber(ctx context.Context, number int) (*entity.Piso, error)
	GetPisobyID(ctx context.Context, ID int64) (*entity.Piso, error)
	GetObraPisos(ctx context.Context, obraID int64) ([]entity.PisoObra, error)
	SaveObraPiso(ctx context.Context, obraID, pisoID int64) error
	UpdatePiso(ctx context.Context, pisoID int64, numero int) error
	SaveCheck(ctx context.Context, estado, observaciones string, version int, fecha string) error
	GetCheckByVersion(ctx context.Context, version int) (*entity.Check, error)
	GetCheckForm(ctx context.Context, FormularioID int64) ([]entity.CheckFormulario, error)
	SaveCheckForm(ctx context.Context, checkID, formularioID int64) error
	DeleteObra(ctx context.Context, nombre string) error
	GetDocumentsByObra(ctx context.Context, obraID int64) ([]models.Document, error)
	GetDocumentsByForm(ctx context.Context, formID int64) ([]models.Document, error)
	getDocumentsByPiso(ctx context.Context, pisoID int64) ([]models.Document, error)
	DeleteDocument(ctx context.Context, DocID int64) error
	// GetDocumentsChecks(ctx context.Context, documents []models.Document) ([]models.Check, error)
	SavePhoto(ctx context.Context, nombre, notas string, formulario_id int) error
	GetPhotoByForm(ctx context.Context, formulario_id int) (*entity.Foto, error)
	GetPhotos(ctx context.Context) ([]entity.Foto, error)
	GetPhotoById(ctx context.Context, id int) (*entity.Foto, error)
	GetPhotoFilePath(ctx context.Context, id int) (string, error)
	GetDocumentChecks(ctx context.Context, documentID int64) ([]models.Check, error)
	UpdateCheck(ctx context.Context, checkID int64, estado, observaciones string) error
	ExportDocument(ctx context.Context, documentID int64) ([]byte, error)
	ExportDocumentsByObra(ctx context.Context, obraID int64) ([]byte, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

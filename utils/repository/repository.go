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
	SaveFrom(ctx context.Context, informacion string, nombre string, version string, fecha string) error
	SaveUserRole(ctx context.Context, userID, roleID int64) error
	SaveControl(ctx context.Context, descripcion, tipo string) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error

	GetUserByEmail(ctx context.Context, email string) (*entity.Usuario, error)
	GetUserRoles(ctx context.Context, userID int64) ([]entity.UsarioRol, error)
	GetUsers(ctx context.Context) ([]entity.Usuario, error)
	GetForm(ctx context.Context) ([]entity.Formulario, error)
	GetControls(ctx context.Context) ([]entity.Control, error)
	GetControlsByForm(ctx context.Context, formID int64) ([]entity.Control, error)
	InsertDocument(ctx context.Context, formularioID int64, obraID int64, pisoID int64) (models.Document, error)
	InsertChecks(ctx context.Context, formularioID int64, documentID int64, controles []models.Control) error
	GetFormByDate(ctx context.Context, fecha string) (*entity.Formulario, error)
	GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error)
	GetFromControles(ctx context.Context, controles string) (*entity.Formulario, error)
	SaveObra(ctx context.Context, nombre string) error
	GetObrabyName(ctx context.Context, name string) (*entity.Obra, error)
	SaveEtapa(ctx context.Context, nombre string) error
	GetEtapabyName(ctx context.Context, nombre string) (*entity.Etapa, error)
	SavePiso(ctx context.Context, number int64) error
	GetPisobyNumber(ctx context.Context, number int64) (*entity.Piso, error)
	GetObraPisos(ctx context.Context, obraID int64) ([]entity.PisoObra, error)
	SaveObraPiso(ctx context.Context, obraID, pisoID int64) error
	SaveCheck(ctx context.Context, estado, observaciones string, version int, fecha string) error
	GetCheckByVersion(ctx context.Context, version int) (*entity.Check, error)
	GetCheckForm(ctx context.Context, FormularioID int64) ([]entity.CheckFormulario, error)
	SaveCheckForm(ctx context.Context, checkID, formularioID int64) error
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

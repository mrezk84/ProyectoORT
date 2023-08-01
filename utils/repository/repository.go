package repository

import (
	"context"
	"proyectoort/utils/entity"

	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps the basic CRUD operations.
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, username, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.Usuario, error)
	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetUserRoles(ctx context.Context, userID int64) ([]entity.UsarioRol, error)
	GetUsers(ctx context.Context) ([]entity.Usuario, error)
	GetFrom(ctx context.Context) ([]entity.Formulario, error)
	SaveFrom(ctx context.Context, informacion string, nombre string, version string, fecha string) error
	GetFormByDate(ctx context.Context, fechaIni, fechaFin string) (*entity.Formulario, error)
	GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error)
	SaveObra(ctx context.Context, nombre string) error
	GetObrabyName(ctx context.Context, name string) (*entity.Obra, error)
	SaveEtapa(ctx context.Context, nombre string) error
	GetEtapabyName(ctx context.Context, name string) (*entity.Etapa, error)
	SavePiso(ctx context.Context, number int64) error
	GetPisobyNumber(ctx context.Context, number int64) (*entity.Piso, error)
	GetObraPisos(ctx context.Context, obraID int64) ([]entity.PisoObra, error)
	SaveObraPiso(ctx context.Context, obraID, pisoID int64) error
	SaveCheck(ctx context.Context, estado, fecha, observaciones string, version int) error
	GetCheckByVersion(ctx context.Context, version int) (*entity.Check, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

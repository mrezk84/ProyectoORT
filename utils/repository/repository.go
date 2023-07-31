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
	SaveFrom(ctx context.Context, informacion string, nombre string, version string, fecha string) error
	SaveUserRole(ctx context.Context, userID, roleID int64) error
	SaveControl(ctx context.Context, descripcion, tipo string) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error

	GetUserByEmail(ctx context.Context, email string) (*entity.Usuario, error)
	GetUserRoles(ctx context.Context, userID int64) ([]entity.UsarioRol, error)
	GetUsers(ctx context.Context) ([]entity.Usuario, error)
	GetForm(ctx context.Context) ([]entity.Formulario, error)
	GetControls(ctx context.Context) ([]entity.Control, error)
	GetFormByDate(ctx context.Context, fecha string) (*entity.Formulario, error)
	GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error)
	GetFromControles(ctx context.Context, controles string) (*entity.Formulario, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

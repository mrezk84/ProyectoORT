package repository

import (
	"context"
	"proyectoort/utils/entity"
	"time"

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
	GetFormByVersion(ctx context.Context, version int) (*entity.Formulario, error)
	GetFrom(ctx context.Context) ([]entity.Formulario, error)
	SaveFrom(ctx context.Context, nombre, informacion string, version int, fecha *time.Time) error
	GetFormByDate(ctx context.Context, fechaIni, fechaFin *time.Time) (*entity.Formulario, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

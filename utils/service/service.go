package service

import (
	"context"
	"proyectoort/utils/models"
	"proyectoort/utils/repository"
	"time"
)

// Service is the business logic of the application.
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.Usuario, error)
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	RegisterFrom(ctx context.Context, nombre, informacion string, version int, fecha *time.Time) error
	GetFormByDate(ctx context.Context, fechaIni, fechaFin *time.Time) (*models.Formulario, error)
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

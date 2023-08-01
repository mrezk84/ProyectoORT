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
	LoginUser(ctx context.Context, email, password string) (*models.Usuario, error)
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	RegisterFrom(ctx context.Context, nombre string, informacion string, version string, fecha string) error
	GetFormByDate(ctx context.Context, fechaIni, fechaFin string) (*models.Formulario, error)
	RegisterObra(ctx context.Context, name string) error
	RegisterEtapa(ctx context.Context, name string) error
	RegisterPiso(ctx context.Context, number int64) error
	AddObraPiso(ctx context.Context, obraID, pisoID int64) error
	RegisterCheck(ctx context.Context, estado string, fecha string, observaciones string, version int) error
	AddCheckForm(ctx context.Context, checkID, formularioID int64) error
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

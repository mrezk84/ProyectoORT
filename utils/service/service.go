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
	RegisterFrom(ctx context.Context, nombre string, informacion string, version string, fecha string) error
	GetFormByDate(ctx context.Context, fecha string) (*models.Formulario, error)
	GetForms(ctx context.Context) ([]models.Formulario, error)
	GetControls(ctx context.Context) ([]models.Control, error)
	GetUsers(ctc context.Context) ([]models.Usuario, error)
	RegisterObra(ctx context.Context, name string) error
	RegisterEtapa(ctx context.Context, nombre string) error
	RegisterPiso(ctx context.Context, number int64) error
	AddObraPiso(ctx context.Context, obraID, pisoID int64) error
	RegisterCheck(ctx context.Context, estado string, observaciones string, version int, fecha string) error
	AddCheckForm(ctx context.Context, checkID, formularioID int64) error
	sendEmail(toEmail, code string) error
	generateRandomCode() string
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

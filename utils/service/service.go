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
	AddUserRole(ctx context.Context, userID, roleID int) error
	AddObraPiso(ctx context.Context, obraID, pisoID int64) error
	AddForm(ctx context.Context, email string, formulario models.Formulario) error
	AddCheckForm(ctx context.Context, checkID, formularioID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int) error
	RegisterFrom(ctx context.Context, informacion string, nombre string, version string, fecha string, etapa_id int, usuario_id int) error
	GetForms(ctx context.Context) ([]models.Formulario, error)
	GetControls(ctx context.Context) ([]models.Control, error)
	GetUsers(ctc context.Context) ([]models.Usuario, error)
	RegisterObra(ctx context.Context, name string) error
	RegisterEtapa(ctx context.Context, nombre string) error
	RegisterPiso(ctx context.Context, id, numero int) error
	RegisterCheck(ctx context.Context, estado string, observaciones string, version int, fecha string) error

	GetAllRoles(ctx context.Context) ([]models.Rol, error)
	RegisterRol(ctx context.Context, id int) error
	GetUsersRole(ctx context.Context, userID int) ([]models.UsuarioRol, error)
	GetPisos(ctx context.Context) ([]models.Piso, error)

	RegisterPhoto(ctx context.Context, nombre, notas string, formulario_id int) error
	GetPhotos(ctx context.Context) ([]models.Foto, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}

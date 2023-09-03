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
	SaveFrom(ctx context.Context, informacion string, version int, nombre string, control_id int, usuario_id int) error
	SaveUserRole(ctx context.Context, userID, roleID int) error
	SaveControl(ctx context.Context, descripcion, tipo string) error
	SaveCheck(ctx context.Context, estado, observaciones string, version int, fecha string) error
	SavePhoto(ctx context.Context, nombre, notas string, formulario_id int) error
	SaveCheckForm(ctx context.Context, checkID, formularioID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int) error
	GetChecks(ctx context.Context) ([]entity.Check, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.Usuario, error)
	GetUsers(ctx context.Context) ([]entity.Usuario, error)
	GetForms(ctx context.Context) ([]entity.Formulario, error)
	GetFormsById(ctx context.Context, id int64) (*entity.Formulario, error)
	GetControls(ctx context.Context) ([]entity.Control, error)
	GetControlById(ctx context.Context, id int) (*entity.Control, error)
	GetUserById(ctx context.Context, id int) (*entity.Usuario, error)
	GetFormById(ctx context.Context, id int) (*entity.Formulario, error)
	GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error)
	GetFormControles(ctx context.Context) (*entity.Formulario, error)
	GetFormPhotos(ctx context.Context) (*entity.Formulario, error)
	GetFromUsers(ctx context.Context) (*entity.Formulario, error)
	SaveObra(ctx context.Context, nombre string) error
	GetObrabyName(ctx context.Context, name string) (*entity.Obra, error)
	SaveEtapa(ctx context.Context, nombre string) error
	GetEtapaByName(ctx context.Context, nombre string) (*entity.Etapa, error)
	GetEtapaById(ctx context.Context, id int64) (*entity.Etapa, error)
	SavePiso(ctx context.Context, id, numero int) error
	GetPisobyNumber(ctx context.Context, numero int) (*entity.Piso, error)
	GetObraPisos(ctx context.Context, obraID int64) ([]entity.PisoObra, error)
	SaveObraPiso(ctx context.Context, obraID, pisoID int64) error
	GetCheckByVersion(ctx context.Context, version int) (*entity.Check, error)
	GetCheckForm(ctx context.Context, FormularioID int64) ([]entity.CheckFormulario, error)
	GetFormByName(ctx context.Context, nombre string) (*entity.Formulario, error)

	SaveRole(ctx context.Context, id int) error
	GetRolByName(ctx context.Context, nombre string) (*entity.Rol, error)
	GetAllRoles(ctx context.Context) ([]entity.Rol, error)
	GetRolById(ctx context.Context, id int) (*entity.Rol, error)
	GetUserRoles(ctx context.Context, userID int) ([]entity.UsuarioRol, error)
	GetPhotoByForm(ctx context.Context, formulario_id int) (*entity.Foto, error)
	GetPisos(ctx context.Context) ([]entity.Piso, error)
	GetPhotos(ctx context.Context) ([]entity.Foto, error)
	GetPhotoById(ctx context.Context, id int) (*entity.Foto, error)
	GetPhotoFilePath(ctx context.Context, id int) (string, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

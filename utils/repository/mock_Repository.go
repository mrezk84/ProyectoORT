// Code generated by mockery v2.33.0. DO NOT EDIT.

package repository

import (
	context "context"
	entity "proyectoort/utils/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// GetCheckByVersion provides a mock function with given fields: ctx, version
func (_m *MockRepository) GetCheckByVersion(ctx context.Context, version int) (*entity.Check, error) {
	ret := _m.Called(ctx, version)

	var r0 *entity.Check
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*entity.Check, error)); ok {
		return rf(ctx, version)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Check); ok {
		r0 = rf(ctx, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Check)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCheckForm provides a mock function with given fields: ctx, FormularioID
func (_m *MockRepository) GetCheckForm(ctx context.Context, FormularioID int64) ([]entity.CheckFormulario, error) {
	ret := _m.Called(ctx, FormularioID)

	var r0 []entity.CheckFormulario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]entity.CheckFormulario, error)); ok {
		return rf(ctx, FormularioID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entity.CheckFormulario); ok {
		r0 = rf(ctx, FormularioID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.CheckFormulario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, FormularioID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChecks provides a mock function with given fields: ctx
func (_m *MockRepository) GetChecks(ctx context.Context) ([]entity.Check, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Check
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Check, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Check); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Check)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetControlById provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetControlById(ctx context.Context, id int) (*entity.Control, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Control
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*entity.Control, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Control); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Control)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetControls provides a mock function with given fields: ctx
func (_m *MockRepository) GetControls(ctx context.Context) ([]entity.Control, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Control
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Control, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Control); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Control)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEtapaById provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetEtapaById(ctx context.Context, id int64) (*entity.Etapa, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Etapa
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.Etapa, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Etapa); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Etapa)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEtapaByName provides a mock function with given fields: ctx, nombre
func (_m *MockRepository) GetEtapaByName(ctx context.Context, nombre string) (*entity.Etapa, error) {
	ret := _m.Called(ctx, nombre)

	var r0 *entity.Etapa
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Etapa, error)); ok {
		return rf(ctx, nombre)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Etapa); ok {
		r0 = rf(ctx, nombre)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Etapa)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nombre)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFormByDate provides a mock function with given fields: ctx, fecha
func (_m *MockRepository) GetFormByDate(ctx context.Context, fecha string) (*entity.Formulario, error) {
	ret := _m.Called(ctx, fecha)

	var r0 *entity.Formulario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Formulario, error)); ok {
		return rf(ctx, fecha)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Formulario); ok {
		r0 = rf(ctx, fecha)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Formulario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fecha)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFormByName provides a mock function with given fields: ctx, nombre
func (_m *MockRepository) GetFormByName(ctx context.Context, nombre string) (*entity.Formulario, error) {
	ret := _m.Called(ctx, nombre)

	var r0 *entity.Formulario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Formulario, error)); ok {
		return rf(ctx, nombre)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Formulario); ok {
		r0 = rf(ctx, nombre)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Formulario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nombre)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFormByVersion provides a mock function with given fields: ctx, version
func (_m *MockRepository) GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error) {
	ret := _m.Called(ctx, version)

	var r0 *entity.Formulario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Formulario, error)); ok {
		return rf(ctx, version)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Formulario); ok {
		r0 = rf(ctx, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Formulario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForms provides a mock function with given fields: ctx
func (_m *MockRepository) GetForms(ctx context.Context) ([]entity.Formulario, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Formulario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Formulario, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Formulario); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Formulario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFormsById provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetFormsById(ctx context.Context, id int64) (*entity.Formulario, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Formulario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.Formulario, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Formulario); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Formulario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFromEtapas provides a mock function with given fields: ctx
func (_m *MockRepository) GetFromEtapas(ctx context.Context) (*entity.Formulario, error) {
	ret := _m.Called(ctx)

	var r0 *entity.Formulario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*entity.Formulario, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *entity.Formulario); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Formulario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFromUsers provides a mock function with given fields: ctx
func (_m *MockRepository) GetFromUsers(ctx context.Context) (*entity.Formulario, error) {
	ret := _m.Called(ctx)

	var r0 *entity.Formulario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*entity.Formulario, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *entity.Formulario); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Formulario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObraPisos provides a mock function with given fields: ctx, obraID
func (_m *MockRepository) GetObraPisos(ctx context.Context, obraID int64) ([]entity.PisoObra, error) {
	ret := _m.Called(ctx, obraID)

	var r0 []entity.PisoObra
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]entity.PisoObra, error)); ok {
		return rf(ctx, obraID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entity.PisoObra); ok {
		r0 = rf(ctx, obraID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.PisoObra)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, obraID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObrabyName provides a mock function with given fields: ctx, name
func (_m *MockRepository) GetObrabyName(ctx context.Context, name string) (*entity.Obra, error) {
	ret := _m.Called(ctx, name)

	var r0 *entity.Obra
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Obra, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Obra); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Obra)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPisobyNumber provides a mock function with given fields: ctx, number
func (_m *MockRepository) GetPisobyNumber(ctx context.Context, number int64) (*entity.Piso, error) {
	ret := _m.Called(ctx, number)

	var r0 *entity.Piso
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.Piso, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Piso); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Piso)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *MockRepository) GetUserByEmail(ctx context.Context, email string) (*entity.Usuario, error) {
	ret := _m.Called(ctx, email)

	var r0 *entity.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Usuario, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Usuario); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetUserById(ctx context.Context, id int) (*entity.Usuario, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*entity.Usuario, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Usuario); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserRoles provides a mock function with given fields: ctx, userID
func (_m *MockRepository) GetUserRoles(ctx context.Context, userID int64) ([]entity.UsarioRol, error) {
	ret := _m.Called(ctx, userID)

	var r0 []entity.UsarioRol
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]entity.UsarioRol, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entity.UsarioRol); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.UsarioRol)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx
func (_m *MockRepository) GetUsers(ctx context.Context) ([]entity.Usuario, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Usuario, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Usuario); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockRepository) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveCheck provides a mock function with given fields: ctx, estado, observaciones, version, fecha
func (_m *MockRepository) SaveCheck(ctx context.Context, estado string, observaciones string, version int, fecha string) error {
	ret := _m.Called(ctx, estado, observaciones, version, fecha)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, string) error); ok {
		r0 = rf(ctx, estado, observaciones, version, fecha)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveCheckForm provides a mock function with given fields: ctx, checkID, formularioID
func (_m *MockRepository) SaveCheckForm(ctx context.Context, checkID int64, formularioID int64) error {
	ret := _m.Called(ctx, checkID, formularioID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, checkID, formularioID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveControl provides a mock function with given fields: ctx, descripcion, tipo
func (_m *MockRepository) SaveControl(ctx context.Context, descripcion string, tipo string) error {
	ret := _m.Called(ctx, descripcion, tipo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, descripcion, tipo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveEtapa provides a mock function with given fields: ctx, nombre
func (_m *MockRepository) SaveEtapa(ctx context.Context, nombre string) error {
	ret := _m.Called(ctx, nombre)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nombre)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveFrom provides a mock function with given fields: ctx, nombre, informacion, version, fecha, etapa_id, usuario_id
func (_m *MockRepository) SaveFrom(ctx context.Context, nombre string, informacion string, version string, fecha string, etapa_id int, usuario_id int) error {
	ret := _m.Called(ctx, nombre, informacion, version, fecha, etapa_id, usuario_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, int, int) error); ok {
		r0 = rf(ctx, nombre, informacion, version, fecha, etapa_id, usuario_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveObra provides a mock function with given fields: ctx, nombre
func (_m *MockRepository) SaveObra(ctx context.Context, nombre string) error {
	ret := _m.Called(ctx, nombre)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nombre)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveObraPiso provides a mock function with given fields: ctx, obraID, pisoID
func (_m *MockRepository) SaveObraPiso(ctx context.Context, obraID int64, pisoID int64) error {
	ret := _m.Called(ctx, obraID, pisoID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, obraID, pisoID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SavePiso provides a mock function with given fields: ctx, number
func (_m *MockRepository) SavePiso(ctx context.Context, number int64) error {
	ret := _m.Called(ctx, number)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, number)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUser provides a mock function with given fields: ctx, email, username, password
func (_m *MockRepository) SaveUser(ctx context.Context, email string, username string, password string) error {
	ret := _m.Called(ctx, email, username, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, email, username, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockRepository) SaveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

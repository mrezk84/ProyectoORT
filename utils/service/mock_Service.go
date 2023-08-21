package service

import (
	"context"

	"proyectoort/utils/entity"
	"proyectoort/utils/models"

	"github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// AddUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockService) AddUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginUser provides a mock function with given fields: ctx, email, password
func (_m *MockService) LoginUser(ctx context.Context, email string, password string) (*models.Usuario, error) {
	ret := _m.Called(ctx, email, password)

	var r0 *models.Usuario
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Usuario); ok {
		r0 = rf(ctx, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Usuario)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, email, name, password
func (_m *MockService) RegisterUser(ctx context.Context, email string, username string, password string) error {
	ret := _m.Called(ctx, email, username, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, email, username, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *MockService) GetUserByEmail(ctx context.Context, email string) (*entity.Usuario, error) {
	ret := _m.Called(ctx, email)

	var r0 *entity.Usuario
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Usuario); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Usuario)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockService) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockService) RegisterFrom(ctx context.Context, nombre string, informacion string, version string, fecha string) error {
	ret := _m.Called(ctx, informacion, nombre, version, fecha)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, informacion, nombre, version, fecha)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockService) GetFormByVersion(ctx context.Context, version string) (*entity.Formulario, error) {
	ret := _m.Called(ctx, version)

	var r0 *entity.Formulario
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Formulario); ok {
		r0 = rf(ctx, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Formulario)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, version)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *MockService) GetFormsById(ctx context.Context, id int64) (*entity.Formulario, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Formulario
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Formulario); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Formulario)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}
func (_m *MockService) GetControls(ctx context.Context) (*models.Control, error) {
	ret := _m.Called(ctx)

	var r0 *models.Control
	if rf, ok := ret.Get(0).(func(context.Context) *models.Control); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Control)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockService) GetControlById(ctx context.Context, id int) (*entity.Control, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Control
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Control); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Control)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *MockService) RegisterCheck(ctx context.Context, estado string, observaciones string, version int, fecha string) error {
	ret := _m.Called(ctx, estado, observaciones, version, fecha)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, string) error); ok {
		r0 = rf(ctx, estado, observaciones, version, fecha)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockService(t mockConstructorTestingTNewMockService) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

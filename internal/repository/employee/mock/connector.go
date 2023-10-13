package mock

import (
	"context"
	"employee/internal/model"
	"github.com/stretchr/testify/mock"
)

type DBMock struct {
	mock.Mock
}

func (m *DBMock) CreateEmployee(ctx context.Context, employee *model.Employee) (int, error) {
	ret := m.Called(ctx, employee)
	return ret.Get(0).(int), ret.Error(1)
}

func (m *DBMock) GetEmployees(ctx context.Context) ([]*model.Employee, error) {
	ret := m.Called(ctx)
	return ret.Get(0).([]*model.Employee), ret.Error(1)
}

func (m *DBMock) GetEmployeeByID(ctx context.Context, employeeID int) (*model.Employee, error) {
	ret := m.Called(ctx, employeeID)
	return ret.Get(0).(*model.Employee), ret.Error(1)
}

func (m *DBMock) UpdateEmployee(ctx context.Context, employee *model.Employee) error {
	ret := m.Called(ctx, employee)
	return ret.Error(0)
}

func (m *DBMock) DeleteEmployee(ctx context.Context, employeeID int) error {
	ret := m.Called(ctx, employeeID)
	return ret.Error(0)
}

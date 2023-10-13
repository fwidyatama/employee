package mock

import (
	"context"
	"employee/internal/transport"
	"github.com/stretchr/testify/mock"
)

type EmployeeUseCaseMock struct {
	mock.Mock
}

func (m *EmployeeUseCaseMock) GetEmployees(ctx context.Context) (*transport.ListEmployees, error) {
	args := m.Called(ctx)

	return args.Get(0).(*transport.ListEmployees), args.Error(1)
}

func (m *EmployeeUseCaseMock) CreateEmployee(ctx context.Context, payload *transport.CreateEmployeeReq) (*transport.EmployeeRes, error) {
	args := m.Called(ctx, payload)

	return args.Get(0).(*transport.EmployeeRes), args.Error(1)
}

func (m *EmployeeUseCaseMock) GetEmployeeByID(ctx context.Context, employeeID int) (*transport.EmployeeRes, error) {
	args := m.Called(ctx, employeeID)

	return args.Get(0).(*transport.EmployeeRes), args.Error(1)
}

func (m *EmployeeUseCaseMock) UpdateEmployee(ctx context.Context, payload *transport.UpdateEmployeeReq) error {
	args := m.Called(ctx, payload)

	return args.Error(0)
}

func (m *EmployeeUseCaseMock) DeleteEmployee(ctx context.Context, employeeID int) error {
	args := m.Called(ctx, employeeID)

	return args.Error(0)
}

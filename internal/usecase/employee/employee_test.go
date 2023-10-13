package employee

import (
	"context"
	"database/sql"
	"employee/internal/model"
	employeeRepoMock "employee/internal/repository/employee/mock"
	"employee/internal/transport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateEmployee(t *testing.T) {

	payload := &transport.CreateEmployeeReq{
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test.com",
		HireDate:  "2023-05-03",
	}

	testCases := []struct {
		name      string
		payload   *transport.CreateEmployeeReq
		buildStub func(
			employeeRepo *employeeRepoMock.DBMock,
		)
		checkReturn func(user *transport.EmployeeRes, err error)
	}{

		{
			name:    "error when create employee",
			payload: payload,
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("CreateEmployee", mock.Anything, mock.Anything).Return(0, sql.ErrConnDone)
			},
			checkReturn: func(user *transport.EmployeeRes, err error) {
				assert.Nil(t, user)
				assert.Error(t, err)
			},
		},
		{
			name:    "success when create employee",
			payload: payload,
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("CreateEmployee", mock.Anything, mock.Anything).Return(1, nil)
			},
			checkReturn: func(user *transport.EmployeeRes, err error) {
				assert.NotNil(t, user)
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			employeeRepository := new(employeeRepoMock.DBMock)
			tc.buildStub(employeeRepository)

			u := NewUseCaseEmployee(employeeRepository)
			result, err := u.CreateEmployee(context.TODO(), tc.payload)

			tc.checkReturn(result, err)

		})
	}
}

func TestGetEmployee(t *testing.T) {

	mockEmployeesResult := []*model.Employee{
		{
			ID:        1,
			FirstName: "test",
			LastName:  "test",
			Email:     "test@mail.com",
			HireDate:  "2023-05-01",
		},
	}

	testCases := []struct {
		name      string
		buildStub func(
			employeeRepo *employeeRepoMock.DBMock,
		)
		checkReturn func(employees *transport.ListEmployees, err error)
	}{

		{
			name: "error when get all employee",
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployees", mock.Anything, mock.Anything).Return([]*model.Employee{}, sql.ErrConnDone)
			},
			checkReturn: func(employees *transport.ListEmployees, err error) {
				assert.Nil(t, employees)
				assert.Error(t, err)
			},
		},
		{
			name: "success when get employee",
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployees", mock.Anything, mock.Anything).Return(mockEmployeesResult, nil)
			},
			checkReturn: func(employees *transport.ListEmployees, err error) {
				assert.NotNil(t, employees)
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			employeeRepository := new(employeeRepoMock.DBMock)
			tc.buildStub(employeeRepository)

			u := NewUseCaseEmployee(employeeRepository)
			result, err := u.GetEmployees(context.TODO())

			tc.checkReturn(result, err)

		})
	}
}

func TestGetEmployeeByID(t *testing.T) {

	mockEmployeesResult := &model.Employee{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Email:     "test@mail.com",
		HireDate:  "2023-05-01",
	}

	testCases := []struct {
		name       string
		employeeID int
		buildStub  func(
			employeeRepo *employeeRepoMock.DBMock,
		)
		checkReturn func(employees *transport.EmployeeRes, err error)
	}{

		{
			name: "error when get all employee",
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(&model.Employee{}, sql.ErrConnDone)
			},
			checkReturn: func(employees *transport.EmployeeRes, err error) {
				assert.Nil(t, employees)
				assert.Error(t, err)
			},
		},
		{
			name: "success when get employee",
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(mockEmployeesResult, nil)
			},
			checkReturn: func(employees *transport.EmployeeRes, err error) {
				assert.NotNil(t, employees)
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			employeeRepository := new(employeeRepoMock.DBMock)
			tc.buildStub(employeeRepository)

			u := NewUseCaseEmployee(employeeRepository)
			result, err := u.GetEmployeeByID(context.TODO(), tc.employeeID)

			tc.checkReturn(result, err)

		})
	}
}

func TestUpdateEmployee(t *testing.T) {

	payload := &transport.UpdateEmployeeReq{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test.com",
		HireDate:  "2023-05-03",
	}

	mockEmployeesResult := &model.Employee{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Email:     "test@mail.com",
		HireDate:  "2023-05-01",
	}

	testCases := []struct {
		name      string
		payload   *transport.UpdateEmployeeReq
		buildStub func(
			employeeRepo *employeeRepoMock.DBMock,
		)
		checkReturn func(err error)
	}{

		{
			name:    "error when get employee detail",
			payload: payload,
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(&model.Employee{}, sql.ErrConnDone)
			},
			checkReturn: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "error when update employee",
			payload: payload,
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(mockEmployeesResult, nil)
				employeeRepoMock.On("UpdateEmployee", mock.Anything, mock.Anything).Return(sql.ErrConnDone)
			},
			checkReturn: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "success when update employee",
			payload: payload,
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(mockEmployeesResult, nil)
				employeeRepoMock.On("UpdateEmployee", mock.Anything, mock.Anything).Return(nil)
			},
			checkReturn: func(err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			employeeRepository := new(employeeRepoMock.DBMock)
			tc.buildStub(employeeRepository)

			u := NewUseCaseEmployee(employeeRepository)
			err := u.UpdateEmployee(context.TODO(), tc.payload)

			tc.checkReturn(err)

		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	mockEmployeesResult := &model.Employee{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Email:     "test@mail.com",
		HireDate:  "2023-05-01",
	}

	testCases := []struct {
		name       string
		employeeID int
		buildStub  func(
			employeeRepo *employeeRepoMock.DBMock,
		)
		checkReturn func(err error)
	}{

		{
			name:       "error when get employee detail",
			employeeID: 1,
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(&model.Employee{}, sql.ErrConnDone)
			},
			checkReturn: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:       "error when update employee",
			employeeID: 1,
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(mockEmployeesResult, nil)
				employeeRepoMock.On("DeleteEmployee", mock.Anything, mock.Anything).Return(sql.ErrConnDone)
			},
			checkReturn: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:       "success when update employee",
			employeeID: 1,
			buildStub: func(employeeRepoMock *employeeRepoMock.DBMock) {
				employeeRepoMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(mockEmployeesResult, nil)
				employeeRepoMock.On("DeleteEmployee", mock.Anything, mock.Anything).Return(nil)
			},
			checkReturn: func(err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			employeeRepository := new(employeeRepoMock.DBMock)
			tc.buildStub(employeeRepository)

			u := NewUseCaseEmployee(employeeRepository)
			err := u.DeleteEmployee(context.TODO(), tc.employeeID)

			tc.checkReturn(err)

		})
	}
}

package employee

import (
	"database/sql"
	"employee/internal/config"
	"employee/internal/transport"
	employeeUCMock "employee/internal/usecase/employee/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateEmployee(t *testing.T) {

	completePayload := `{
    "first_name":"farid",
    "last_name":"widyatama",
    "email":"email@mil.com",
    "hire_date":"2023-04-05"
	}`

	incompletePayload := `
	{
    "first_name":"farid",
    "last_name":"widyatama",
    "hire_date":"2023-04-05"
	}`

	invalidPayload := `{invalid json}`

	testCases := []struct {
		name      string
		payload   string
		buildStub func(
			employeeUCMock *employeeUCMock.EmployeeUseCaseMock,
		)
		checkReturn func(resp *httptest.ResponseRecorder)
	}{
		{
			name:    "failed when marshall json",
			payload: invalidPayload,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, resp.Code)
			},
		},
		{
			name:    "failed when doing validation",
			payload: incompletePayload,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, resp.Code)
			},
		},
		{
			name:    "failed when create employee",
			payload: completePayload,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("CreateEmployee", mock.Anything, mock.Anything).Return(&transport.EmployeeRes{}, sql.ErrConnDone)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, resp.Code)
			},
		},
		{
			name:    "success create employee",
			payload: completePayload,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("CreateEmployee", mock.Anything, mock.Anything).Return(&transport.EmployeeRes{}, nil)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, resp.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodPost, "/employees", strings.NewReader(tc.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			employeeUC := new(employeeUCMock.EmployeeUseCaseMock)
			tc.buildStub(employeeUC)

			cfg := new(config.Config)
			h := NewEmployeeHandler(employeeUC, *cfg)
			_ = h.CreateEmployee(c)

			tc.checkReturn(rec)
		})
	}
}

func TestGetEmployee(t *testing.T) {

	mockEmployeeResult := &transport.ListEmployees{
		Employees: []*transport.EmployeeRes{
			{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				Email:     "test@mail.com",
				HireDate:  "2023-05-01",
			},
		}}

	testCases := []struct {
		name      string
		buildStub func(
			employeeUCMock *employeeUCMock.EmployeeUseCaseMock,
		)
		checkReturn func(resp *httptest.ResponseRecorder)
	}{

		{
			name: "failed when get employee",
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("GetEmployees", mock.Anything).Return(&transport.ListEmployees{}, sql.ErrConnDone)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, resp.Code)
			},
		},
		{
			name: "success create employee",
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("GetEmployees", mock.Anything).Return(mockEmployeeResult, nil)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, resp.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/employees", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			employeeUC := new(employeeUCMock.EmployeeUseCaseMock)
			tc.buildStub(employeeUC)

			cfg := new(config.Config)
			h := NewEmployeeHandler(employeeUC, *cfg)
			_ = h.GetEmployee(c)

			tc.checkReturn(rec)
		})
	}
}

func TestGetEmployeeByID(t *testing.T) {

	mockEmployeeResult := &transport.EmployeeRes{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Email:     "test@mail.com",
		HireDate:  "2023-05-01",
	}

	testCases := []struct {
		name      string
		buildStub func(
			employeeUCMock *employeeUCMock.EmployeeUseCaseMock,
		)
		checkReturn func(resp *httptest.ResponseRecorder)
	}{

		{
			name: "failed when get employee by id",
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(&transport.EmployeeRes{}, sql.ErrConnDone)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, resp.Code)
			},
		},
		{
			name: "success create employee",
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(mockEmployeeResult, nil)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, resp.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/employees/:employee_id", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			employeeUC := new(employeeUCMock.EmployeeUseCaseMock)
			tc.buildStub(employeeUC)

			cfg := new(config.Config)
			h := NewEmployeeHandler(employeeUC, *cfg)
			_ = h.GetEmployeeByID(c)

			tc.checkReturn(rec)
		})
	}
}

func TestUpdateEmployee(t *testing.T) {

	completePayload := `{
    "first_name":"farid",
    "last_name":"widyatama",
    "email":"email@mil.com",
    "hire_date":"2023-04-05"
	}`

	incompletePayload := `
	{
    "first_name":"farid",
    "last_name":"widyatama",
    "hire_date":"2023-04-05"
	}`

	invalidPayload := `{invalid json}`

	testCases := []struct {
		name      string
		payload   string
		buildStub func(
			employeeUCMock *employeeUCMock.EmployeeUseCaseMock,
		)
		checkReturn func(resp *httptest.ResponseRecorder)
	}{
		{
			name:    "failed when marshall json",
			payload: invalidPayload,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, resp.Code)
			},
		},
		{
			name:    "failed when doing validation",
			payload: incompletePayload,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, resp.Code)
			},
		},
		{
			name:    "failed when create employee",
			payload: completePayload,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("CreateEmployee", mock.Anything, mock.Anything).Return(&transport.EmployeeRes{}, sql.ErrConnDone)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, resp.Code)
			},
		},
		{
			name:    "success create employee",
			payload: completePayload,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("CreateEmployee", mock.Anything, mock.Anything).Return(&transport.EmployeeRes{}, nil)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, resp.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodPut, "/employees", strings.NewReader(tc.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			employeeUC := new(employeeUCMock.EmployeeUseCaseMock)
			tc.buildStub(employeeUC)

			cfg := new(config.Config)
			h := NewEmployeeHandler(employeeUC, *cfg)
			_ = h.CreateEmployee(c)

			tc.checkReturn(rec)
		})
	}
}

func TestDeleteEmployee(t *testing.T) {

	testCases := []struct {
		name       string
		employeeID int
		buildStub  func(
			employeeUCMock *employeeUCMock.EmployeeUseCaseMock,
		)
		checkReturn func(resp *httptest.ResponseRecorder)
	}{
		{
			name:       "failed delete employee",
			employeeID: 1,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("DeleteEmployee", mock.Anything, mock.Anything).Return(sql.ErrConnDone)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, resp.Code)
			},
		},
		{
			name:       "success delete employee",
			employeeID: 1,
			buildStub: func(employeeUCMock *employeeUCMock.EmployeeUseCaseMock) {
				employeeUCMock.On("DeleteEmployee", mock.Anything, mock.Anything).Return(nil)
			},
			checkReturn: func(resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, resp.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodDelete, "/employees/:employee_id", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			employeeUC := new(employeeUCMock.EmployeeUseCaseMock)
			tc.buildStub(employeeUC)

			cfg := new(config.Config)
			h := NewEmployeeHandler(employeeUC, *cfg)
			_ = h.DeleteEmployee(c)

			tc.checkReturn(rec)
		})
	}
}

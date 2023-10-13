package employee

import (
	"context"
	"employee/internal/config"
	"employee/internal/response"
	"employee/internal/transport"
	"employee/internal/usecase/employee"
	"errors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var (
	logger = log.WithField("handler", "handler.employee")
)

type Handler struct {
	uc  employee.UseCaseEmployee
	cfg config.Config
}

func NewEmployeeHandler(employeeUC employee.UseCaseEmployee, cfg config.Config) *Handler {
	return &Handler{uc: employeeUC, cfg: cfg}
}

func (h *Handler) CreateEmployee(c echo.Context) error {
	hLog := logger.WithField("handler", "CreateEmployee")

	ctx := context.Background()

	payload := new(transport.CreateEmployeeReq)

	if err := c.Bind(payload); err != nil {
		hLog.Errorf("echo bind got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	if err := transport.ValidateStruct(payload); err != nil {
		hLog.Errorf("error when validate body, got %s", err)
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	res, err := h.uc.CreateEmployee(ctx, payload)
	if err != nil {
		hLog.Errorf("error when call u.CreateEmployee got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, res)
}

func (h *Handler) GetEmployee(c echo.Context) error {
	hLog := logger.WithField("handler", "GetEmployee")

	ctx := context.Background()

	res, err := h.uc.GetEmployees(ctx)
	if err != nil {
		hLog.Errorf("error when call u.GetEmployees got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, res)
}

func (h *Handler) GetEmployeeByID(c echo.Context) error {
	hLog := logger.WithField("handler", "GetEmployeeByID")

	ctx := context.Background()

	employeeIDStr := c.Param("employee_id")
	employeeID, _ := strconv.Atoi(employeeIDStr)

	res, err := h.uc.GetEmployeeByID(ctx, employeeID)

	if res == nil {
		err := errors.New("employee not found")
		hLog.Errorf("error when call u.GetEmployeeByID got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusNotFound)
	}

	if err != nil {
		hLog.Errorf("error when call u.GetEmployeeByID got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, res)
}

func (h *Handler) UpdateEmployee(c echo.Context) error {
	hLog := logger.WithField("handler", "UpdateEmployee")

	ctx := context.Background()

	employeeIDStr := c.Param("employee_id")
	employeeID, _ := strconv.Atoi(employeeIDStr)

	payload := new(transport.UpdateEmployeeReq)
	payload.ID = employeeID

	if err := c.Bind(payload); err != nil {
		hLog.Errorf("echo bind got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	if err := transport.ValidateStruct(payload); err != nil {
		hLog.Errorf("error when validate body, got %s", err)
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	err := h.uc.UpdateEmployee(ctx, payload)

	if err != nil && err.Error() == "employee not found" {
		hLog.Errorf("employee not found got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusNotFound)
	}

	if err != nil {
		hLog.Errorf("error when call u.UpdateEmployee got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, nil)
}

func (h *Handler) DeleteEmployee(c echo.Context) error {
	hLog := logger.WithField("handler", "DeleteEmployee")

	ctx := context.Background()

	employeeIDStr := c.Param("employee_id")
	employeeID, _ := strconv.Atoi(employeeIDStr)

	err := h.uc.DeleteEmployee(ctx, employeeID)

	if err != nil {
		hLog.Errorf("error when call uc.DeleteEmployee got %s", err.Error())
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, nil)
}

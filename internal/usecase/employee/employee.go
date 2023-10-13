package employee

import (
	"context"
	"employee/internal/model"
	eRepo "employee/internal/repository/employee"
	"employee/internal/transport"
	"errors"
	log "github.com/sirupsen/logrus"
)

var (
	logger = log.WithField("useCase", "useCase.Employee")
)

type UseCaseEmployee interface {
	CreateEmployee(ctx context.Context, payload *transport.CreateEmployeeReq) (*transport.EmployeeRes, error)
	GetEmployees(ctx context.Context) (*transport.ListEmployees, error)
	GetEmployeeByID(ctx context.Context, employeeID int) (*transport.EmployeeRes, error)
	UpdateEmployee(ctx context.Context, payload *transport.UpdateEmployeeReq) error
	DeleteEmployee(ctx context.Context, employeeID int) error
}

type useCaseEmployee struct {
	employeeRepo eRepo.UserRepo
}

func NewUseCaseEmployee(employeeRepo eRepo.UserRepo) UseCaseEmployee {
	return &useCaseEmployee{employeeRepo: employeeRepo}
}

func (u *useCaseEmployee) CreateEmployee(ctx context.Context, payload *transport.CreateEmployeeReq) (*transport.EmployeeRes, error) {
	uLog := logger.WithContext(ctx).WithField("function", "Register")

	employee := &model.Employee{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		HireDate:  payload.HireDate,
	}

	currentID, err := u.employeeRepo.CreateEmployee(ctx, employee)
	if err != nil {
		uLog.Errorf("error when call employeeRepo.CreateEmployee got %s", err.Error())
		return nil, err
	}

	result := &transport.EmployeeRes{
		ID:        currentID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		HireDate:  payload.HireDate,
	}

	return result, nil
}

func (u *useCaseEmployee) GetEmployees(ctx context.Context) (*transport.ListEmployees, error) {
	uLog := logger.WithContext(ctx).WithField("function", "GetEmployees")

	employees, err := u.employeeRepo.GetEmployees(ctx)
	if err != nil {
		uLog.Errorf("error when call employeeRepo.GetEmployees got %s", err.Error())
		return nil, err
	}

	employeesResData := make([]*transport.EmployeeRes, 0)
	for _, employee := range employees {
		emp := &transport.EmployeeRes{
			ID:        employee.ID,
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
			Email:     employee.Email,
			HireDate:  employee.HireDate,
		}
		employeesResData = append(employeesResData, emp)
	}

	employeesRes := &transport.ListEmployees{Employees: employeesResData}

	return employeesRes, nil
}

func (u *useCaseEmployee) GetEmployeeByID(ctx context.Context, employeeID int) (*transport.EmployeeRes, error) {
	uLog := logger.WithContext(ctx).WithField("function", "GetEmployeeByID")

	employee, err := u.employeeRepo.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		uLog.Errorf("error when call employeeRepo.GetEmployeeByID got %s", err.Error())
		return nil, err
	}

	if employee != nil {
		employeeRes := &transport.EmployeeRes{
			ID:        employee.ID,
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
			Email:     employee.Email,
			HireDate:  employee.HireDate,
		}
		return employeeRes, nil
	}

	return nil, nil

}

func (u *useCaseEmployee) UpdateEmployee(ctx context.Context, payload *transport.UpdateEmployeeReq) error {
	uLog := logger.WithContext(ctx).WithField("function", "UpdateEmployee")

	employee, err := u.employeeRepo.GetEmployeeByID(ctx, payload.ID)
	if err != nil {
		uLog.Errorf("error when call employeeRepo.GetEmployeeByID got %s", err.Error())
		return err
	}

	if employee == nil {
		err = errors.New("employee not found")
		uLog.Errorf("error when call employeeRepo.GetEmployeeByID got %s", err.Error())
		return err
	}

	employeePayload := &model.Employee{
		ID:        payload.ID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		HireDate:  payload.HireDate,
	}

	err = u.employeeRepo.UpdateEmployee(ctx, employeePayload)
	if err != nil {
		uLog.Errorf("error when call employeeRepo.UpdateEmployee got %s", err.Error())
		return err
	}

	return nil

}

func (u *useCaseEmployee) DeleteEmployee(ctx context.Context, employeeID int) error {
	uLog := logger.WithContext(ctx).WithField("function", "WithField")

	employee, err := u.employeeRepo.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		uLog.Errorf("error when call employeeRepo.GetEmployeeByID got %s", err.Error())
		return err
	}

	if employee == nil {
		err = errors.New("employee not found")
		uLog.Errorf("error when call employeeRepo.GetEmployeeByID got %s", err.Error())
		return err
	}

	err = u.employeeRepo.DeleteEmployee(ctx, employeeID)
	if err != nil {
		uLog.Errorf("error when call employeeRepo.DeleteEmployee got %s", err.Error())
		return err
	}
	return nil

}

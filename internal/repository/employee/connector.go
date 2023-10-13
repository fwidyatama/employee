package employee

import (
	"context"
	"database/sql"
	"employee/internal/model"
	log "github.com/sirupsen/logrus"
)

var (
	logRepo = log.WithField("package", "repository.employee")
)

type UserRepo interface {
	CreateEmployee(ctx context.Context, employee *model.Employee) (int, error)
	GetEmployees(ctx context.Context) ([]*model.Employee, error)
	GetEmployeeByID(ctx context.Context, employeeID int) (*model.Employee, error)
	UpdateEmployee(ctx context.Context, employee *model.Employee) error
	DeleteEmployee(ctx context.Context, employeeID int) error
}

type userRepo struct {
	sqlConn *sql.DB
}

func NewRepoUser(sqlConn *sql.DB) UserRepo {
	return &userRepo{sqlConn: sqlConn}
}

func (u *userRepo) CreateEmployee(ctx context.Context, employee *model.Employee) (int, error) {
	rLog := logRepo.WithField("function", "CreateEmployee")

	var currentInsertedID int

	query := `INSERT INTO employees 
		(first_name, last_name,email,hire_date )
		values ($1, $2, $3, $4) returning id`

	values := []interface{}{
		employee.FirstName,
		employee.LastName,
		employee.Email,
		employee.HireDate,
	}

	err := u.sqlConn.QueryRowContext(ctx, query, values...).Scan(&currentInsertedID)
	if err != nil {
		rLog.Errorf("error when create employee got: %s", err.Error())
		return 0, err
	}

	return currentInsertedID, nil
}

func (u *userRepo) GetEmployees(ctx context.Context) ([]*model.Employee, error) {
	rLog := logRepo.WithField("function", "GetEmployee")

	var employees []*model.Employee

	query := `select id, first_name, last_name, email, hire_date from employees order by id DESC`

	rows, err := u.sqlConn.QueryContext(ctx, query)
	if err != nil {
		rLog.Errorf("error when get employees got: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		temp := &model.Employee{}
		err := rows.Scan(&temp.ID, &temp.FirstName, &temp.LastName, &temp.Email, &temp.HireDate)
		if err != nil {
			rLog.Errorf("error when scan: %s", err.Error())
			return nil, err
		}

		employees = append(employees, temp)

	}

	return employees, nil
}

func (u *userRepo) GetEmployeeByID(ctx context.Context, employeeID int) (*model.Employee, error) {
	rLog := logRepo.WithField("function", "GetEmployeeByID")

	employees := &model.Employee{}

	query := `select id, first_name, last_name, email, hire_date from employees where id = $1`

	row := u.sqlConn.QueryRowContext(ctx, query, employeeID)

	err := row.Scan(&employees.ID, &employees.FirstName, &employees.LastName, &employees.Email, &employees.HireDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		rLog.Errorf("error when scan: %s", err.Error())
		return nil, err
	}

	return employees, nil
}

func (u *userRepo) UpdateEmployee(ctx context.Context, employee *model.Employee) error {
	rLog := logRepo.WithField("function", "GetEmployeeByID")

	values := []interface{}{employee.FirstName, employee.LastName, employee.Email, employee.HireDate, employee.ID}

	query := `UPDATE employees  SET first_name=$1, last_name=$2, email=$3, hire_date=$4 where id = $5`

	_, err := u.sqlConn.ExecContext(ctx, query, values...)
	if err != nil {
		rLog.Error(err)
		return err
	}

	return nil
}

func (u *userRepo) DeleteEmployee(ctx context.Context, employeeID int) error {
	rLog := logRepo.WithField("function", "DeleteEmployee")

	query := `DELETE FROM employees WHERE id = $1`

	_, err := u.sqlConn.ExecContext(ctx, query, employeeID)
	if err != nil {
		rLog.Error(err)
		return err
	}

	return nil
}

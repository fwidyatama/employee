package employee

import (
	"context"
	"database/sql"
	"employee/internal/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	query := `INSERT INTO employees 
		(first_name, last_name,email,hire_date )
		values ($1, $2, $3, $4) returning id`

	employee := &model.Employee{
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test",
		HireDate:  "2023-05-02",
	}

	testCase := []struct {
		name        string
		payload     *model.Employee
		buildStub   func(mock sqlmock.Sqlmock)
		checkReturn func(resultID int, err error)
	}{

		{
			name:    "error connection when create employee",
			payload: employee,
			buildStub: func(mock sqlmock.Sqlmock) {
				runQueryCount := regexp.QuoteMeta(query)
				mock.ExpectQuery(runQueryCount).WillReturnError(sql.ErrConnDone)
			},
			checkReturn: func(resultID int, err error) {
				assert.Error(t, err)
				assert.Zero(t, resultID)
			},
		},
		{
			name:    "success",
			payload: employee,
			buildStub: func(mock sqlmock.Sqlmock) {
				runQueryCount := regexp.QuoteMeta(query)
				mock.ExpectQuery(runQueryCount).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
			},
			checkReturn: func(resultID int, err error) {
				assert.NoError(t, err)
				assert.NotZero(t, resultID)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			defer db.Close()

			tc.buildStub(mock)

			repo := NewRepoUser(db)

			result, err := repo.CreateEmployee(context.TODO(), tc.payload)

			tc.checkReturn(result, err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetEmployees(t *testing.T) {
	query := `select id, first_name, last_name, email, hire_date from employees order by id DESC`

	testCase := []struct {
		name        string
		buildStub   func(mock sqlmock.Sqlmock)
		checkReturn func(result []*model.Employee, err error)
	}{

		{
			name: "error connection when get employee",
			buildStub: func(mock sqlmock.Sqlmock) {
				runQueryCount := regexp.QuoteMeta(query)
				mock.ExpectQuery(runQueryCount).WillReturnError(sql.ErrConnDone)
			},
			checkReturn: func(result []*model.Employee, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
			},
		},
		{
			name: "success",
			buildStub: func(mock sqlmock.Sqlmock) {
				runQueryCount := regexp.QuoteMeta(query)
				mock.ExpectQuery(runQueryCount).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "hire_date"}).
					AddRow("1", "test", "test", "test@mail.com", "2023-05-03"))
			},
			checkReturn: func(result []*model.Employee, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			defer db.Close()

			tc.buildStub(mock)

			repo := NewRepoUser(db)

			result, err := repo.GetEmployees(context.TODO())

			tc.checkReturn(result, err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetEmployeeByID(t *testing.T) {
	query := `select id, first_name, last_name, email, hire_date from employees where id = $1`

	testCase := []struct {
		name        string
		employeeID  int
		buildStub   func(mock sqlmock.Sqlmock)
		checkReturn func(result *model.Employee, err error)
	}{

		{
			name: "error connection when get employee by id",
			buildStub: func(mock sqlmock.Sqlmock) {
				runQueryCount := regexp.QuoteMeta(query)
				mock.ExpectQuery(runQueryCount).WillReturnError(sql.ErrConnDone)
			},
			checkReturn: func(result *model.Employee, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
			},
		},
		{
			name:       "success",
			employeeID: 1,
			buildStub: func(mock sqlmock.Sqlmock) {

				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "hire_date"}).
					AddRow(1, "test", "test", "test@mail.com", "2023-05-03")
				runQuery := regexp.QuoteMeta(query)

				mock.ExpectQuery(runQuery).WillReturnRows(rows)
			},
			checkReturn: func(result *model.Employee, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			defer db.Close()

			tc.buildStub(mock)

			repo := NewRepoUser(db)

			result, err := repo.GetEmployeeByID(context.TODO(), tc.employeeID)

			tc.checkReturn(result, err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUpdateEmployee(t *testing.T) {
	query := `UPDATE employees  SET first_name=$1, last_name=$2, email=$3, hire_date=$4 where id = $5`

	employee := &model.Employee{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test",
		HireDate:  "2023-05-02",
	}

	testCase := []struct {
		name        string
		model       *model.Employee
		buildStub   func(mock sqlmock.Sqlmock)
		checkReturn func(err error)
	}{

		{
			name:  "error connection when update employee",
			model: employee,
			buildStub: func(mock sqlmock.Sqlmock) {
				runQueryCount := regexp.QuoteMeta(query)
				mock.ExpectExec(runQueryCount).WillReturnError(sql.ErrConnDone)
			},
			checkReturn: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "success",
			model: employee,
			buildStub: func(mock sqlmock.Sqlmock) {

				runQuery := regexp.QuoteMeta(query)
				mock.ExpectExec(runQuery).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			checkReturn: func(err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			defer db.Close()

			tc.buildStub(mock)

			repo := NewRepoUser(db)

			err = repo.UpdateEmployee(context.TODO(), tc.model)

			tc.checkReturn(err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	query := `DELETE FROM employees WHERE id = $1`

	testCase := []struct {
		name        string
		employeeID  int
		buildStub   func(mock sqlmock.Sqlmock)
		checkReturn func(err error)
	}{

		{
			name:       "error connection when delete employee",
			employeeID: 1,
			buildStub: func(mock sqlmock.Sqlmock) {
				runQueryCount := regexp.QuoteMeta(query)
				mock.ExpectExec(runQueryCount).WillReturnError(sql.ErrConnDone)
			},
			checkReturn: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:       "success",
			employeeID: 1,
			buildStub: func(mock sqlmock.Sqlmock) {

				runQuery := regexp.QuoteMeta(query)
				mock.ExpectExec(runQuery).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			checkReturn: func(err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			defer db.Close()

			tc.buildStub(mock)

			repo := NewRepoUser(db)

			err = repo.DeleteEmployee(context.TODO(), tc.employeeID)

			tc.checkReturn(err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

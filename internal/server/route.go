package server

import (
	empHandler "employee/internal/handler/employee"
	mdlwr "employee/internal/middleware"
	empRepo "employee/internal/repository/employee"
	empUsecase "employee/internal/usecase/employee"
	log "github.com/sirupsen/logrus"
)

type Router struct {
}

var rLog = log.WithField("module", "router")

func (r *Rest) ConfigureRoutes() {
	cfg := *r.Config

	r.Echo.Use(mdlwr.LoggingMiddleware)

	employeeRepo := empRepo.NewRepoUser(r.SQL)
	employeeUseCase := empUsecase.NewUseCaseEmployee(employeeRepo)
	employeeHandler := empHandler.NewEmployeeHandler(employeeUseCase, cfg)

	r.Echo.POST("/employees", employeeHandler.CreateEmployee)
	r.Echo.GET("/employees", employeeHandler.GetEmployee)
	r.Echo.GET("/employees/:employee_id", employeeHandler.GetEmployeeByID)
	r.Echo.PUT("/employees/:employee_id", employeeHandler.UpdateEmployee)
	r.Echo.DELETE("/employees/:employee_id", employeeHandler.DeleteEmployee)

}

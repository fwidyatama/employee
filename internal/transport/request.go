package transport

type CreateEmployeeReq struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required"`
	HireDate  string `json:"hire_date" validate:"required,date"`
}
type UpdateEmployeeReq struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" `
	Email     string `json:"email" validate:"required"`
	HireDate  string `json:"hire_date" validate:"required,date"`
}

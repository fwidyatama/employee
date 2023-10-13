package transport

type EmployeeRes struct {
	ID        int    `json:"id" swaggo:"example=1"`
	FirstName string `json:"first_name" swaggo:"minLength=1,example=John"`
	LastName  string `json:"last_name" swaggo:"minLength=1,example=Mayer"`
	Email     string `json:"email" swaggo:"format=email,example=johndoe@example.com"`
	HireDate  string `json:"hire_date" swaggo:"format=date,example=2023-01-15"`
}

type ListEmployees struct {
	Employees []*EmployeeRes `json:"employees"`
}

package domains

import (
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
)

// Repository
type AuthUserRepository interface {
	CheckEmailEmployeeIsExist(email string) (isExist bool, err error)
	CheckUsernameEmployeeIsExist(username string) (isExist bool, err error)
	CreateEmployee(req *models.Employee) (res *response.EmployeeProfileResponse, err error)

	CheckEmailCustomerIsExist(email string) (isExist bool, err error)
	CheckUsernameCustomerIsExist(username string) (isExist bool, err error)
	CreateCustomer(req models.Customer) (res *response.CustomerProfileResponse, err error)
}

// Usecase
type AuthUserUsecase interface {
	CreateEmployee(req *models.Employee) (res *response.EmployeeProfileResponse, err error)
	EmployeeLogin(loginData response.LoginEmployeeRequest) (empToken *string, err error)

	CreateCustomer(req models.Customer) (res *response.CustomerProfileResponse, err error)
	CustomerLogin(loginData response.LoginCustomerRequest) (cusToken *string, err error)
}

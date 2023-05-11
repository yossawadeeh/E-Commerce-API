package repository

import (
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"

	"gorm.io/gorm"
)

type authUserRepository struct {
	DB *gorm.DB
}

func NewAuthUserRepository(db *gorm.DB) domains.AuthUserRepository {
	return &authUserRepository{DB: db}
}

func (t *authUserRepository) CheckEmailEmployeeIsExist(email string) (isExist bool, err error) {
	var amountRes int64
	if err := t.DB.Model(&models.Employee{}).Where("email = ?", email).Count(&amountRes).Error; err != nil {
		return false, err
	}
	exists := amountRes > 0
	return exists, nil
}

func (t *authUserRepository) CheckUsernameEmployeeIsExist(username string) (isExist bool, err error) {
	var amountRes int64
	if err := t.DB.Model(&models.Employee{}).Where("username = ?", username).Count(&amountRes).Error; err != nil {
		return false, err
	}
	exists := amountRes > 0
	return exists, nil
}

func (t *authUserRepository) CreateEmployee(req *response.RegisterEmployeeRequest) (res *response.EmployeeProfileResponse, err error) {
	var resEmp *response.EmployeeProfileResponse

	newEmployee := models.Employee{
		Email:       req.Email,
		Username:    req.Username,
		FirstName:   req.Firstname,
		LastName:    req.Lastname,
		Password:    req.Password,
		Phone:       req.Phone,
		RoleId:      req.RoleId,
		ShopOwnerId: req.ShopOwnerId,
	}

	if err := t.DB.Create(&newEmployee).Error; err != nil {
		return nil, err
	}

	if err := t.DB.Model(&models.Employee{}).Where("id = ?", newEmployee.ID).First(&resEmp).Error; err != nil {
		return nil, err
	}
	return resEmp, nil
}

func (t *authUserRepository) CheckEmailCustomerIsExist(email string) (isExist bool, err error) {
	var amountRes int64
	if err := t.DB.Model(&models.Customer{}).Where("email = ?", email).Count(&amountRes).Error; err != nil {
		return false, err
	}
	exists := amountRes > 0
	return exists, nil
}

func (t *authUserRepository) CheckUsernameCustomerIsExist(username string) (isExist bool, err error) {
	var amountRes int64
	if err := t.DB.Model(&models.Customer{}).Where("username = ?", username).Count(&amountRes).Error; err != nil {
		return false, err
	}
	exists := amountRes > 0
	return exists, nil
}

func (t *authUserRepository) CreateCustomer(req response.RegisterCustomerRequest) (res *response.CustomerProfileResponse, err error) {
	var customer *response.CustomerProfileResponse

	newCustomer := models.Customer{
		Email:        req.Email,
		Username:     req.Username,
		FirstName:    req.Firstname,
		LastName:     req.Lastname,
		Password:     req.Password,
		Phone:        req.Phone,
		Age:          req.Age,
		Birthday:     req.Birthday,
		BirthdayText: req.BirthdayText,
	}

	if err := t.DB.Create(&newCustomer).Error; err != nil {
		return nil, err
	}

	err = t.DB.Model(&models.Customer{}).Where("id = ?", newCustomer.ID).First(&customer).Error

	return customer, err
}

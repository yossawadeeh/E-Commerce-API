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

func (t *authUserRepository) CreateEmployee(req *models.Employee) (res *response.EmployeeProfileResponse, err error) {
	var resEmp *response.EmployeeProfileResponse
	if err := t.DB.Create(&req).Error; err != nil {
		return nil, err
	}

	if err := t.DB.Model(&models.Employee{}).Where("id = ?", req.ID).First(&resEmp).Error; err != nil {
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

func (t *authUserRepository) CreateCustomer(req models.Customer) (res *response.CustomerProfileResponse, err error) {
	var customer *response.CustomerProfileResponse
	if err := t.DB.Create(&req).Error; err != nil {
		return nil, err
	}

	err = t.DB.Model(&models.Customer{}).Where("id = ?", req.ID).First(&customer).Error

	return customer, err
}

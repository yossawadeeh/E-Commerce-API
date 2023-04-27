package repository

import (
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"

	"gorm.io/gorm"
)

type shopRepository struct {
	DB *gorm.DB
}

func NewShopRepository(db *gorm.DB) domains.ShopRepository {
	return &shopRepository{DB: db}
}

func (t *shopRepository) GetEmployeeProfile(empId uint, emp *response.EmployeeProfileResponse) (err error) {
	if err := t.DB.Model(&models.Employee{}).Where("id = ?", empId).First(&emp).Error; err != nil {
		return err
	}
	return nil
}

func (t *shopRepository) GetAllRolesData() (role []models.Role, err error) {
	var res []models.Role
	if err := t.DB.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (t *shopRepository) GetEmployeeByEmail(email string) (emp *models.Employee, err error) {
	var empRes models.Employee
	if err := t.DB.Where("email = ?", email).First(&empRes).Error; err != nil {
		return nil, err
	}
	return &empRes, nil
}

func (t *shopRepository) GetEmployeeByUsername(username string) (emp *models.Employee, err error) {
	var empRes models.Employee
	if err := t.DB.Where("username = ?", username).First(&empRes).Error; err != nil {
		return nil, err
	}
	return &empRes, nil
}

func (t *shopRepository) GetShopById(shopId uint) (shop *models.ShopOwner, err error) {
	var shopRes models.ShopOwner
	if err := t.DB.Where("id = ?", shopId).First(&shopRes).Error; err != nil {
		return nil, err
	}
	return &shopRes, nil
}

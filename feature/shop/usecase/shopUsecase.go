package usecase

import (
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
)

var err error

type shopUsecase struct {
	shopRepo domains.ShopRepository
}

func NewShopUsecase(repo domains.ShopRepository) domains.ShopUsecase {
	return &shopUsecase{
		shopRepo: repo,
	}
}

func (t *shopUsecase) GetEmployeeProfile(empId uint) (emp *response.EmployeeProfileResponse, err error) {
	empRes := &response.EmployeeProfileResponse{}
	err = t.shopRepo.GetEmployeeProfile(empId, empRes)
	return empRes, err
}

func (t *shopUsecase) GetAllRolesData() (role []models.Role, err error) {
	role, err = t.shopRepo.GetAllRolesData()
	return role, err
}

func (t *shopUsecase) GetEmployeeByEmail(email string) (emp *models.Employee, err error) {
	res, err := t.shopRepo.GetEmployeeByEmail(email)
	return res, err
}

func (t *shopUsecase) GetEmployeeByUsername(username string) (emp *models.Employee, err error) {
	res, err := t.shopRepo.GetEmployeeByUsername(username)
	return res, err
}

func (t *shopUsecase) GetShopById(shopId uint) (shop *models.ShopOwner, err error) {
	res, err := t.shopRepo.GetShopById(shopId)
	return res, err
}

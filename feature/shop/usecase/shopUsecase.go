package usecase

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"errors"
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

func (t *shopUsecase) GetDailyReports(req response.DailyReportsRequest) (res *response.DailyReportsResponse, err error) {
	reports, err := t.shopRepo.GetDailyReports(req)
	return reports, err
}

func (t *shopUsecase) GetOrderReportsPeriodDate(req response.OrderReportsPeriodDateRequest) (res *response.DailyReportsResponse, err error) {
	reports, err := t.shopRepo.GetOrderReportsPeriodDate(req)
	return reports, err
}

func (t *shopUsecase) CreateShop(req *models.ShopOwner) (err error) {
	if req.Name == "" {
		return errors.New(constant.InvalidField)
	}

	var isExist bool
	if isExist, err = t.shopRepo.CheckIsExistShopName(req.Name); err != nil {
		if err.Error() != constant.RecordNotFound {
			return err
		}
	}
	if isExist == true {
		return errors.New(constant.DupicateShopName)
	}

	err = t.shopRepo.CreateShop(req)
	return err
}

func (t *shopUsecase) UpdateShop(req *models.ShopOwner) (err error) {
	if req.Name == "" || req.ID == 0 {
		return errors.New(constant.InvalidField)
	}

	var tempShop *models.ShopOwner
	if tempShop, err = t.shopRepo.GetShopById(req.ID); err != nil {
		return err
	}

	var isExist bool
	if req.Name != tempShop.Name {
		if isExist, err = t.shopRepo.CheckIsExistShopName(req.Name); err != nil {
			if err.Error() != constant.RecordNotFound {
				return err
			}
		}
		if isExist == true {
			return errors.New(constant.DupicateShopName)
		}
	}

	tempShop.Name = req.Name
	tempShop.Description = req.Description

	err = t.shopRepo.UpdateShop(tempShop)
	return err
}

func (t *shopUsecase) DeleteShop(shopId uint) (res uint, err error) {
	err = t.shopRepo.DeleteShop(shopId)
	if err != nil {
		return 0, err
	}
	return shopId, nil
}

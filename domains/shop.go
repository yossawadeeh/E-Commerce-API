package domains

import (
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
)

// Repository
type ShopRepository interface {
	GetEmployeeProfile(empId uint, emp *response.EmployeeProfileResponse) (err error)
	GetAllRolesData() (role []models.Role, err error)
	GetEmployeeByEmail(email string) (emp *models.Employee, err error)
	GetEmployeeByUsername(username string) (emp *models.Employee, err error)
	GetShopById(shopId uint) (shop *models.ShopOwner, err error)
	CheckIsExistShopName(shopName string) (res bool, err error)
	CreateShop(req *models.ShopOwner) (err error)
	UpdateShop(req *models.ShopOwner) (err error)
	DeleteShop(shopId uint) (err error)

	GetDailyReports(req response.DailyReportsRequest) (res *response.DailyReportsResponse, err error)
}

// Usecase
type ShopUsecase interface {
	GetEmployeeProfile(empId uint) (emp *response.EmployeeProfileResponse, err error)
	GetAllRolesData() (role []models.Role, err error)
	GetEmployeeByEmail(email string) (emp *models.Employee, err error)
	GetEmployeeByUsername(username string) (emp *models.Employee, err error)
	GetShopById(shopId uint) (shop *models.ShopOwner, err error)
	CreateShop(req *models.ShopOwner) (err error)
	UpdateShop(req *models.ShopOwner) (err error)
	DeleteShop(shopId uint) (res uint, err error)

	GetDailyReports(req response.DailyReportsRequest) (res *response.DailyReportsResponse, err error)
}

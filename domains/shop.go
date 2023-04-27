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
}

// Usecase
type ShopUsecase interface {
	GetEmployeeProfile(empId uint) (emp *response.EmployeeProfileResponse, err error)
	GetAllRolesData() (role []models.Role, err error)
	GetEmployeeByEmail(email string) (emp *models.Employee, err error)
	GetEmployeeByUsername(username string) (emp *models.Employee, err error)
	GetShopById(shopId uint) (shop *models.ShopOwner, err error)

	// CreateShop -> cretae employee admin
	// UpdateShop
	// DeleteShop
}

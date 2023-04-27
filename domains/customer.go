package domains

import (
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
)

// Repository
type CustomerRepository interface {
	GetCustomerProfileById(customerId uint) (customer response.CustomerProfileResponse, err error)
	GetCustomerByEmail(email string) (cus *models.Customer, err error)

	UpdateCarts(req models.Cart) (res *response.AddToCartResponse, err error)
	DeleteProductFromCarts(customerId uint, productId uint) (err error)
	CreateAddress(req *models.Address) (err error)
}

// Usecase
type CustomerUsecase interface {
	GetCustomerProfileById(customerId uint) (customer response.CustomerProfileResponse, err error)
	GetCustomerByEmail(email string) (cus *models.Customer, err error)

	UpdateCarts(req models.Cart) (res *response.AddToCartResponse, err error)
	DeleteProductFromCarts(customerId uint, productId uint) (err error)
	CreateAddress(req *models.Address) (err error)
}

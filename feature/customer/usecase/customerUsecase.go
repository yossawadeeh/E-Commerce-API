package usecase

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"errors"
)

var err error

type customerUsecase struct {
	customerRepo domains.CustomerRepository
	productRepo  domains.ProductRepository
}

func NewCustomerUsecase(customerRepo domains.CustomerRepository, productRepo domains.ProductRepository) domains.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
		productRepo:  productRepo,
	}
}

func (t *customerUsecase) GetCustomerProfileById(customerId uint) (customer response.CustomerProfileResponse, err error) {
	var customerRes response.CustomerProfileResponse
	customerRes, err = t.customerRepo.GetCustomerProfileById(customerId)
	return customerRes, err
}

func (t *customerUsecase) GetCustomerByEmail(email string) (cus *models.Customer, err error) {
	var customerRes *models.Customer
	customerRes, err = t.customerRepo.GetCustomerByEmail(email)
	return customerRes, err
}

func (t *customerUsecase) UpdateCarts(req models.Cart) (res *response.AddToCartResponse, err error) {
	var product *models.Product
	if req.ProductId <= 0 {
		return nil, errors.New(constant.InvalidField)
	}

	product, err = t.productRepo.GetProductById(req.ProductId)
	if err != nil {
		return nil, err
	}

	if req.Amount > product.Amount {
		return nil, errors.New(constant.ProductNotEnough)
	}

	var prodToCart *response.AddToCartResponse
	prodToCart, err = t.customerRepo.UpdateCarts(req)
	return prodToCart, err
}

func (t *customerUsecase) DeleteProductFromCarts(customerId uint, productId uint) (err error) {
	if _, err := t.productRepo.GetProductById(productId); err != nil {
		return errors.New(constant.ProductIdNotFound)
	}
	err = t.customerRepo.DeleteProductFromCarts(customerId, productId)
	return err
}

func (t *customerUsecase) CreateAddress(req *models.Address) (err error) {
	if req.AddressDetail == "" || req.Districts == "" || req.Province == "" || req.Country == "" || req.PostalCode == "" {
		return errors.New(constant.InvalidField)
	}

	if err := t.customerRepo.CreateAddress(req); err != nil {
		return err
	}
	return nil
}

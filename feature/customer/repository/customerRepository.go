package repository

import (
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"

	"gorm.io/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

var err error

func NewCustomerRepository(db *gorm.DB) domains.CustomerRepository {
	return &customerRepository{DB: db}
}

func (t *customerRepository) GetCustomerProfileById(customerId uint) (customer response.CustomerProfileResponse, err error) {
	var res response.CustomerProfileResponse
	err = t.DB.Model(models.Customer{}).Where("id = ?", customer).First(&res).Error
	return res, err
}

func (t *customerRepository) GetCustomerByEmail(email string) (cus *models.Customer, err error) {
	var res *models.Customer
	err = t.DB.Where("email = ?", email).First(&res).Error
	return res, err
}

func (t *customerRepository) UpdateCarts(req models.Cart) (res *response.AddToCartResponse, err error) {
	prodToCart := &response.AddToCartResponse{}

	err = t.DB.Where("customer_id = ? and product_id = ?", req.CustomerId, req.ProductId).Save(&req).Error
	if err != nil {
		err = t.DB.Create(&req).Error
		if err != nil {
			return nil, err
		}
	}

	var product models.Product
	err = t.DB.Where("id = ?", req.ProductId).First(&product).Error
	if err != nil {
		return nil, err
	}

	prodToCart.CustomerId = req.CustomerId
	prodToCart.Product = product
	prodToCart.Amount = req.Amount

	return prodToCart, err
}

func (t *customerRepository) DeleteProductFromCarts(customerId uint, productId uint) (err error) {
	err = t.DB.Where("customer_id = ? and product_id = ?", customerId, productId).Delete(&models.Cart{}).Error
	return err
}

func (t *customerRepository) CreateAddress(req *models.Address) (err error) {
	if err := t.DB.Save(req).Error; err != nil {
		return err
	}
	return nil
}

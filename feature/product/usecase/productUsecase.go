package usecase

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/models"
	"errors"
	"fmt"
)

type productUsecase struct {
	productUsecase domains.ProductRepository
}

func NewProductUsecase(repo domains.ProductRepository) domains.ProductUsecase {
	return &productUsecase{
		productUsecase: repo,
	}
}

func (t *productUsecase) GetAllProducts(shopId uint) (products *[]models.Product, err error) {
	res, err := t.productUsecase.GetAllProducts(shopId)
	return res, err
}

func (t *productUsecase) GetProductById(productId uint) (product *models.Product, err error) {
	res, err := t.productUsecase.GetProductById(productId)
	return res, err
}

func (t *productUsecase) CreateProduct(product *models.Product) (err error) {
	isExist, err := t.productUsecase.CheckProductNameIsExist(product.ShopOwnerId, product.Name)
	if isExist {
		return errors.New(constant.DupicateProductName)
	}
	return t.productUsecase.CreateProduct(product)
}

func (t *productUsecase) UpdateProduct(product *models.Product) (err error) {

	productTemp, err := t.productUsecase.GetProductById(product.ID)
	if err != nil {
		return errors.New(constant.ProductIdNotFound)
	}

	isExist, err := t.productUsecase.CheckProductNameIsExist(product.ShopOwnerId, product.Name)
	if isExist && productTemp.Name != product.Name {
		return errors.New(constant.DupicateProductName)
	}

	productTemp.Name = product.Name
	productTemp.Description = product.Description
	productTemp.Amount = product.Amount
	productTemp.Price = product.Price
	productTemp.IsActive = product.IsActive
	productTemp.ProductCategory = product.ProductCategory
	return t.productUsecase.UpdateProduct(product)
}

func (t *productUsecase) DeleteProduct(shopId uint, productId uint) (deletedId uint, err error) {
	var product *models.Product
	var id uint
	if product, err = t.productUsecase.GetProductById(productId); err != nil {
		return id, errors.New(constant.ProductIdNotFound)
	}

	fmt.Println(product)

	if id, err = t.productUsecase.DeleteProduct(product); err != nil {
		return id, err
	}
	return id, nil
}

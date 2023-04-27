package repository

import (
	"e-commerce-api/domains"
	"e-commerce-api/models"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) domains.ProductRepository {
	return &productRepository{DB: db}
}

func (t *productRepository) GetProductCategoryById(productCatId uint) (productCat *[]models.ProductCategory, err error) {
	var productCatRes *[]models.ProductCategory
	if err := t.DB.Where("id = ?", productCatId).First(&productCatRes).Error; err != nil {
		return nil, err
	}
	return productCatRes, err
}

func (t *productRepository) GetAllProducts(shopId uint) (products *[]models.Product, err error) {
	var productsRes *[]models.Product
	if err := t.DB.Where("shop_owner_id = ?", shopId).Find(&productsRes).Error; err != nil {
		return nil, err
	}
	return productsRes, err
}

func (t *productRepository) GetProductById(productId uint) (product *models.Product, err error) {
	var productRes *models.Product
	if err := t.DB.Where("id = ?", productId).First(&productRes).Error; err != nil {
		return nil, err
	}

	return productRes, err
}

func (t *productRepository) CreateProduct(product *models.Product) (err error) {
	err = t.DB.Create(product).Error
	return err
}

func (t *productRepository) CheckProductNameIsExist(shopId uint, productName string) (isExist bool, err error) {
	var amountRes int64
	if err := t.DB.Model(&models.Product{}).Where("name = ? and shop_owner_id = ?", productName, shopId).Count(&amountRes).Error; err != nil {
		return false, err
	}
	exists := amountRes > 0
	return exists, nil
}

func (t *productRepository) UpdateProduct(product *models.Product) (err error) {
	err = t.DB.Save(&product).Error
	return err
}

func (t *productRepository) DeleteProduct(product *models.Product) (deletedId uint, err error) {
	err = t.DB.Delete(&product).Error
	return product.ID, err
}

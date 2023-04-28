package domains

import "e-commerce-api/models"

type ProductRepository interface {
	GetProductCategoryById(productCatId uint) (productCat *[]models.ProductCategory, err error)

	GetAllProducts() (products *[]models.Product, err error)
	GetAllProductsByShopId(shopId uint) (products *[]models.Product, err error)
	GetProductById(productId uint) (product *models.Product, err error)
	FilterProductByCategoryId(categoryId uint) (product *[]models.Product, err error)
	FilterProductByCategoryIdAndShopId(categoryId uint, shopId uint) (product *[]models.Product, err error)

	CheckProductNameIsExist(shopId uint, productName string) (isExist bool, err error)
	CreateProduct(product *models.Product) (err error)
	UpdateProduct(product *models.Product) (err error)
	DeleteProduct(product *models.Product) (deletedId uint, err error)
}

type ProductUsecase interface {
	GetAllProducts() (products *[]models.Product, err error)
	GetAllProductsByShopId(shopId uint) (products *[]models.Product, err error)
	GetProductById(productId uint) (product *models.Product, err error)
	FilterProductByCategoryId(categoryId uint) (product *[]models.Product, err error)
	FilterProductByCategoryIdAndShopId(categoryId uint, shopId uint) (product *[]models.Product, err error)

	CreateProduct(product *models.Product) (err error)
	UpdateProduct(product *models.Product) (err error)
	DeleteProduct(shopId uint, productId uint) (deletedId uint, err error)
}

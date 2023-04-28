package repository

import (
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"time"

	"gorm.io/gorm"
)

type shopRepository struct {
	DB *gorm.DB
}

func NewShopRepository(db *gorm.DB) domains.ShopRepository {
	return &shopRepository{DB: db}
}

func (t *shopRepository) GetEmployeeProfile(empId uint, emp *response.EmployeeProfileResponse) (err error) {
	if err := t.DB.Model(&models.Employee{}).Where("id = ?", empId).First(&emp).Error; err != nil {
		return err
	}
	return nil
}

func (t *shopRepository) GetAllRolesData() (role []models.Role, err error) {
	var res []models.Role
	if err := t.DB.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (t *shopRepository) GetEmployeeByEmail(email string) (emp *models.Employee, err error) {
	var empRes models.Employee
	if err := t.DB.Where("email = ?", email).First(&empRes).Error; err != nil {
		return nil, err
	}
	return &empRes, nil
}

func (t *shopRepository) GetEmployeeByUsername(username string) (emp *models.Employee, err error) {
	var empRes models.Employee
	if err := t.DB.Where("username = ?", username).First(&empRes).Error; err != nil {
		return nil, err
	}
	return &empRes, nil
}

func (t *shopRepository) GetShopById(shopId uint) (shop *models.ShopOwner, err error) {
	var shopRes models.ShopOwner
	if err := t.DB.Where("id = ?", shopId).First(&shopRes).Error; err != nil {
		return nil, err
	}
	return &shopRes, nil
}

func (t *shopRepository) GetDailyReports(req response.DailyReportsRequest) (res *response.DailyReportsResponse, err error) {
	reportRes := &response.DailyReportsResponse{}
	var orders []models.Order
	orderIds := []uint{}
	var products []response.ProductTotalSales
	var categories []response.CategoryTotalSales

	date := req.Date.Truncate(24 * time.Hour)
	if err := t.DB.Where("DATE(order_date) = ?", date).Find(&orders).Error; err != nil {
		return nil, err
	}

	for _, order := range orders {
		orderIds = append(orderIds, order.ID)
	}

	totalSalesPrice := 0.00

	// SummaryByProductsDaily
	var summaryProducts []response.SummaryProductsDaily
	err = t.DB.Model(&models.OrderDetail{}).
		Select("order_details.product_id, products.name, products.description, products.product_category_id, sum(order_details.amount) as total_amount, sum(order_details.price) as total_price").
		Joins("JOIN products ON order_details.product_id = products.id").
		Joins("JOIN orders ON orders.id = order_details.order_id").
		Where("products.shop_owner_id = ? and orders.id IN (?)", req.ShopId, orderIds).
		Group("order_details.product_id, products.name, products.description, products.product_category_id").
		Scan(&summaryProducts).Error
	if err != nil {
		return nil, err
	}

	for _, summaryProduct := range summaryProducts {
		product := response.ProductTotalSales{
			ProductId:         summaryProduct.ProductID,
			Name:              summaryProduct.Name,
			Description:       summaryProduct.Description,
			ProductCategoryId: summaryProduct.ProductCategoryId,
			SalesPrice:        float64(summaryProduct.TotalPrice),
			SalesAmount:       float64(summaryProduct.TotalAmount),
		}
		products = append(products, product)
		totalSalesPrice += summaryProduct.TotalPrice
	}

	// SummaryByCategories
	var summaryCategories []response.CategoryTotalSales
	err = t.DB.Model(&models.OrderDetail{}).
		Select("product_categories.id as category_id, product_categories.name as category_name, sum(order_details.amount) as total_amount, sum(order_details.price) as total_price").
		Joins("JOIN orders ON orders.id = order_details.order_id").
		Joins("JOIN products ON order_details.product_id = products.id").
		Joins("JOIN product_categories ON product_categories.id = products.product_category_id").
		Where("products.shop_owner_id = ? and orders.id IN (?)", req.ShopId, orderIds).
		Group("product_categories.id, product_categories.name").
		Scan(&summaryCategories).Error
	if err != nil {
		return nil, err
	}

	for _, summaryCategory := range summaryCategories {
		category := response.CategoryTotalSales{
			CategoryId:   summaryCategory.CategoryId,
			CategoryName: summaryCategory.CategoryName,
			TotalPrice:   summaryCategory.TotalPrice,
			TotalAmount:  summaryCategory.TotalAmount,
		}
		categories = append(categories, category)
	}

	reportRes.Date = date
	reportRes.Products = products
	reportRes.Categories = categories
	reportRes.TotalSalesPrice = totalSalesPrice

	return reportRes, err
}

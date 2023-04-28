package response

import (
	"time"
)

type DailyReportsRequest struct {
	Date     time.Time `json:"date"`
	DateText string    `json:"date_text"`
	ShopId   uint      `json:"shopId"`
}

type CategoryTotalSales struct {
	CategoryId   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	TotalPrice   float64 `json:"total_price"`
	TotalAmount  float64 `json:"total_amount"`
}

type ProductTotalSales struct {
	ProductId         uint    `json:"product_id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	ProductCategoryId uint    `json:"product_category_id"`
	SalesPrice        float64 `json:"sales_price"`
	SalesAmount       float64 `json:"sales_amount"`
}

type DailyReportsResponse struct {
	Date            time.Time            `json:"date"`
	TotalSalesPrice float64              `json:"total_sales_price"`
	Categories      []CategoryTotalSales `json:"categories"`
	Products        []ProductTotalSales  `json:"products"`
}

type SummaryProductsDaily struct {
	ProductID         uint
	Name              string
	Description       string
	ProductCategoryId uint
	TotalAmount       int
	TotalPrice        float64
}

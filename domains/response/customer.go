package response

import (
	"e-commerce-api/models"
	"time"
)

type CustomerProfileResponse struct {
	ID        uint      `json:"employee_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Age       string    `json:"age"`
	Birthday  time.Time `json:"birthday"`
	Phone     string    `json:"phone"`
}

type AddToCartResponse struct {
	CustomerId uint           `json:"customer_id"`
	Product    models.Product `json:"product"`
	Amount     int64          `json:"amount"`
}

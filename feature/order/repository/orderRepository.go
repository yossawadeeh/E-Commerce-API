package repository

import (
	"e-commerce-api/domains"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domains.OrderRepository {
	return &orderRepository{DB: db}
}

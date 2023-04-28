package repository

import (
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domains.OrderRepository {
	return &orderRepository{DB: db}
}

func (t *orderRepository) CreateOrder(req *response.OrderRequest) (err error) {
	var order models.Order
	var orderDetail models.OrderDetail

	tx := t.DB.Begin()
	if err := tx.Model(&order).Create(&req.Order).Error; err != nil {
		return err
	}

	for index := range req.OrderDetails {
		req.OrderDetails[index].OrderId = req.Order.ID
	}

	err = tx.Model(&orderDetail).Create(req.OrderDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (t *orderRepository) GetOrderCustomerById(orderId uint, customerId uint) (res *models.Order, err error) {
	var order *models.Order
	err = t.DB.Preload("Customer", "id = ?", customerId).Preload("Address").Preload("OrderStatus").Preload("Shipper").Where("id = ?", orderId).First(&order).Error
	return order, err
}

func (t *orderRepository) GetOrderById(orderId uint) (res *models.Order, err error) {
	var order *models.Order
	err = t.DB.Preload("Customer").Preload("Address").Preload("OrderStatus").Preload("Shipper").Where("id = ?", orderId).First(&order).Error
	return order, err
}

func (t *orderRepository) GetAllCustomerOrders(customerId uint) (res *[]models.Order, err error) {
	var order *[]models.Order

	if err = t.DB.Where("customer_id = ?", customerId).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (t *orderRepository) GetPaymentByOrderId(orderId uint) (res *models.Payment, err error) {
	var payment *models.Payment
	err = t.DB.Where("order_id = ?", orderId).First(&payment).Error
	return payment, err
}

func (t *orderRepository) CreatePayment(req *models.Payment) (err error) {
	err = t.DB.Create(req).Error
	return err
}

func (t *orderRepository) UpdateOrder(req *models.Order) (err error) {
	err = t.DB.Model(models.Order{}).Where("id = ?", req.ID).Save(req).Error
	return err
}

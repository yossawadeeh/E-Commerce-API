package domains

import (
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
)

type OrderRepository interface {
	CreateOrder(req *response.OrderRequest) (err error)
	GetOrderCustomerById(orderId uint, customerId uint) (res *models.Order, err error)
	GetOrderById(orderId uint) (res *models.Order, err error)
	GetAllCustomerOrders(customerId uint) (res *[]models.Order, err error)
	UpdateOrder(req *models.Order) (err error)

	GetPaymentByOrderId(orderId uint) (res *models.Payment, err error)
	CreatePayment(req *models.Payment) (err error)
}

type OrderUsecase interface {
	CreateOrder(req *response.OrderRequest) (err error)
	GetOrderCustomerById(orderId uint, customerId uint) (res *models.Order, err error)
	GetOrderById(orderId uint) (res *models.Order, err error)
	GetAllCustomerOrders(customerId uint) (res *[]models.Order, err error)
	UpdateOrder(req response.UpdateOrderRequest) (res *models.Order, err error)

	CreatePayment(req *models.Payment, customerId uint) (err error)
}

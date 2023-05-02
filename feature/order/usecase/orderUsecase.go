package usecase

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"errors"
	"time"
)

var err error

type orderUsecase struct {
	orderRepo   domains.OrderRepository
	productRepo domains.ProductRepository
}

func NewCustomerUsecase(orderRepo domains.OrderRepository, productRepo domains.ProductRepository) domains.OrderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (t *orderUsecase) CreateOrder(req *response.OrderRequest) (err error) {
	var totalPrice float64
	for index, orderDetail := range req.OrderDetails {
		if orderDetail.ProductId == 0 {
			return errors.New(constant.InvalidField)
		}
		var productRemain *models.Product
		if productRemain, err = t.productRepo.GetProductById(orderDetail.ProductId); err != nil {
			return errors.New(constant.InvalidField)
		}
		if orderDetail.Amount > productRemain.Amount {
			return errors.New(constant.ProductNotEnough)
		}
		req.OrderDetails[index].Price = float64(orderDetail.Amount) * productRemain.Price
		totalPrice += req.OrderDetails[index].Price
	}

	req.Order.OrderStatusId = 1 // Waiting for payment
	req.Order.OrderDate = time.Now()
	req.Order.TotalPrice = totalPrice

	if err = t.orderRepo.CreateOrder(req); err != nil {
		return err
	}

	// update product remain
	for _, val := range req.OrderDetails {
		product, _ := t.productRepo.GetProductById(val.ProductId)
		product.Amount = product.Amount - val.Amount
		err = t.productRepo.UpdateProduct(product)
	}

	return err
}

func (t *orderUsecase) GetOrderCustomerById(orderId uint, customerId uint) (res *models.Order, err error) {
	var order *models.Order
	order, err = t.orderRepo.GetOrderCustomerById(orderId, customerId)
	return order, err
}

func (t *orderUsecase) GetOrderCustomerByIdResponse(orderId uint, customerId uint) (res *response.OrderResponse, err error) {
	var response *response.OrderResponse
	response, err = t.orderRepo.GetOrderCustomerByIdResponse(orderId, customerId)
	return response, err
}

func (t *orderUsecase) GetOrderById(orderId uint) (res *models.Order, err error) {
	var order *models.Order
	order, err = t.orderRepo.GetOrderById(orderId)
	return order, err
}

func (t *orderUsecase) GetOrderByIdResponse(orderId uint) (res *response.OrderResponse, err error) {
	var response *response.OrderResponse
	response, err = t.orderRepo.GetOrderByIdResponse(orderId)
	return response, err
}

func (t *orderUsecase) GetAllCustomerOrders(customerId uint) (res *[]models.Order, err error) {
	var orders *[]models.Order
	orders, err = t.orderRepo.GetAllCustomerOrders(customerId)
	return orders, err
}

func (t *orderUsecase) CreatePayment(req *models.Payment, customerId uint) (err error) {
	var order *models.Order
	if order, err = t.orderRepo.GetOrderCustomerById(req.OrderId, customerId); err != nil {
		return err
	}

	if _, err = t.orderRepo.GetPaymentByOrderId(req.OrderId); err != nil {
		if err.Error() != constant.RecordNotFound {
			return err
		}
	} else {
		return errors.New(constant.PaymentPaid)
	}

	if req.PaymentAmount < order.TotalPrice {
		return errors.New(constant.PaymentInvalid)
	}

	req.PaymentDate = time.Now()
	if err = t.orderRepo.CreatePayment(req); err != nil {
		return err
	}

	// Update order status
	order.OrderStatusId = 2 // paid
	if err = t.orderRepo.UpdateOrder(order); err != nil {
		return err
	}

	return nil
}

func (t *orderUsecase) UpdateOrder(req response.UpdateOrderRequest) (res *models.Order, err error) {
	var order *models.Order
	if order, err = t.orderRepo.GetOrderById(req.OrderId); err != nil {
		return nil, err
	}

	order.ShipperId = req.ShipperId
	order.OrderStatusId = req.OrderStatusId
	order.AddressId = req.AddressId
	order.TrackNo = req.TrackNo

	err = t.orderRepo.UpdateOrder(order)
	return order, err
}

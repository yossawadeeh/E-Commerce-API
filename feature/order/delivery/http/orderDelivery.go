package delivery

import "e-commerce-api/domains"

var err error

type OrderHandler struct {
	orderUsecase domains.OrderUsecase
}

func NewCustomerHandler(usecase domains.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: usecase,
	}
}

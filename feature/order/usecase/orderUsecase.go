package usecase

import "e-commerce-api/domains"

var err error

type orderUsecase struct {
	orderRepo domains.OrderRepository
}

func NewCustomerUsecase(orderRepo domains.OrderRepository) domains.OrderUsecase {
	return &orderUsecase{
		orderRepo: orderRepo,
	}
}

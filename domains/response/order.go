package response

import (
	"e-commerce-api/models"
)

type OrderRequest struct {
	OrderDetails []models.OrderDetail `json:"order_details"`
	Order        models.Order         `json:"order"`
}

type OrderResponse struct {
	Order        models.Order         `json:"order"`
	OrderDetails []models.OrderDetail `json:"order_details"`
}

type UpdateOrderRequest struct {
	OrderId       uint    `json:"order_id"`
	ShipperId     uint    `json:"shipper_id"`
	OrderStatusId uint    `json:"order_status_id"`
	AddressId     uint    `json:"address_id"`
	TrackNo       *string `json:"track_no"`
}

type ProductIdRequest struct {
	ProductIds []uint `json:"products_id"`
}

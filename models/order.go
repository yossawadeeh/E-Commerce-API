package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	StatusName  string         `gorm:"not null" json:"status_name"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type Order struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderDate time.Time      `gorm:"not null" json:"order_date"`
	TrackNo   *string        `json:"track_no"`
	CreatedAt time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`

	CustomerId uint      `json:"customer_id"`
	Customer   *Customer `gorm:"foreignKey:CustomerId; not null" json:"customer"`

	ShipperId uint      `json:"shipper_id"`
	Shipper   *Customer `gorm:"foreignKey:ShipperId; not null" json:"shipper"`

	OrderStatusId uint         `json:"order_status_id"`
	OrderStatus   *OrderStatus `gorm:"foreignKey:OrderStatusId; not null" json:"order_status"`

	AddressId uint     `json:"address_id"`
	Address   *Address `gorm:"foreignKey:AddressId; not null" json:"address"`
}

type OrderDetail struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Amount    int64          `gorm:"not null" json:"amount"`
	Price     float64        `gorm:"not null" json:"price"`
	CreatedAt time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`

	ProductId uint     `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductId; not null" json:"product"`

	OrderId uint   `json:"order_id"`
	Order   *Order `gorm:"foreignKey:OrderId; not null" json:"order"`
}

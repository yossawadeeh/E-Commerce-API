package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"not null" json:"username"`
	Email        string         `gorm:"not null" json:"email"`
	Password     string         `gorm:"not null"`
	FirstName    string         `json:"firstname"`
	LastName     string         `json:"lastname"`
	Age          *int64         `json:"age"`
	Birthday     time.Time      `json:"birthday"`
	BirthdayText string         `json:"birthday_text"`
	Phone        string         `json:"phone"`
	CreatedAt    time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

type Address struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	AddressDetail string         `gorm:"not null" json:"address_detail"`
	Districts     string         `gorm:"not null" json:"districts"`
	Province      string         `gorm:"not null" json:"province"`
	Country       string         `gorm:"not null" json:"country"`
	PostalCode    string         `gorm:"not null" json:"postal_code"`
	CreatedAt     time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"-"`

	CustomerId uint      `json:"customer_id"`
	Customer   *Customer `gorm:"foreignKey:CustomerId; not null" json:"customer"`
}

type Cart struct {
	Amount int64 `gorm:"not null" json:"amount"`

	CustomerId uint      `json:"customer_id"`
	Customer   *Customer `gorm:"foreignKey:CustomerId; not null" json:"customer"`

	ProductId uint     `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductId; not null" json:"product"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentType struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TypeName  string         `gorm:"not null" json:"type_name"`
	CreatedAt time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Payment struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	PaymentAmount float64        `gorm:"not null" json:"payment_amount"`
	PaymentDate   time.Time      `gorm:"not null; default:now()" json:"payment_date"`
	CreatedAt     time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

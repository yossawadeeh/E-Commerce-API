package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipper struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CompanyName string         `gorm:"not null" json:"company_name"`
	Phone       *string        `json:"phone"`
	CreatedAt   time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

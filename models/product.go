package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description *string        `json:"description"`
	Amount      int64          `gorm:"default:0" json:"amount"`
	Price       float64        `gorm:"default:0.00" json:"price"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-"`

	ShopOwnerId uint       `json:"shop_owner_id"`
	ShopOwner   *ShopOwner `gorm:"foreignKey:ShopOwnerId; not null" json:"shop_owner"`

	ProductCategoryId uint             `json:"product_category_id"`
	ProductCategory   *ProductCategory `gorm:"foreignKey:ProductCategoryId; not null" json:"product_category"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	PermissionName string         `gorm:"not null" json:"permission_name"`
	Detail         *string        `json:"detail"`
	CreatedAt      time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}

type Role struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	RoleName  string         `gorm:"not null" json:"role_name"`
	CreatedAt time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type RolePermission struct {
	PermissionId uint        `json:"permission_id"`
	Permission   *Permission `gorm:"foreignKey:PermissionId; not null" json:"permission"`

	RoleId uint  `json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleId; not null" json:"role"`
}

type Employee struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"not null" json:"username"`
	Email     string         `gorm:"not null" json:"email"`
	Password  string         `gorm:"not null"`
	FirstName string         `json:"firstname"`
	LastName  string         `json:"lastname"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`

	ShopOwnerId uint       `json:"shop_id"`
	ShopOwner   *ShopOwner `gorm:"foreignKey:ShopOwnerId; not null" json:"shop_owner"`

	RoleId uint  `gorm:"default:1" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleId; not null" json:"role"`
}

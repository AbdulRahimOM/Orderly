package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Credentials struct {
	ID             uuid.UUID `json:"id" gorm:"column:id"`
	Username       string    `json:"username" gorm:"column:username"`
	HashedPassword string    `json:"-" gorm:"column:hashed_password"`
}

type AdminInList struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	Designation string    `gorm:"column:designation" json:"designation"`
	IsActive    bool      `gorm:"column:is_active" json:"isActive"`
}

type Admin struct {
	ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
	Email       string         `gorm:"column:email;unique" json:"email"`
	Name        string         `gorm:"column:name" json:"name"`
	Phone       string         `gorm:"column:phone" json:"phone"`
	Designation string         `gorm:"column:designation" json:"designation"`
	IsActive    bool           `gorm:"column:is_active" json:"isActive"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	IsDeleted   bool           `gorm:"is_deleted" json:"isDeleted"`
}

type UserSignInDetails struct {
	Name      string `gorm:"column:name" json:"name"`
	IsBlocked bool   `gorm:"column:is_blocked" json:"isBlocked"`
	Phone     string `gorm:"column:phone" json:"phone"`
}

type User struct {
	ID        uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Email     string         `gorm:"column:email;unique" json:"email"`
	Phone     string         `gorm:"column:phone" json:"phone"`
	IsBlocked bool           `gorm:"column:is_blocked;default:false" json:"isBlocked"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	IsDeleted bool           `gorm:"is_deleted" json:"isDeleted"`
}

type UserInList struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	IsBlocked bool      `gorm:"column:is_blocked;default:false" json:"isBlocked"`
	IsDeleted bool      `gorm:"is_deleted" json:"isDeleted"`
}

type ProductInList struct {
	ID               int     `gorm:"column:id;primaryKey" json:"id"`
	Name             string  `gorm:"column:name" json:"name"`
	Description      string  `gorm:"column:description" json:"description"`
	CurrentSalePrice float64 `gorm:"column:current_sale_price" json:"currentSalePrice"`
	MaxSalePrice     float64 `gorm:"column:max_sale_price" json:"maxSalePrice"`
	CurrentStock     int     `gorm:"column:current_stock" json:"currentStock"`
}

type Product struct {
	ID               int            `gorm:"column:id;primaryKey" json:"id"`
	Name             string         `gorm:"column:name" json:"name"`
	Description      string         `gorm:"column:description" json:"description"`
	CategoryID       int            `gorm:"column:category_id" json:"categoryId"`
	CategoryName     string         `gorm:"column:category_name" json:"categoryName"`
	MinSalePrice     float64        `gorm:"column:min_sale_price" json:"minSalePrice"`
	MaxSalePrice     float64        `gorm:"column:max_sale_price" json:"maxSalePrice"`
	BasePrice        float64        `gorm:"column:base_price" json:"basePrice"`
	CurrentSalePrice float64        `gorm:"column:current_sale_price" json:"currentSalePrice"`
	OptimalStock     int            `gorm:"column:optimal_stock" json:"optimalStock"`
	CurrentStock     int            `gorm:"column:current_stock" json:"currentStock"`
	CreatedAt        time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	IsDeleted        bool           `gorm:"column:is_deleted" json:"isDeleted"`
}

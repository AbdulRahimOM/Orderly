package dto

import (
	"time"

	"gorm.io/gorm"
)

type Credentials struct {
	ID             int    `json:"id" gorm:"column:id"`
	Username       string `json:"username" gorm:"column:username"`
	HashedPassword string `json:"-" gorm:"column:hashed_password"`
}

type AdminInList struct {
	ID          int    `gorm:"column:id;primaryKey" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Designation string `gorm:"column:designation" json:"designation"`
	IsBlocked   bool   `gorm:"column:is_blocked;default:false" json:"isBlocked"`
}

type Admin struct {
	ID          int            `gorm:"column:id;primaryKey" json:"id"`
	Email       string         `gorm:"column:email;unique" json:"email"`
	Name        string         `gorm:"column:name" json:"name"`
	Phone       string         `gorm:"column:phone" json:"phone"`
	Designation string         `gorm:"column:designation" json:"designation"`
	IsBlocked   bool           `gorm:"column:is_blocked;default:false" json:"isBlocked"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	IsDeleted   bool           `gorm:"is_deleted" json:"isDeleted"`
}

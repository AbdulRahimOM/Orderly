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
	IsActive	bool      `gorm:"column:is_active" json:"isActive"`
}

type Admin struct {
	ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
	Email       string         `gorm:"column:email;unique" json:"email"`
	Name        string         `gorm:"column:name" json:"name"`
	Phone       string         `gorm:"column:phone" json:"phone"`
	Designation string         `gorm:"column:designation" json:"designation"`
	IsActive	bool           `gorm:"column:is_active" json:"isActive"`
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

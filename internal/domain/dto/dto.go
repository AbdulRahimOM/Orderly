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

type UserProfile struct {
	Name      string `gorm:"column:name" json:"name"`
	Email     string `gorm:"column:email" json:"email"`
	Phone     string `gorm:"column:phone" json:"phone"`
	IsBlocked bool   `gorm:"column:is_blocked" json:"isBlocked"`
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

type UserAddress struct {
	ID       uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	House    string    `gorm:"column:house" json:"house"`
	Street1  string    `gorm:"column:street1" json:"street1"`
	Street2  string    `gorm:"column:street2" json:"street2"`
	City     string    `gorm:"column:city" json:"city"`
	State    string    `gorm:"column:state" json:"state"`
	Pincode  string    `gorm:"column:pincode" json:"pincode"`
	Landmark string    `gorm:"column:landmark" json:"landmark"`
	Country  string    `gorm:"column:country" json:"country"`
}

type Cart struct {
	ProductID          int     `gorm:"column:product_id;primaryKey" json:"productId"`
	Quantity           int     `gorm:"column:quantity" json:"quantity"`
	PriceWhenPutInCart float64 `gorm:"column:price_when_put_in_cart" json:"priceWhenPutInCart"`
	CurrentSalePrice   float64 `gorm:"column:current_sale_price" json:"currentSalePrice"`
}

type OrderInListForAdmin struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	UserID        uuid.UUID `gorm:"column:user_id" json:"userId"`
	CustomerName  string    `gorm:"column:customer_name" json:"customerName"`
	OrderTime     time.Time `gorm:"column:order_time" json:"orderTime"`
	TotalAmount   float64   `gorm:"column:total_amount" json:"totalAmount"`
	PaymentMethod string    `gorm:"column:payment_method" json:"paymentMethod"`
	OrderStatus   string    `gorm:"column:order_status" json:"orderStatus"`
}

type OrderInListForUser struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	OrderTime     time.Time `gorm:"column:order_time" json:"orderTime"`
	TotalAmount   float64   `gorm:"column:total_amount" json:"totalAmount"`
	PaymentMethod string    `gorm:"column:payment_method" json:"paymentMethod"`
	OrderStatus   string    `gorm:"column:order_status" json:"orderStatus"`
}

type OrderForAdmin struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	UserID        uuid.UUID `gorm:"column:user_id" json:"userId"`
	CustomerName  string    `gorm:"column:customer_name" json:"customerName"`
	OrderTime     time.Time `gorm:"column:order_time" json:"orderTime"`
	TotalAmount   float64   `gorm:"column:total_amount" json:"totalAmount"`
	PaymentMethod string    `gorm:"column:payment_method" json:"paymentMethod"`
	PaymentID     string    `gorm:"column:payment_id" json:"paymentId"`
	OrderStatus   string    `gorm:"column:order_status" json:"orderStatus"`
	CancelledAt   *time.Time `gorm:"column:cancelled_at" json:"cancelledAt,omitempty"`
	DeliveredAt   *time.Time `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`

	//Shipping Address
	House    string `gorm:"column:house" json:"house"`
	Street1  string `gorm:"column:street1" json:"street1"`
	Street2  string `gorm:"column:street2" json:"street2"`
	City     string `gorm:"column:city" json:"city"`
	State    string `gorm:"column:state" json:"state"`
	Pincode  string `gorm:"column:pincode" json:"pincode"`
	Landmark string `gorm:"column:landmark" json:"landmark"`
	Country  string `gorm:"column:country" json:"country"`
}

type OrderForUser struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	OrderTime     time.Time `gorm:"column:order_time" json:"orderTime"`
	TotalAmount   float64   `gorm:"column:total_amount" json:"totalAmount"`
	PaymentMethod string    `gorm:"column:payment_method" json:"paymentMethod"`
	PaymentID     string    `gorm:"column:payment_id" json:"paymentId"`
	OrderStatus   string    `gorm:"column:order_status" json:"orderStatus"`
	CancelledAt   *time.Time `gorm:"column:cancelled_at" json:"cancelledAt,omitempty"`
	DeliveredAt   *time.Time `gorm:"column:delivered_at" json:"deliveredAt,omitempty"`

	//Shipping Address
	House    string `gorm:"column:house" json:"house"`
	Street1  string `gorm:"column:street1" json:"street1"`
	Street2  string `gorm:"column:street2" json:"street2"`
	City     string `gorm:"column:city" json:"city"`
	State    string `gorm:"column:state" json:"state"`
	Pincode  string `gorm:"column:pincode" json:"pincode"`
	Landmark string `gorm:"column:landmark" json:"landmark"`
	Country  string `gorm:"column:country" json:"country"`
}

type CartItemsForOrder struct {
	ProductID        int     `gorm:"column:product_id;primaryKey" json:"productId"`
	Quantity         int     `gorm:"column:quantity" json:"quantity"`
	CurrentSalePrice float64 `gorm:"column:current_sale_price" json:"currentSalePrice"`
	CurrentStock     int     `gorm:"column:current_stock" json:"currentStock"`
}

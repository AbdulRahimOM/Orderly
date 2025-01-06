package models

import (
	"fmt"
	"net/http"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const ()

// Table names as constants
const (
	Addresses_TableName            = "addresses"
	AdminPrivileges_TableName      = "admin_privileges"
	Admins_TableName               = "admins"
	CartItems_TableName            = "cart_items"
	Category_TableName             = "categories"
	IncomingTransactions_TableName = "incoming_transactions"
	Orders_TableName               = "orders"
	OrderProducts_TableName        = "order_products"
	ProductRating_TableName        = "product_ratings"
	ProductReviews_TableName       = "product_reviews"
	Products_TableName             = "products"
	Returns_TableName              = "returns"
	RefundTransactions_TableName   = "refund_transactions"
	SignupUsers_TableName          = "signup_users"
	SuperAdmin_TableName           = "super_admin"
	Users_TableName                = "users"
)

// Users table
type User struct {
	ID             uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
	Username       string         `gorm:"column:username;unique" json:"username"`
	HashedPassword string         `gorm:"column:hashed_password" json:"-"`
	Email          string         `gorm:"column:email" json:"email"`
	Name           string         `gorm:"column:name" json:"name"`
	Phone          string         `gorm:"column:phone" json:"phone"`
	CreatedAt      time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	IsActive       bool           `gorm:"column:is_active;default:true" json:"isActive"`
}

func (User) TableName() string {
	return Users_TableName
}

func (u User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New() // Generate a new UUID
	}
	return
}

func (User) PostTableCreation(db *gorm.DB) error {
	err := db.Raw(`
		CREATE UNIQUE INDEX IF NOT EXISTS uni_users_email 
		ON users (email)
		WHERE deleted_at IS NULL;
	`).Error
	if err != nil {
		return fmt.Errorf("error creating unique index on users table in email column: %v", err)
	}

	err = db.Raw(`
		CREATE UNIQUE INDEX IF NOT EXISTS uni_users_phone
		ON users (phone)
		WHERE deleted_at IS NULL;
	`).Error
	if err != nil {
		return fmt.Errorf("error creating unique index on users table in phone column: %v", err)
	}

	return nil
}

// Addresses table
type Address struct {
	ID       uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"column:user_id" json:"userId"`
	House    string    `gorm:"column:house" json:"house"`
	Street1  string    `gorm:"column:street1" json:"street1"`
	Street2  string    `gorm:"column:street2" json:"street2"`
	City     string    `gorm:"column:city" json:"city"`
	State    string    `gorm:"column:state" json:"state"`
	Pincode  string    `gorm:"column:pincode" json:"pincode"`
	Landmark string    `gorm:"column:landmark" json:"landmark"`
	Country  string    `gorm:"column:country" json:"country"`
}

func (Address) TableName() string {
	return Addresses_TableName
}

func (m Address) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New() // Generate a new UUID
	}
	return
}

type SuperAdmin struct {
	ID             uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	Username       string    `gorm:"column:username;unique" json:"username"`
	Email          string    `gorm:"column:email;unique" json:"email"`
	HashedPassword string    `gorm:"column:hashed_password" json:"-"`
}

func (SuperAdmin) TableName() string {
	return SuperAdmin_TableName
}

func (m SuperAdmin) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New() // Generate a new UUID
	}
	return
}

// Admins table
type Admin struct {
	ID             uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
	Username       string         `gorm:"column:username;unique" json:"username"`
	HashedPassword string         `gorm:"column:hashed_password" json:"-"`
	Email          string         `gorm:"column:email;unique" json:"email"`
	Name           string         `gorm:"column:name" json:"name"`
	Phone          string         `gorm:"column:phone" json:"phone"`
	Designation    string         `gorm:"column:designation" json:"designation"`
	IsActive       bool           `gorm:"column:is_active;default:true" json:"isActive"`
	CreatedAt      time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (Admin) TableName() string {
	return Admins_TableName
}

func (m Admin) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New() // Generate a new UUID
	}
	return
}

func (Admin) GetResponseFromDBError(err error) *response.Response {
	switch err.Error() {
	case "ERROR: duplicate key value violates unique constraint \"uni_admins_username\" (SQLSTATE 23505)":
		return response.ErrorResponse(http.StatusConflict, respcode.USERNAME_ALREADY_EXISTS, err)
	case "ERROR: duplicate key value violates unique constraint \"uni_admins_email\" (SQLSTATE 23505)":
		return response.ErrorResponse(http.StatusConflict, respcode.EMAIL_ALREADY_EXISTS, err)
	default:
		return response.DBErrorResponse(err)
	}
}

// AdminPrivileges table
type AdminPrivilege struct {
	AdminID    uuid.UUID `gorm:"column:admin_id;primaryKey" json:"adminId"`
	AccessRole string    `gorm:"column:access_role;primaryKey" json:"accessRole"`

	Admin Admin `gorm:"foreignKey:AdminID;references:ID" json:"-"`
}

func (AdminPrivilege) TableName() string {
	return AdminPrivileges_TableName
}

// Products table
type Product struct {
	ID               int            `gorm:"column:id;primaryKey" json:"id"`
	Name             string         `gorm:"column:name" json:"name"`
	Description      string         `gorm:"column:description" json:"description"`
	CategoryID       int            `gorm:"column:category_id" json:"CategoryId"`
	MinSalePrice     float64        `gorm:"column:min_sale_price" json:"minSalePrice"`
	MaxSalePrice     float64        `gorm:"column:max_sale_price" json:"maxSalePrice"`
	BasePrice        float64        `gorm:"column:base_price" json:"basePrice"` //foundational price that serves as a reference for adjustments based on market conditions.
	CurrentSalePrice float64        `gorm:"column:current_sale_price" json:"currentSalePrice"`
	OptimalStock     int            `gorm:"column:optimal_stock" json:"optimalStock"` //the ideal amount of stock to have on hand.
	CurrentStock     int            `gorm:"column:current_stock" json:"currentStock"`
	CreatedAt        time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Category Category `gorm:"foreignKey:CategoryID;references:ID" json:"-"`
}

func (Product) TableName() string {
	return Products_TableName
}

func (m Product) PostTableCreation(db *gorm.DB) error {
	err := db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS uni_products_name 
		ON products (name)
		WHERE deleted_at IS NULL;
	`).Error
	if err != nil {
		return fmt.Errorf("error creating unique index on products table in name column: %v", err)
	}

	return nil
}

// Orders table
type Order struct {
	ID                     uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	UserID                 uuid.UUID `gorm:"column:user_id" json:"userId"`
	OrderDate              string    `gorm:"column:order_date;type:date" json:"orderDate"`
	TotalAmount            float64   `gorm:"column:total_amount" json:"totalAmount"`
	PaymentMethod          string    `gorm:"column:payment_method" json:"paymentMethod"`
	PaymentID              int       `gorm:"column:payment_id" json:"paymentId"`
	OrderStatus            string    `gorm:"column:order_status" json:"orderStatus"`
	DeliveredOrCancelledAt time.Time `gorm:"column:delivered_or_cancelled_at;type:date" json:"deliveredOrCancelledAt"`

	User User `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

func (Order) TableName() string {
	return Orders_TableName
}

func (m Order) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New() // Generate a new UUID
	}
	return
}

// OrderProducts table
type OrderProduct struct {
	OrderID          uuid.UUID `gorm:"column:order_id;primaryKey" json:"orderId"`
	ProductID        int       `gorm:"column:product_id;primaryKey" json:"productId"`
	Quantity         int       `gorm:"column:quantity" json:"quantity"`
	PerUnitSalePrice int       `gorm:"column:per_unit_sale_price" json:"perUnitSalePrice"`

	Order   Order   `gorm:"foreignKey:OrderID;references:ID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}

func (OrderProduct) TableName() string {
	return OrderProducts_TableName
}

// Returns table
type Return struct {
	ID                      uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	ProductID               int       `gorm:"column:product_id" json:"productId"`
	OrderID                 uuid.UUID `gorm:"column:order_id" json:"orderId"`
	ReturnRequestDate       time.Time `gorm:"column:return_request_date;type:date" json:"returnRequestDate"`
	ReturnCollectedDate     time.Time `gorm:"column:return_collected_date;type:date" json:"returnCollectedDate"`
	ReturnItemReachBackDate time.Time `gorm:"column:return_item_reach_back_date;type:date" json:"returnItemReachBackDate"`
	RefundTransactionID     uuid.UUID `gorm:"column:refund_transaction_id" json:"refundTransactionId"`

	Order             Order             `gorm:"foreignKey:OrderID;references:ID" json:"-"`
	Product           Product           `gorm:"foreignKey:ProductID;references:ID" json:"-"`
	RefundTransaction RefundTransaction `gorm:"foreignKey:RefundTransactionID;references:ID" json:"-"`
}

func (Return) TableName() string {
	return Returns_TableName
}

func (m Return) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New() // Generate a new UUID
	}
	return
}

// IncomingTransactions table
type IncomingTransaction struct {
	ID                uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	PaymentAmount     float64   `gorm:"column:payment_amount" json:"paymentAmount"`
	TransactionTime   time.Time `gorm:"column:transaction_time" json:"transactionTime"`
	TransactionMethod string    `gorm:"column:transaction_method" json:"transactionMethod"`
	Status            string    `gorm:"column:status" json:"status"`
}

func (IncomingTransaction) TableName() string {
	return IncomingTransactions_TableName
}

func (m IncomingTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New() // Generate a new UUID
	}
	return
}

// RefundTransactions table
type RefundTransaction struct {
	ID                uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	PaymentAmount     float64   `gorm:"column:payment_amount" json:"paymentAmount"`
	TransactionTime   time.Time `gorm:"column:transaction_time" json:"transactionTime"`
	TransactionMethod string    `gorm:"column:transaction_method" json:"transactionMethod"`
	Status            string    `gorm:"column:status" json:"status"`
}

func (RefundTransaction) TableName() string {
	return RefundTransactions_TableName
}

func (m RefundTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New() // Generate a new UUID
	}
	return
}

// CartItems table
type CartItem struct {
	UserID             uuid.UUID `gorm:"column:user_id;primaryKey" json:"userId"`
	ProductID          int       `gorm:"column:product_id;primaryKey" json:"productId"`
	Quantity           int       `gorm:"column:quantity" json:"quantity"`
	PriceWhenPutInCart float64   `gorm:"column:price_when_put_in_cart" json:"priceWhenPutInCart"`

	User    User    `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}

func (CartItem) TableName() string {
	return CartItems_TableName
}

type ProductRating struct {
	UserID    uuid.UUID `gorm:"column:user_id;primaryKey" json:"userId"`
	ProductID int       `gorm:"column:product_id;primaryKey" json:"productId"`
	Rating    float64   `gorm:"column:rating" json:"rating"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`

	User    User    `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}

func (ProductRating) TableName() string {
	return ProductRating_TableName
}

type ProductReview struct {
	UserID    uuid.UUID      `gorm:"column:user_id;primaryKey" json:"userId"`
	ProductID int            `gorm:"column:product_id;primaryKey" json:"productId"`
	Review    string         `gorm:"column:review" json:"review"`
	PicLinks  pq.StringArray `gorm:"column:pic_links;type:varchar(100)[]" json:"picLinks"`

	User    User    `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}

func (ProductReview) TableName() string {
	return ProductReviews_TableName
}

type Category struct {
	ID        int            `gorm:"column:id;primaryKey" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Category) TableName() string {
	return Category_TableName
}

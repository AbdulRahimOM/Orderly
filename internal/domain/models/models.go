package models

import (
	"net/http"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	"time"

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
	IncomingTransactions_TableName = "incoming_transactions"
	Orders_TableName               = "orders"
	OrderProducts_TableName        = "order_products"
	ProductCategory_TableName      = "product_categories"
	ProductRating_TableName        = "product_ratings"
	ProductReviews_TableName       = "product_reviews"
	Products_TableName             = "products"
	Returns_TableName              = "returns"
	RefundTransactions_TableName   = "refund_transactions"
	SuperAdmin_TableName           = "super_admin"
	Users_TableName                = "users"
)

// Users table
type User struct {
	ID             int            `gorm:"column:id;primaryKey" json:"id"`
	Username       string         `gorm:"column:username;unique" json:"username"`
	HashedPassword string         `gorm:"column:hashed_password" json:"-"`
	Email          string         `gorm:"column:email;unique" json:"email"`
	Name           string         `gorm:"column:name" json:"name"`
	Phone          string         `gorm:"column:phone" json:"phone"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	IsBlocked      bool           `gorm:"column:is_blocked;default:false" json:"isBlocked"`
}

func (User) TableName() string {
	return Users_TableName
}

// Addresses table
type Address struct {
	ID       int    `gorm:"column:id;primaryKey" json:"id"`
	UserID   int    `gorm:"column:user_id" json:"userId"`
	House    string `gorm:"column:house" json:"house"`
	Street1  string `gorm:"column:street1" json:"street1"`
	Street2  string `gorm:"column:street2" json:"street2"`
	City     string `gorm:"column:city" json:"city"`
	State    string `gorm:"column:state" json:"state"`
	Pincode  string `gorm:"column:pincode" json:"pincode"`
	Landmark string `gorm:"column:landmark" json:"landmark"`
	Country  string `gorm:"column:country" json:"country"`
}

func (Address) TableName() string {
	return Addresses_TableName
}

type SuperAdmin struct {
	ID             int    `gorm:"column:id;primaryKey" json:"id"`
	Username       string `gorm:"column:username;unique" json:"username"`
	Email          string `gorm:"column:email;unique" json:"email"`
	HashedPassword string `gorm:"column:hashed_password" json:"-"`
}

func (SuperAdmin) TableName() string {
	return SuperAdmin_TableName
}

// Admins table
type Admin struct {
	ID             int            `gorm:"column:id;primaryKey" json:"id"`
	Username       string         `gorm:"column:username;unique" json:"username"`
	HashedPassword string         `gorm:"column:hashed_password" json:"-"`
	Email          string         `gorm:"column:email;unique" json:"email"`
	Name           string         `gorm:"column:name" json:"name"`
	Phone          string         `gorm:"column:phone" json:"phone"`
	Designation    string         `gorm:"column:designation" json:"designation"`
	IsBlocked      bool           `gorm:"column:is_blocked;default:false" json:"isBlocked"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (Admin) TableName() string {
	return Admins_TableName
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
	AdminID    int    `gorm:"column:admin_id;primaryKey" json:"adminId"`
	AccessRole string `gorm:"column:access_role;primaryKey" json:"accessRole"`

	Admin Admin `gorm:"foreignKey:AdminID;references:ID" json:"-"`
}

func (AdminPrivilege) TableName() string {
	return AdminPrivileges_TableName
}

// Products table
type Product struct {
	ID                int            `gorm:"column:id;primaryKey" json:"id"`
	Name              string         `gorm:"column:name" json:"name"`
	Description       string         `gorm:"column:description" json:"description"`
	ProductCategoryID int            `gorm:"column:product_category_id" json:"productCategoryId"`
	MinSalePrice      float64        `gorm:"column:min_sale_price" json:"minSalePrice"`
	MaxSalePrice      float64        `gorm:"column:max_sale_price" json:"maxSalePrice"`
	DefaultSalePrice  float64        `gorm:"column:default_sale_price" json:"defaultSalePrice"`
	CurrentSalePrice  float64        `gorm:"column:current_sale_price" json:"currentSalePrice"`
	OptimalStock      int            `gorm:"column:optimal_stock" json:"optimalStock"`
	CurrentStock      int            `gorm:"column:current_stock" json:"currentStock"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	ProductCategory ProductCategory `gorm:"foreignKey:ProductCategoryID;references:ID" json:"-"`
}

func (Product) TableName() string {
	return Products_TableName
}

// Orders table
type Order struct {
	ID                     int       `gorm:"column:id;primaryKey" json:"id"`
	UserID                 int       `gorm:"column:user_id" json:"userId"`
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

// OrderProducts table
type OrderProduct struct {
	OrderID          int `gorm:"column:order_id;primaryKey" json:"orderId"`
	ProductID        int `gorm:"column:product_id;primaryKey" json:"productId"`
	Quantity         int `gorm:"column:quantity" json:"quantity"`
	PerUnitSalePrice int `gorm:"column:per_unit_sale_price" json:"perUnitSalePrice"`

	Order   Order   `gorm:"foreignKey:OrderID;references:ID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}

func (OrderProduct) TableName() string {
	return OrderProducts_TableName
}

// Returns table
type Return struct {
	ID                      int       `gorm:"column:id;primaryKey" json:"id"`
	ProductID               int       `gorm:"column:product_id" json:"productId"`
	OrderID                 int       `gorm:"column:order_id" json:"orderId"`
	ReturnRequestDate       time.Time `gorm:"column:return_request_date;type:date" json:"returnRequestDate"`
	ReturnCollectedDate     time.Time `gorm:"column:return_collected_date;type:date" json:"returnCollectedDate"`
	ReturnItemReachBackDate time.Time `gorm:"column:return_item_reach_back_date;type:date" json:"returnItemReachBackDate"`
	RefundTransactionID     string    `gorm:"column:refund_transaction_id" json:"refundTransactionId"`

	Order             Order             `gorm:"foreignKey:OrderID;references:ID" json:"-"`
	Product           Product           `gorm:"foreignKey:ProductID;references:ID" json:"-"`
	RefundTransaction RefundTransaction `gorm:"foreignKey:RefundTransactionID;references:ID" json:"-"`
}

func (Return) TableName() string {
	return Returns_TableName
}

// IncomingTransactions table
type IncomingTransaction struct {
	ID                string    `gorm:"column:id;primaryKey" json:"id"`
	PaymentAmount     float64   `gorm:"column:payment_amount" json:"paymentAmount"`
	TransactionTime   time.Time `gorm:"column:transaction_time" json:"transactionTime"`
	TransactionMethod string    `gorm:"column:transaction_method" json:"transactionMethod"`
	Status            string    `gorm:"column:status" json:"status"`
}

func (IncomingTransaction) TableName() string {
	return IncomingTransactions_TableName
}

// RefundTransactions table
type RefundTransaction struct {
	ID                string    `gorm:"column:id;primaryKey" json:"id"`
	PaymentAmount     float64   `gorm:"column:payment_amount" json:"paymentAmount"`
	TransactionTime   time.Time `gorm:"column:transaction_time" json:"transactionTime"`
	TransactionMethod string    `gorm:"column:transaction_method" json:"transactionMethod"`
	Status            string    `gorm:"column:status" json:"status"`
}

func (RefundTransaction) TableName() string {
	return RefundTransactions_TableName
}

// CartItems table
type CartItem struct {
	UserID             int     `gorm:"column:user_id;primaryKey" json:"userId"`
	ProductID          int     `gorm:"column:product_id;primaryKey" json:"productId"`
	Quantity           int     `gorm:"column:quantity" json:"quantity"`
	PriceWhenPutInCart float64 `gorm:"column:price_when_put_in_cart" json:"priceWhenPutInCart"`

	User    User    `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}

func (CartItem) TableName() string {
	return CartItems_TableName
}

type ProductRating struct {
	UserID    int     `gorm:"column:user_id;primaryKey" json:"userId"`
	ProductID int     `gorm:"column:product_id;primaryKey" json:"productId"`
	Rating    float64 `gorm:"column:rating" json:"rating"`

	User    User    `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}

func (ProductRating) TableName() string {
	return ProductRating_TableName
}

type ProductReview struct {
	UserID    int            `gorm:"column:user_id;primaryKey" json:"userId"`
	ProductID int            `gorm:"column:product_id;primaryKey" json:"productId"`
	Review    string         `gorm:"column:review" json:"review"`
	PicLinks  pq.StringArray `gorm:"column:pic_links;type:varchar(100)[]" json:"picLinks"`

	User    User    `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}

func (ProductReview) TableName() string {
	return ProductReviews_TableName
}

type ProductCategory struct {
	ID   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (ProductCategory) TableName() string {
	return ProductCategory_TableName
}

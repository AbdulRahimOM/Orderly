package constants

import "time"

const (
	Role   = "role"
	UserID = "user_id"

	// Pagination, Get request
	Param_IncludeDeleted = "includeDeleted"
	Param_Limit          = "limit"
	Param_Page           = "p"
	Param_Offset         = "offset"

	// Roles
	RoleSuperAdmin = "superadmin"
	RoleAdmin      = "admin"
	RoleUser       = "user"
	UnverifiedUser = "unverifiedUser"

	// Development purpose:
	UniversalPassword = "password"

	// Order Status
	PaymentMethod_COD    = "COD"
	PaymentMethod_Online = "Online Payment"
	Order_Pending        = "pending"
	Order_Placed         = "placed"
	Order_Cancelled      = "cancelled"
	Order_Delivered      = "delivered"

	//Token
	DefaultTokenExpiry          = time.Hour * 24 * 30 // 1 month
	Privilege                   = "privilege"
	Privilege_User_Management   = "user_manager"
	Privilege_Sales_Manager     = "sales_manager"
	Privilege_Inventory_Manager = "inventory_manager"
)

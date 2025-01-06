package constants

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
)

package usecase

import (
	"context"
	"orderly/internal/domain/request"
	"orderly/internal/domain/response"
)

type Usecase interface {
	CommonUsecase
	// Session Management
	SuperAdminSignin(ctx context.Context, req *request.SigninReq) *response.Response
	AdminSignin(ctx context.Context, req *request.SigninReq) *response.Response
	UserSignin(ctx context.Context, req *request.SigninReq) *response.Response
	UserSignUpGetOTP(ctx context.Context, req *request.UserSignupReq) *response.Response
	UserSignUpVerifyOTP(ctx context.Context, req *request.VerifyOTPReq) *response.Response

	// Admin Management
	CreateAdmin(ctx context.Context, req *request.CreateAdminReq) *response.Response
	GetAdmins(ctx context.Context, req *request.GetRequest) *response.Response
	GetAdminByID(ctx context.Context, id string) *response.Response
	UpdateAdminByID(ctx context.Context,id string, req *request.UpdateAdminReq) *response.Response

	//Access Privilege Management
	CreateAccessPrivilege(ctx context.Context, req *request.AccessPrivilegeReq) *response.Response
	GetAccessPrivileges(ctx context.Context) *response.Response
	GetAccessPrivilegeByAdminID(ctx context.Context, adminID string) *response.Response
	DeleteAccessPrivilege(ctx context.Context, adminID string, privilegeID string) *response.Response

	// User Management
	GetUsers(ctx context.Context, req *request.GetRequest) *response.Response
	GetUserByID(ctx context.Context, id string) *response.Response

	// User Profile Management
	GetUserProfile(ctx context.Context) *response.Response

	// Address Management
	GetUserAddresses(ctx context.Context) *response.Response
	GetUserAddressByID(ctx context.Context, id string) *response.Response
	CreateUserAddress(ctx context.Context, req *request.UserAddressReq) *response.Response
	UpdateUserAddressByID(ctx context.Context, id string, req *request.UserAddressReq) *response.Response

	// Category Management
	CreateCategory(ctx context.Context, req *request.CategoryReq) *response.Response
	GetCategories(ctx context.Context, req *request.GetRequest) *response.Response
	GetCategoryByID(ctx context.Context, id int) *response.Response
	UpdateCategoryByID(ctx context.Context, id int, req *request.CategoryReq) *response.Response

	// Product Management
	CreateProduct(ctx context.Context, req *request.ProductReq) *response.Response
	GetProducts(ctx context.Context, req *request.GetRequest) *response.Response
	GetProductByID(ctx context.Context, id int) *response.Response
	UpdateProductByID(ctx context.Context, id int, req *request.UpdateProductReq) *response.Response

	// Stock Management
	GetProductStockByID(ctx context.Context, id int) *response.Response
	AddProductStockByID(ctx context.Context, id int, req *request.AddProductStockReq) *response.Response

	// Cart Management
	GetCart(ctx context.Context) *response.Response
	AddToCart(ctx context.Context, req *request.AddToCartReq) *response.Response
	RemoveProductFromCart(ctx context.Context, productId int) *response.Response
	UpdateCartItemQuantity(ctx context.Context, req *request.UpdateCartItemQuantityReq) *response.Response
	ClearCart(ctx context.Context) *response.Response

	// Order Management
	//User side
	CreateOrder(ctx context.Context, req *request.CreateOrderReq) *response.Response
	GetMyOrders(ctx context.Context, req *request.Pagination) *response.Response
	GetMyOrderDetails(ctx context.Context, orderID string) *response.Response
	CancelMyOrder(ctx context.Context, orderID string) *response.Response

	//Admin side
	GetOrders(ctx context.Context, req *request.Pagination) *response.Response
	GetOrderDetails(ctx context.Context, orderID string) *response.Response
	CancelOrder(ctx context.Context, orderID string) *response.Response
	MarkOrderAsDelivered(ctx context.Context, orderID string) *response.Response
}

// common functions
type CommonUsecase interface {
	SoftDeleteRecordByID(ctx context.Context, tableName string, id any) *response.Response
	UndoSoftDeleteRecordByID(ctx context.Context, tableName string, id any) *response.Response
	ActivateByID(ctx context.Context, tableName string, id any) *response.Response
	DeactivateByID(ctx context.Context, tableName string, id any) *response.Response
	HardDeleteRecordByID(ctx context.Context, tableName string, id any) *response.Response
}

package repositoryinterface

import (
	"context"
	"orderly/internal/domain/dto"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
)

type Repository interface {
	CommonRepositoryFunctions

	// Session Management
	GetCredential(ctx context.Context, username string, role string) (*dto.Credentials, error)

	// Admin Management
	GetAdmins(ctx context.Context, req *request.GetRequest) ([]dto.AdminInList, error)
	GetAdminByID(ctx context.Context, id string) (*dto.Admin, error)
	UpdateAdminByID(ctx context.Context, id string, req *request.UpdateAdminReq) error
	GetAccessPrivileges(ctx context.Context) ([]dto.AccessPrivilege, error)
	GetAccessPrivilegeByAdminID(ctx context.Context, adminID string) (*dto.AccessPrivilege, error)
	DeleteAccessPrivilege(ctx context.Context, adminID string, privilegeID string) error

	// User Management
	GetUserSignInDetails(ctx context.Context, id string) (*dto.UserSignInDetails, error)
	CheckIfUsernameEmailOrPhoneExistsInUser(ctx context.Context, username, email, phone string) (usernameExists, emailExists, phoneExists bool, err error)
	GetUsers(ctx context.Context, req *request.GetRequest) ([]dto.UserInList, error)
	GetUserByID(ctx context.Context, id string) (*dto.User, error)
	GetUserProfile(ctx context.Context) (*dto.UserProfile, error)
	GetUserAddresses(ctx context.Context) ([]dto.UserAddress, error)
	GetUserAddressByID(ctx context.Context, id string) (*dto.UserAddress, error)
	UpdateUserAddressByID(ctx context.Context, id string, req *request.UserAddressReq) error

	// Cart Management
	GetCart(ctx context.Context) ([]dto.Cart, error)
	AddToCart(ctx context.Context, req *request.AddToCartReq) error
	RemoveProductFromCart(ctx context.Context, productId int) error
	UpdateCartItemQuantity(ctx context.Context, req *request.UpdateCartItemQuantityReq) error
	ClearCart(ctx context.Context) error

	// Category Management
	GetCategories(ctx context.Context, req *request.GetRequest) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, id int) (*models.Category, error)
	UpdateCategoryByID(ctx context.Context, id int, req *request.CategoryReq) error

	// Product Management
	GetProducts(ctx context.Context, req *request.GetRequest) ([]dto.ProductInList, error)
	GetProductByID(ctx context.Context, id int) (*dto.Product, error)
	UpdateProductByID(ctx context.Context, id int, req *request.UpdateProductReq) error

	// Stock Management
	GetProductStockByID(ctx context.Context, id int) (int, error)
	AddProductStockByID(ctx context.Context, id int, addingQuantity int) (int, error)

	// Order Management
	GetOrders(ctx context.Context, req *request.Pagination) ([]*dto.OrderInListForAdmin, error)
	GetMyOrders(ctx context.Context, req *request.Pagination) ([]*dto.OrderInListForUser, error)
	GetOrderDetails(ctx context.Context, orderID string) (*dto.OrderDetailed, error)
	CancelOrder(ctx context.Context, id string) error
	GetCartItemsForOrder(ctx context.Context) ([]*dto.CartItemsForOrder, error)
	CreateOrder(ctx context.Context, order *models.Order, orderProducts []*models.OrderProduct) error
	CheckOrderBelongsToUser(ctx context.Context, orderID string) (bool, error)
	GetOrderStatus(ctx context.Context, orderID string) (string, bool, error)
	MarkOrderAsDelivered(ctx context.Context, orderID string) error
}

type CommonRepositoryFunctions interface {
	CreateRecord(ctx context.Context, record interface{}) error
	SoftDeleteRecordByID(ctx context.Context, tableName string, id any) error
	UndoSoftDeleteRecordByID(ctx context.Context, tableName string, id any) (string, error)
	ActivateByID(ctx context.Context, tableName string, id any) error
	DeactivateByID(ctx context.Context, tableName string, id any) error
	HardDeleteRecordByID(ctx context.Context, tableName string, id any) error
}

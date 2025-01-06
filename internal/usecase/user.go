package uc

import (
	"context"
	"fmt"
	"net/http"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	"orderly/pkg/utils/helper"

	"github.com/google/uuid"
)

func (uc *Usecase) GetCart(ctx context.Context) *response.Response {
	cart, err := uc.repo.GetCart(ctx)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, map[string]interface{}{
		"cart": cart,
	})
}

func (uc *Usecase) AddToCart(ctx context.Context, req *request.AddToCartReq) *response.Response {
	err := uc.repo.AddToCart(ctx, req)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) RemoveProductFromCart(ctx context.Context, productId int) *response.Response {
	err := uc.repo.RemoveProductFromCart(ctx, productId)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) UpdateCartItemQuantity(ctx context.Context, req *request.UpdateCartItemQuantityReq) *response.Response {
	err := uc.repo.UpdateCartItemQuantity(ctx, req)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) ClearCart(ctx context.Context) *response.Response {
	err := uc.repo.ClearCart(ctx)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) CreateOrder(ctx context.Context, req *request.CreateOrderReq) *response.Response {

	var (
		totalSum float64
	)

	cartItemsForOrder, err := uc.repo.GetCartItemsForOrder(ctx)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	if len(cartItemsForOrder) == 0 {
		return response.ErrorResponse(http.StatusBadRequest, "CART_EMPTY", fmt.Errorf("cart is empty"))
	}

	insufficientStockProductIDs := make([]int, 0)
	for _, item := range cartItemsForOrder {
		if item.Quantity > item.CurrentStock {
			insufficientStockProductIDs = append(insufficientStockProductIDs, item.ProductID)
		}
		totalSum += item.CurrentSalePrice * float64(item.Quantity)
	}
	if len(insufficientStockProductIDs) > 0 {
		return &response.Response{
			HttpStatusCode: http.StatusConflict,
			ResponseCode:   "INSUFFICIENT_STOCK",
			Error:          fmt.Errorf("insufficient stock for some products"),
			Data: map[string]interface{}{
				"insufficient_stock_productIDs": insufficientStockProductIDs,
			},
		}
	}

	status := constants.Order_Pending
	if req.PaymentMethod == constants.PaymentMethod_COD {
		status = constants.Order_Placed
	}

	address, err := uc.repo.GetUserAddressByID(ctx, req.AddressID)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting user address: %v", err))
	}

	order := &models.Order{
		ID:            uuid.New(),
		UserID:        helper.GetUserIdFromContext(ctx),
		TotalAmount:   totalSum,
		PaymentMethod: req.PaymentMethod,
		OrderStatus:   status,

		//shipping address
		House:    address.House,
		Street1:  address.Street1,
		Street2:  address.Street2,
		City:     address.City,
		Pincode:  address.Pincode,
		State:    address.State,
		Country:  address.Country,
		Landmark: address.Landmark,
	}

	orderProducts := make([]*models.OrderProduct, 0)
	for _, item := range cartItemsForOrder {
		orderProducts = append(orderProducts, &models.OrderProduct{
			OrderID:      order.ID,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			SellingPrice: item.CurrentSalePrice,
		})
	}

	err = uc.repo.CreateOrder(ctx, order, orderProducts)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	//can add: send email to user, invoice generation, etc.

	return response.SuccessResponse(201, respcode.Created, map[string]interface{}{
		"id": order.ID,
	})
}

func (uc *Usecase) GetMyOrders(ctx context.Context, req *request.Pagination) *response.Response {
	orders, err := uc.repo.GetMyOrders(ctx, req)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, map[string]interface{}{
		"orders": orders,
	})
}

func (uc *Usecase) GetMyOrderDetails(ctx context.Context, orderId string) *response.Response {
	//check if order belongs to user
	belongsToUser, err := uc.repo.CheckOrderBelongsToUser(ctx, orderId)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	if !belongsToUser {
		return response.ErrorResponse(http.StatusForbidden, respcode.Forbidden, fmt.Errorf("order does not belong to user"))
	}

	orderDetails, err := uc.repo.GetOrderDetails(ctx, orderId)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, map[string]interface{}{
		"order_details": orderDetails,
	})
}

func (uc *Usecase) CancelMyOrder(ctx context.Context, orderId string) *response.Response {
	//check if order belongs to user
	belongsToUser, err := uc.repo.CheckOrderBelongsToUser(ctx, orderId)
	if err != nil {
		return response.DBErrorResponse(err)
	}
	if !belongsToUser {
		return response.ErrorResponse(http.StatusForbidden, respcode.Forbidden, fmt.Errorf("order does not belong to user"))
	}

	return uc.cancelOrder(ctx, orderId)
}

func (uc *Usecase) cancelOrder(ctx context.Context, orderId string) *response.Response {

	orderStatus, paymentDone, err := uc.repo.GetOrderStatus(ctx, orderId)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	switch orderStatus {
	case constants.Order_Pending, constants.Order_Placed:
		goto gate_CancelOrder
	default:
		return response.ErrorResponse(400, respcode.BadRequest, fmt.Errorf("cannot mark order as delivered: details: order status: %s, payment done: %v", orderStatus, paymentDone))
	}

gate_CancelOrder:
	err = uc.repo.CancelOrder(ctx, orderId)
	if err != nil {
		return response.DBErrorResponse(err)
	}
	if paymentDone {
		//refund the amount	//todo
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

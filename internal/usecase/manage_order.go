package uc

import (
	"context"
	"fmt"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
)

func (u *Usecase) GetOrders(ctx context.Context, req *request.Pagination) *response.Response {
	orders, err := u.repo.GetOrders(ctx, req)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, orders)
}

func (u *Usecase) GetOrderDetails(ctx context.Context, id string) *response.Response {
	order, err := u.repo.GetOrderDetails(ctx, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, order)
}

func (u *Usecase) CancelOrder(ctx context.Context, id string) *response.Response {
	return u.cancelOrder(ctx, id)
}

func (u *Usecase) MarkOrderAsDelivered(ctx context.Context, id string) *response.Response {
	orderStatus, paymentDone, err := u.repo.GetOrderStatus(ctx, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}
	if !paymentDone {
		return response.ErrorResponse(400, "PAYMENT_NOT_DONE", fmt.Errorf("cannot mark order as delivered: payment not done"))
	} else if orderStatus != constants.Order_Placed {
		return response.ErrorResponse(400, respcode.BadRequest, fmt.Errorf("cannot mark order as delivered: details: order status: %s, payment done: %v", orderStatus, paymentDone))
	}

	//allow to mark as delivered
	err = u.repo.MarkOrderAsDelivered(ctx, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}
	return response.SuccessResponse(200, respcode.Success, nil)

}

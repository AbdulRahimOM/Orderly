package uc

import (
	"context"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
)

func (uc *Usecase) GetCart(ctx context.Context) *response.Response {
	cart,err:=uc.repo.GetCart(ctx)
	if err!=nil{
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, map[string]interface{}{
		"cart": cart,
	})
}

func (uc *Usecase) AddToCart(ctx context.Context, req *request.AddToCartReq) *response.Response {
	err:=uc.repo.AddToCart(ctx,req)
	if err!=nil{
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) RemoveProductFromCart(ctx context.Context, productId int) *response.Response {
	err:=uc.repo.RemoveProductFromCart(ctx,productId)
	if err!=nil{
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) UpdateCartItemQuantity(ctx context.Context, req *request.UpdateCartItemQuantityReq) *response.Response {
	err:=uc.repo.UpdateCartItemQuantity(ctx,req)
	if err!=nil{
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) ClearCart(ctx context.Context) *response.Response {
	err:=uc.repo.ClearCart(ctx)
	if err!=nil{
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}
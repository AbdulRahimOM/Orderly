package usecase

import (
	"context"
	"orderly/internal/domain/request"
	"orderly/internal/domain/response"
)

type AccountUsecase interface {
	SuperAdminSignin(ctx context.Context, req *request.SigninReq) *response.Response
}

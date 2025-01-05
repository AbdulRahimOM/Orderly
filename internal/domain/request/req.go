package request

import (
	"log"
	"net/http"
	"orderly/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

const (
	defaultLimit = 10
)

type SigninReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type VerifyOTPReq struct {
	OTP string `json:"otp" validate:"required"`
}

type UserSignupReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required,e164"`
}

type CreateAdminReq struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	Phone       string `json:"phone" validate:"required,e164"`
	Designation string `json:"designation"`
}

type UpdateAdminReq struct {
	Email       string `json:"email" validate:"omitempty,email"`
	Name        string `json:"name" validate:"omitempty"`
	Phone       string `json:"phone" validate:"omitempty,e164"`
	Designation string `json:"designation" validate:"omitempty"`
}

type Pagination struct {
	Page   int `query:"page"`
	Limit  int `query:"limit"`
	Offset int `query:"-"`
}

type GetRequest struct {
	IsDeleted bool `query:"is_deleted"`
	Pagination
}

func GetListRequest(ctx *fiber.Ctx) (*GetRequest, error) {
	req := new(GetRequest)
	if err := ctx.QueryParser(req); err != nil {
		log.Println("error parsing request:", err)
		return nil, response.Response{
			HttpStatusCode: http.StatusBadRequest,
			Status:         false,
			ResponseCode:   "URL_QUERY_BINDING_ERROR",
			Error:          err,
		}.WriteToJSON(ctx)
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = defaultLimit
	}

	req.Offset = (req.Page - 1) * req.Limit
	return req, nil
}

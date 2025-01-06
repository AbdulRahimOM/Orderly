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

func GetPaginations(ctx *fiber.Ctx) (*Pagination, error) {
	req := new(Pagination)
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
package handler

import (
	"orderly/internal/domain/request"
	"orderly/internal/domain/response"
	"orderly/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetCart(c *fiber.Ctx) error {
	response := h.uc.GetCart(c.Context())
	return response.WriteToJSON(c)
}

func (h *Handler) AddToCart(c *fiber.Ctx) error {
	req := new(request.AddToCartReq)
	if ok, errResponse := validation.BindAndValidateJSONRequest(c, req); !ok {
		return errResponse
	}

	response := h.uc.AddToCart(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) RemoveProductFromCart(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}

	response := h.uc.RemoveProductFromCart(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) UpdateCartItemQuantity(c *fiber.Ctx) error {
	req := new(request.UpdateCartItemQuantityReq)
	if ok, errResponse := validation.BindAndValidateJSONRequest(c, req); !ok {
		return errResponse
	}

	response := h.uc.UpdateCartItemQuantity(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) ClearCart(c *fiber.Ctx) error {
	response := h.uc.ClearCart(c.Context())
	return response.WriteToJSON(c)
}

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	req := new(request.CreateOrderReq)
	if ok, errResponse := validation.BindAndValidateJSONRequest(c, req); !ok {
		return errResponse
	}

	response := h.uc.CreateOrder(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetMyOrders(c *fiber.Ctx) error {
	req,errResponse:=request.GetPaginations(c)
	if req==nil{
		return errResponse
	}

	response := h.uc.GetMyOrders(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetMyOrderDetails(c *fiber.Ctx) error {
	id:=c.Params("id")
	response := h.uc.GetMyOrderDetails(c.Context(), id)
	return response.WriteToJSON(c)
}
	
func (h *Handler) CancelMyOrder(c *fiber.Ctx) error {
	id:=c.Params("id")
	response := h.uc.CancelMyOrder(c.Context(), id)
	return response.WriteToJSON(c)
}

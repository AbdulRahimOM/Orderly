package handler

import (
	"orderly/internal/domain/request"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetOrders(c *fiber.Ctx) error {
	req, errResponse := request.GetPaginations(c)
	if req == nil {
		return errResponse
	}

	response := h.uc.GetOrders(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) CancelOrder (c *fiber.Ctx) error {
	id := c.Params("id")
	response := h.uc.CancelOrder(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) GetOrderDetails (c *fiber.Ctx) error {
	id := c.Params("id")
	response := h.uc.GetOrderDetails(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) MarkOrderAsDelivered (c *fiber.Ctx) error {
	id := c.Params("id")
	response := h.uc.MarkOrderAsDelivered(c.Context(), id)
	return response.WriteToJSON(c)
}
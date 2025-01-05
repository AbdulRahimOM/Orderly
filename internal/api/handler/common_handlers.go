package handler

import (
	"orderly/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SoftDeleteRecordByID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
		}
		response := h.uc.SoftDeleteRecordByID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) SoftDeleteRecordByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.SoftDeleteRecordByUUID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) UndoSoftDeleteRecordByID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
		}
		response := h.uc.UndoSoftDeleteRecordByID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) UndoSoftDeleteRecordByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.UndoSoftDeleteRecordByUUID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) ActivateByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.ActivateByUUID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) DeactivateByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.DeactivateByUUID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

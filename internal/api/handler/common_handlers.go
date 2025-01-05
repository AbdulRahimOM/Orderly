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
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
		}
		response := h.uc.UndoSoftDeleteRecordByID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) BlockByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.BlockByUUID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) UnblockByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.UnblockByUUID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

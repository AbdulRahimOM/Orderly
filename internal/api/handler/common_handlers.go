package handler

import (
	"orderly/internal/domain/response"
	jwttoken "orderly/pkg/jwt-token"

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
		response := h.uc.SoftDeleteRecordByID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) SoftDeleteAccountByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.SoftDeleteRecordByID(c.Context(), tableName, id)
		if response.Status {
			jwttoken.RevokeExistingAuthToken(id)
		}
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
		response := h.uc.UndoSoftDeleteRecordByID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) ActivateByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.ActivateByID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) DeactivateByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.DeactivateByID(c.Context(), tableName, id)
		if response.Status {
			jwttoken.RevokeExistingAuthToken(id)
		}
		return response.WriteToJSON(c)
	}
}

func (h *Handler) DeactivateAccountByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.DeactivateByID(c.Context(), tableName, id)
		if response.Status {
			jwttoken.RevokeExistingAuthToken(id)
		}
		return response.WriteToJSON(c)
	}
}

func (h *Handler) HardDeleteRecordByID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
		}
		response := h.uc.HardDeleteRecordByID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) HardDeleteRecordByUUID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		response := h.uc.HardDeleteRecordByID(c.Context(), tableName, id)
		return response.WriteToJSON(c)
	}
}

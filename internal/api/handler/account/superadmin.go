package accounthandler

import (
	"orderly/internal/domain/request"
	"orderly/internal/domain/response"
	"orderly/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SuperAdminSignin(c *fiber.Ctx) error {

	req := new(request.SuperAdminSigninReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.SuperAdminSignin(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) CreateAdmin(c *fiber.Ctx) error {

	req := new(request.CreateAdminReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.CreateAdmin(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetAdmins(c *fiber.Ctx) error {
	req, errResponse := request.GetListRequest(c)
	if req == nil {
		return errResponse
	}

	response := h.uc.GetAdmins(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetAdminByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}
	response := h.uc.GetAdminByID(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) UpdateAdminByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}

	req := new(request.UpdateAdminReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.UpdateAdminByID(c.Context(), id, req)
	return response.WriteToJSON(c)
}

func (h *Handler) SoftDeleteRecordByID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
		}
		response := h.uc.SoftDeleteRecordByID(c.Context(),tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) UndoSoftDeleteRecordByID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
		}
		response := h.uc.UndoSoftDeleteRecordByID(c.Context(),tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) BlockByID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
		}
		response := h.uc.BlockByID(c.Context(),tableName, id)
		return response.WriteToJSON(c)
	}
}

func (h *Handler) UnblockByID(tableName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
		}
		response := h.uc.UnblockByID(c.Context(),tableName, id)
		return response.WriteToJSON(c)
	}
}

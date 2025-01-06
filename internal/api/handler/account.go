package handler

import (
	"orderly/internal/domain/request"
	"orderly/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

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
	id := c.Params("id")
	response := h.uc.GetAdminByID(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) UpdateAdminByID(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(request.UpdateAdminReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.UpdateAdminByID(c.Context(), id, req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	req, errResponse := request.GetListRequest(c)
	if req == nil {
		return errResponse
	}

	response := h.uc.GetUsers(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	response := h.uc.GetUserByID(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) GetUserProfile(c *fiber.Ctx) error {
	response := h.uc.GetUserProfile(c.Context())
	return response.WriteToJSON(c)
}

func (h *Handler) GetUserAddresses(c *fiber.Ctx) error {
	response := h.uc.GetUserAddresses(c.Context())
	return response.WriteToJSON(c)
}

func (h *Handler) GetUserAddressByID(c *fiber.Ctx) error {
	id := c.Params("id")
	response := h.uc.GetUserAddressByID(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) CreateUserAddress(c *fiber.Ctx) error {
	req := new(request.UserAddressReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.CreateUserAddress(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) UpdateUserAddressByID(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(request.UserAddressReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.UpdateUserAddressByID(c.Context(), id, req)
	return response.WriteToJSON(c)
}

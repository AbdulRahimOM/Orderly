package handler

import (
	"orderly/internal/domain/request"
	jwttoken "orderly/pkg/jwt-token"
	"orderly/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateAccessPrivilege(c *fiber.Ctx) error {
	req := new(request.AccessPrivilegeReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.CreateAccessPrivilege(c.Context(), req)
	if response.Status {
		jwttoken.RevokeExistingAuthToken(req.AdminID.String())
	}
	return response.WriteToJSON(c)
}

func (h *Handler) GetAccessPrivileges(c *fiber.Ctx) error {
	response := h.uc.GetAccessPrivileges(c.Context())
	return response.WriteToJSON(c)
}

func (h *Handler) GetAccessPrivilegeByAdminID(c *fiber.Ctx) error {
	id := c.Params("admin_id")
	response := h.uc.GetAccessPrivilegeByAdminID(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) DeleteAccessPrivilege(c *fiber.Ctx) error {
	id := c.Params("admin_id")
	privilege := c.Params("privilege")
	response := h.uc.DeleteAccessPrivilege(c.Context(), id, privilege)
	if response.Status {
		jwttoken.RevokeExistingAuthToken(id)
	}
	return response.WriteToJSON(c)
}

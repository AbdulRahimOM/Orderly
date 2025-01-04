package accounthandler

import (
	"orderly/internal/domain/request"
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
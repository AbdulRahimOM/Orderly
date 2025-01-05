package accounthandler

import (
	"orderly/internal/domain/request"
	"orderly/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SuperAdminSignin(c *fiber.Ctx) error {

	req := new(request.SigninReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.SuperAdminSignin(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) AdminSignin(c *fiber.Ctx) error {
	
	req := new(request.SigninReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.AdminSignin(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) UserSignIn(c *fiber.Ctx) error {
	
	req := new(request.SigninReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.UserSignin(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) UserSignUpGetOTP(c *fiber.Ctx) error {
	
	req := new(request.UserSignupReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.UserSignUpGetOTP(c.Context(), req)
	return response.WriteToJSON(c)
}
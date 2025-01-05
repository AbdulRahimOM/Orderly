package handler

import (
	"fmt"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
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

func (h *Handler) UserSignUpVerifyOTP(c *fiber.Ctx) error {
	//verify token
	role := c.Locals("role").(string)
	if role != constants.UnverifiedUser {
		return response.ErrorResponse(403, respcode.Forbidden, fmt.Errorf("user is not an unverified user")).WriteToJSON(c)
	}

	req := new(request.VerifyOTPReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.UserSignUpVerifyOTP(c.Context(), req)
	return response.WriteToJSON(c)
}

package middleware

import (
	"orderly/internal/domain/response"
	jwttoken "orderly/pkg/jwt-token"

	"github.com/gofiber/fiber/v2"
)

func RevokeAuthTokenUsingParamId() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		jwttoken.RevokeExistingAuthToken(id)
		return response.SuccessResponse(fiber.StatusOK, "Token Revoked", nil).WriteToJSON(c)
	}
}

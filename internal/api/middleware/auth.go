package middleware

import (
	"fmt"
	cons "orderly/internal/domain/constants"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	jwttoken "orderly/pkg/jwt-token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ValidateJWT validates the jwt token in the Authorization header.
//
// If the token is valid, it sets the adminID, role and addlInfo in the context.
// Authentication middlewares for specific roles should be used after this middleware.
func ValidateJWT(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	tokenData, err := jwttoken.GetDataFromToken(tokenString)
	if err != nil {
		if err == jwttoken.ErrTokenExpired {
			return response.ErrorResponse(401, respcode.TokenExpired, fmt.Errorf("token expired: %w", err)).WriteToJSON(c)
		} else {
			return response.ErrorResponse(401, respcode.InvalidToken, fmt.Errorf("invalid token: %w", err)).WriteToJSON(c)
		}
	}
	c.Locals(cons.UserID, tokenData.Id)
	c.Locals(cons.Role, tokenData.Role)

	addlInfo, ok := tokenData.AddlInfo.(map[string]interface{})
	if !ok {
		if tokenData.AddlInfo != nil {
			return response.BugResponse(fmt.Errorf("tokenData.AddlInfo is not of type map[string]interface{}. Value of addlInfo=%v", tokenData.AddlInfo)).WriteToJSON(c)
			// } else {
			// 	log.Debug("tokenData.AddlInfo is nil")
		}
	}
	for key, value := range addlInfo {
		c.Locals(key, value)
	}
	return c.Next()
}

// ValidateAnyOfTheseRoles validates whether the role in the context is matching with any of the roles passed.
func ValidateAnyOfTheseRoles(roles ...string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		role := c.Locals(cons.Role).(string)
		for _, r := range roles {
			if r == role {
				return c.Next()
			}
		}
		return response.ErrorResponse(403, respcode.Forbidden, fmt.Errorf("role %s is not allowed to access this resource", role)).WriteToJSON(c)
	}
}
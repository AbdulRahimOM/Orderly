package middleware

import (
	"fmt"
	"orderly/internal/domain/constants"
	cons "orderly/internal/domain/constants"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	jwttoken "orderly/pkg/jwt-token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
		TokenExpired = "TOKEN_EXPIRED"
		InvalidToken = "INVALID_TOKEN"
		TokenRevoke  = "TOKEN_REVOKED"
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
			return response.ErrorResponse(401, TokenExpired, fmt.Errorf("token expired: %w", err)).WriteToJSON(c)
		} else if err == jwttoken.ErrTokenIsInvalid {
			return response.ErrorResponse(401, InvalidToken, fmt.Errorf("invalid token: %w", err)).WriteToJSON(c)
		} else if err == jwttoken.ErrTokenRevoked {
			return response.ErrorResponse(401, TokenRevoke, fmt.Errorf("token revoked: %w", err)).WriteToJSON(c)
		} else {
			return response.ErrorResponse(401, InvalidToken, fmt.Errorf("invalid token: %w", err)).WriteToJSON(c)
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

func HavePrivilege(privilege string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		privileges, ok := c.Locals(constants.Privilege).([]any)
		if !ok {
			if c.Locals(constants.Privilege) == nil {
				return response.UnauthorizedResponse(fmt.Errorf("special privileges in token is nil. required privilege=%s", privilege)).WriteToJSON(c)
			}
			return response.BugResponse(fmt.Errorf("special privileges in token is not of type []any. Value of special privileges=%v", c.Locals(constants.Privilege))).WriteToJSON(c)
		}
		if !contains(privileges, privilege) {
			return response.UnauthorizedResponse(fmt.Errorf("special privileges in token does not contain required privilege. required privilege=%s", privilege)).WriteToJSON(c)
		}
		return c.Next()
	}
}

func contains(slice []any, requiredString string) bool {
	for _, element := range slice {
		if str, ok := element.(string); !ok {
			break
		} else if str == requiredString {
			return true
		}
	}
	return false
}

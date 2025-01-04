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

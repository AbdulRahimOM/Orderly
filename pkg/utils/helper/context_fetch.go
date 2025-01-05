package helper

import (
	"context"
	"orderly/internal/domain/constants"
)

func GetRoleFromContext(ctx context.Context) string {
	role := ctx.Value(constants.Role)
	if role != nil {
		return role.(string)
	}

	return ""
}

func GetUserIdFromContext(ctx context.Context) int {
	userID := ctx.Value(constants.UserID)
	if userID != nil {
		return userID.(int)
	}

	return 0
}
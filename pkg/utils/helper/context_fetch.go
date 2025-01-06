package helper

import (
	"context"
	"orderly/internal/domain/constants"

	"github.com/google/uuid"
)

func GetRoleFromContext(ctx context.Context) string {
	role := ctx.Value(constants.Role)
	if role != nil {
		return role.(string)
	}

	return ""
}

func GetUserIdFromContext(ctx context.Context) uuid.UUID {
	userID,ok := ctx.Value(constants.UserID).(uuid.UUID)
	if !ok {
		return uuid.Nil
	}

	return userID
}
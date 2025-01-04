package accountuc

import (
	"context"
	"net/http"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

// SoftDeleteRecordByID
func (uc *AccountUC) SoftDeleteRecordByID(ctx context.Context, tableName string, id int) *response.Response {
	err := uc.repo.SoftDeleteRecordByID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *AccountUC) UndoSoftDeleteRecordByID(ctx context.Context, tableName string, id int) *response.Response {
	responsecode, err := uc.repo.UndoSoftDeleteRecordByID(ctx, tableName, id)
	if err != nil {
		if responsecode == respcode.UniqueFieldViolation {
			return response.ErrorResponse(http.StatusConflict, respcode.UniqueFieldViolation, err)
		}
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *AccountUC) BlockByID(ctx context.Context, tableName string, id int) *response.Response {
	err := uc.repo.BlockByID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *AccountUC) UnblockByID(ctx context.Context, tableName string, id int) *response.Response {
	err := uc.repo.UnblockByID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

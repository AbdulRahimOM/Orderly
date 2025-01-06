package uc

import (
	"context"
	"net/http"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

// SoftDeleteRecordByID
func (uc *Usecase) SoftDeleteRecordByID(ctx context.Context, tableName string, id any) *response.Response {
	err := uc.repo.SoftDeleteRecordByID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *Usecase) UndoSoftDeleteRecordByID(ctx context.Context, tableName string, id any) *response.Response {
	responsecode, err := uc.repo.UndoSoftDeleteRecordByID(ctx, tableName, id)
	if err != nil {
		if responsecode == respcode.UniqueFieldViolation {
			return response.ErrorResponse(http.StatusConflict, respcode.UniqueFieldViolation, err)
		}
		if responsecode == respcode.NotFound {
			return response.ErrorResponse(http.StatusNotFound, respcode.NotFound, err)
		}
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}


func (uc *Usecase) ActivateByID(ctx context.Context, tableName string, id any) *response.Response {
	err := uc.repo.ActivateByID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *Usecase) DeactivateByID(ctx context.Context, tableName string, id string) *response.Response {
	err := uc.repo.DeactivateByID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *Usecase) HardDeleteRecordByID(ctx context.Context, tableName string, id any) *response.Response {
	err := uc.repo.HardDeleteRecordByID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}


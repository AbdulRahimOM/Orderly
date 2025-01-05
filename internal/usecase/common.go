package uc

import (
	"context"
	"net/http"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

// SoftDeleteRecordByID
func (uc *Usecase) SoftDeleteRecordByID(ctx context.Context, tableName string, id int) *response.Response {
	err := uc.repo.SoftDeleteRecordByID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *Usecase) UndoSoftDeleteRecordByID(ctx context.Context, tableName string, id int) *response.Response {
	responsecode, err := uc.repo.UndoSoftDeleteRecordByID(ctx, tableName, id)
	if err != nil {
		if responsecode == respcode.UniqueFieldViolation {
			return response.ErrorResponse(http.StatusConflict, respcode.UniqueFieldViolation, err)
		}
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *Usecase) SoftDeleteRecordByUUID(ctx context.Context, tableName string, id string) *response.Response {
	err := uc.repo.SoftDeleteRecordByUUID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *Usecase) UndoSoftDeleteRecordByUUID(ctx context.Context, tableName string, id string) *response.Response {
	responsecode, err := uc.repo.UndoSoftDeleteRecordByUUID(ctx, tableName, id)
	if err != nil {
		if responsecode == respcode.UniqueFieldViolation {
			return response.ErrorResponse(http.StatusConflict, respcode.UniqueFieldViolation, err)
		}
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *Usecase) BlockByUUID(ctx context.Context, tableName string, id string) *response.Response {
	err := uc.repo.BlockByUUID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

func (uc *Usecase) UnblockByUUID(ctx context.Context, tableName string, id string) *response.Response {
	err := uc.repo.UnblockByUUID(ctx, tableName, id)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(fiber.StatusOK, respcode.Success, nil)
}

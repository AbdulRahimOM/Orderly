package uc

import (
	"context"
	"fmt"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	"orderly/pkg/utils/hashpassword"
	"orderly/pkg/utils/helper"
	"time"

	"github.com/google/uuid"
)

const (
	randomPasswordLength = 8
)

func (uc *Usecase) CreateAdmin(ctx context.Context, req *request.CreateAdminReq) *response.Response {

	password := helper.GenerateRandomAlphanumeric(randomPasswordLength)
	hashedPw, err := hashpassword.GetHashedPassword(password)
	if err != nil {
		return response.InternalServerErrorResponse(fmt.Errorf("error hashing password: %v", err))
	}
	// Create a new admin
	admin := models.Admin{
		ID:             uuid.New(),
		Name:           req.Name,
		Username:       req.Username,
		HashedPassword: hashedPw,
		Email:          req.Email,
		Phone:          req.Phone,
		Designation:    req.Designation,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Save the admin
	if err := uc.repo.CreateRecord(ctx, &admin); err != nil {
		return admin.GetResponseFromDBError(err)
	}

	return response.CreatedResponse(admin.ID)
}

func (uc *Usecase) GetAdmins(ctx context.Context, req *request.GetRequest) *response.Response {
	admins, err := uc.repo.GetAdmins(ctx, req)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting admins: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, admins)
}

func (uc *Usecase) GetAdminByID(ctx context.Context, id string) *response.Response {
	admin, err := uc.repo.GetAdminByID(ctx, id)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting admin: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, admin)
}

func (uc *Usecase) UpdateAdminByID(ctx context.Context, id string, req *request.UpdateAdminReq) *response.Response {
	err := uc.repo.UpdateAdminByID(ctx, id, req)
	if err != nil {
		return models.Admin{}.GetResponseFromDBError(err)
	}
	return response.SuccessResponse(200, respcode.Success, nil)
}

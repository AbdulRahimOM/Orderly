package uc

import (
	"context"
	"fmt"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	jwttoken "orderly/pkg/jwt-token"
	"orderly/pkg/utils/email"
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

	// Send email to the admin
	err = email.SendCredentials(admin.Email, admin.Username, password)
	if err != nil {
		return response.InternalServerErrorResponse(fmt.Errorf("error sending email: %v", err))
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

func (uc *Usecase) GetUsers(ctx context.Context, req *request.GetRequest) *response.Response {
	users, err := uc.repo.GetUsers(ctx, req)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting users: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, users)
}

func (uc *Usecase) GetUserByID(ctx context.Context, id string) *response.Response {
	user, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting user: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, user)
}

func (uc *Usecase) GetUserProfile(ctx context.Context) *response.Response {
	user, err := uc.repo.GetUserProfile(ctx)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting user profile: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, user)
}

func (uc *Usecase) GetUserAddresses(ctx context.Context) *response.Response {
	addresses, err := uc.repo.GetUserAddresses(ctx)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting user addresses: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, addresses)
}

func (uc *Usecase) GetUserAddressByID(ctx context.Context, id string) *response.Response {
	address, err := uc.repo.GetUserAddressByID(ctx, id)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting user address: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, address)
}

func (uc *Usecase) CreateUserAddress(ctx context.Context, req *request.UserAddressReq) *response.Response {
	userID := helper.GetUserIdFromContext(ctx)
	address := models.Address{
		ID:       uuid.New(),
		UserID:   userID,
		House:    req.House,
		Street1:  req.Street1,
		Street2:  req.Street2,
		City:     req.City,
		State:    req.State,
		Pincode:  req.Pincode,
		Country:  req.Country,
		Landmark: req.Landmark,
	}

	if err := uc.repo.CreateRecord(ctx, &address); err != nil {
		return response.DBErrorResponse(fmt.Errorf("error creating user address: %v", err))
	}

	return response.CreatedResponse(address.ID)
}

func (uc *Usecase) UpdateUserAddressByID(ctx context.Context, id string, req *request.UserAddressReq) *response.Response {
	err := uc.repo.UpdateUserAddressByID(ctx, id, req)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error updating user address: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) CreateAccessPrivilege(ctx context.Context, req *request.AccessPrivilegeReq) *response.Response {
	accessPrivilege := models.AdminPrivilege{
		AdminID:    req.AdminID,
		AccessRole: req.AccessRole,
	}

	if err := uc.repo.CreateRecord(ctx, &accessPrivilege); err != nil {
		return accessPrivilege.GetResponseFromDBError(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (uc *Usecase) GetAccessPrivileges(ctx context.Context) *response.Response {
	accessPrivileges, err := uc.repo.GetAccessPrivileges(ctx)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting access privileges: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, accessPrivileges)
}

func (uc *Usecase) GetAccessPrivilegeByAdminID(ctx context.Context, adminID string) *response.Response {
	accessPrivileges, err := uc.repo.GetAccessPrivilegeByAdminID(ctx, adminID)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error getting access privileges: %v", err))
	}
	return response.SuccessResponse(200, respcode.Success, accessPrivileges)
}

func (uc *Usecase) DeleteAccessPrivilege(ctx context.Context, adminID string, privilege string) *response.Response {
	err := uc.repo.DeleteAccessPrivilege(ctx, adminID, privilege)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error deleting access privileges: %v", err))
	}

	jwttoken.RevokeExistingAuthToken(adminID)

	return response.SuccessResponse(200, respcode.Success, nil)
}

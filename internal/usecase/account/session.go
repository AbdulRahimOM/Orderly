package accountuc

import (
	"context"
	"fmt"
	"net/http"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	"orderly/internal/infrastructure/config"
	repo "orderly/internal/repository"
	jwttoken "orderly/pkg/jwt-token"
	"orderly/pkg/utils/hashpassword"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	USERNAME_NOT_REGISTERED = "USERNAME_NOT_REGISTERED"
	PASSWORD_MISMATCH       = "PASSWORD_MISMATCH"
	defaultTokenExpiry      = time.Hour * 24 * 7 // 1 week
)

var (
	invalidUsernameResponse = response.ErrorResponse(http.StatusUnauthorized, USERNAME_NOT_REGISTERED, fmt.Errorf("Username not registered"))
	invalidPasswordResponse = response.ErrorResponse(http.StatusUnauthorized, PASSWORD_MISMATCH, fmt.Errorf("Invalid password"))
)

func comparePassword(hashedPassword string, password string) error {
	if config.Configs.Dev_AllowUniversalPassword && password == constants.UniversalPassword {
		fmt.Println("Dev: Allowing universal password")
		return nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (uc *AccountUC) SuperAdminSignin(ctx context.Context, req *request.SigninReq) *response.Response {
	superAdminCredentials, err := uc.repo.GetCredential(ctx, req.Username, constants.RoleSuperAdmin)
	if err != nil {
		if err == repo.ErrRecordNotFound {
			return invalidUsernameResponse
		}
		return response.DBErrorResponse(err)
	}

	//compare password
	if err := comparePassword(superAdminCredentials.HashedPassword, req.Password); err != nil {
		return invalidPasswordResponse
	}

	//generate token
	token, err := jwttoken.GenerateToken(superAdminCredentials.ID, constants.RoleSuperAdmin, "", nil, defaultTokenExpiry)
	if err != nil {
		return response.ErrorResponse(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"id":    superAdminCredentials.ID,
		"token": token,
	})
}

func (uc *AccountUC) AdminSignin(ctx context.Context, req *request.SigninReq) *response.Response {
	adminCredentials, err := uc.repo.GetCredential(ctx, req.Username, constants.RoleAdmin)
	if err != nil {
		if err == repo.ErrRecordNotFound {
			return invalidUsernameResponse
		}
		return response.DBErrorResponse(err)
	}

	//compare password
	if err := comparePassword(adminCredentials.HashedPassword, req.Password); err != nil {
		return invalidPasswordResponse
	}

	//generate token
	token, err := jwttoken.GenerateToken(adminCredentials.ID, constants.RoleSuperAdmin, "", nil, defaultTokenExpiry)
	if err != nil {
		return response.ErrorResponse(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"id":    adminCredentials.ID,
		"token": token,
	})
}

func (uc *AccountUC) UserSignin(ctx context.Context, req *request.SigninReq) *response.Response {
	userCredentials, err := uc.repo.GetCredential(ctx, req.Username, constants.RoleUser)
	if err != nil {
		if err == repo.ErrRecordNotFound {
			return invalidUsernameResponse
		}
		return response.DBErrorResponse(err)
	}

	//compare password
	if err := comparePassword(userCredentials.HashedPassword, req.Password); err != nil {
		return invalidPasswordResponse
	}

	//generate token
	token, err := jwttoken.GenerateToken(userCredentials.ID, constants.RoleSuperAdmin, "", nil, defaultTokenExpiry)
	if err != nil {
		return response.ErrorResponse(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"id":    userCredentials.ID,
		"token": token,
	})
}

func (uc *AccountUC) UserSignUpGetOTP(ctx context.Context, req *request.UserSignupReq) *response.Response {

	hashedPassword, err := hashpassword.GetHashedPassword(req.Password)
	if err != nil {
		return response.ErrorResponse(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in hashing password: %v", err))
	}

	user := models.User{
		Username:       req.Username,
		Email:          req.Email,
		Name:           req.Name,
		Phone:          req.Phone,
		HashedPassword: hashedPassword,
		IsBlocked:      false,
		IsVerified:     false,
	}
	if err := uc.repo.CreateRecord(ctx, &user); err != nil {
		return response.DBErrorResponse(err)
	}

	//generate token
	token, err := jwttoken.GenerateToken(user.ID, constants.UnverifiedUser, "", nil, time.Minute*5)
	if err != nil {
		return response.InternalServerErrorResponse(fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"id":              user.ID,
		"temporary_token": token,
		"name":            user.Name,
	})
}

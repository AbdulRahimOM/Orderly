package accountuc

import (
	"context"
	"fmt"
	"net/http"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	"orderly/internal/infrastructure/config"
	repo "orderly/internal/repository"
	jwttoken "orderly/pkg/jwt-token"

	"golang.org/x/crypto/bcrypt"
)

const (
	USERNAME_NOT_REGISTERED = "USERNAME_NOT_REGISTERED"
	PASSWORD_MISMATCH       = "PASSWORD_MISMATCH"
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

func (uc *AccountUC) SuperAdminSignin(ctx context.Context, req *request.SuperAdminSigninReq) *response.Response {
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
	token, err := jwttoken.GenerateToken(superAdminCredentials.ID, constants.RoleSuperAdmin, "", nil)
	if err != nil {
		return response.ErrorResponse(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"id":    superAdminCredentials.ID,
		"token": token,
	})
}

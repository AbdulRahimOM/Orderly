package uc

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

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	USERNAME_NOT_REGISTERED = "USERNAME_NOT_REGISTERED"
	PASSWORD_MISMATCH       = "PASSWORD_MISMATCH"
	defaultTokenExpiry      = constants.DefaultTokenExpiry
)

type SignupUser struct {
	Username  string    `gorm:"column:username;unique" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"column:email" json:"email"`
	Name      string    `gorm:"column:name" json:"name"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
}

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

func (uc *Usecase) SuperAdminSignin(ctx context.Context, req *request.SigninReq) *response.Response {
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
	token, err := jwttoken.GenerateToken(superAdminCredentials.ID, constants.RoleSuperAdmin, nil, defaultTokenExpiry)
	if err != nil {
		return response.ErrorResponse(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"id":    superAdminCredentials.ID,
		"token": token,
	})
}

func (uc *Usecase) AdminSignin(ctx context.Context, req *request.SigninReq) *response.Response {
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

	privileges, err := uc.repo.GetAccessPrivilegeByAdminID(ctx, adminCredentials.ID.String())
	if err != nil {
		return response.DBErrorResponse(err)
	}

	//generate token
	token, err := jwttoken.GenerateToken(
		adminCredentials.ID,
		constants.RoleAdmin,
		map[string]interface{}{
			constants.Privilege: privileges.AccessRoles,
		},
		defaultTokenExpiry,
	)
	if err != nil {
		return response.ErrorResponse(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"id":         adminCredentials.ID, //dev purpose
		"token":      token,
		"privileges": privileges.AccessRoles,
	})
}

func (uc *Usecase) UserSignin(ctx context.Context, req *request.SigninReq) *response.Response {
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
	token, err := jwttoken.GenerateToken(userCredentials.ID, constants.RoleUser, nil, defaultTokenExpiry)
	if err != nil {
		return response.ErrorResponse(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"id":    userCredentials.ID, //dev purpose
		"token": token,
	})
}

func (uc *Usecase) UserSignUpGetOTP(ctx context.Context, req *request.UserSignupReq) *response.Response {
	//check if username, email, phone already exists
	usernameExists, emailExists, phoneExists, err := uc.repo.CheckIfUsernameEmailOrPhoneExistsInUser(ctx, req.Username, req.Email, req.Phone)
	switch {
	case err != nil:
		return response.DBErrorResponse(fmt.Errorf("error in checking if username exists: %v", err))
	case usernameExists:
		return response.ErrorResponse(http.StatusConflict, respcode.USERNAME_ALREADY_EXISTS, fmt.Errorf("username already exists"))
	case emailExists:
		return response.ErrorResponse(http.StatusConflict, respcode.EMAIL_ALREADY_EXISTS, fmt.Errorf("email already exists"))
	case phoneExists:
		return response.ErrorResponse(http.StatusConflict, respcode.PHONE_ALREADY_EXISTS, fmt.Errorf("phone already exists"))
	}

	user := SignupUser{
		Username:  req.Username,
		Email:     req.Email,
		Name:      req.Name,
		Phone:     req.Phone,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}

	//generate token
	token, err := jwttoken.GenerateToken(uuid.Nil, constants.UnverifiedUser, map[string]interface{}{
		"user": user,
	}, time.Minute*5)
	if err != nil {
		return response.InternalServerErrorResponse(fmt.Errorf("error in generating token: %v", err))
	}

	//send OTP via SMS
	err = uc.smsOtpClient.SendOtp(user.Phone)
	if err != nil {
		return response.InternalServerErrorResponse(fmt.Errorf("error in sending OTP: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"temporary_token": token,
		"name":            user.Name,
	})
}

func (uc *Usecase) UserSignUpVerifyOTP(ctx context.Context, req *request.VerifyOTPReq) *response.Response {
	userInToken, _ := ctx.Value("user").(map[string]interface{})

	user := models.User{}
	user.Username, _ = userInToken["username"].(string)
	user.Email, _ = userInToken["email"].(string)
	user.Name, _ = userInToken["name"].(string)
	user.Phone, _ = userInToken["phone"].(string)
	password, _ := userInToken["password"].(string)

	//check if username, email, phone already exists
	usernameExists, emailExists, phoneExists, err := uc.repo.CheckIfUsernameEmailOrPhoneExistsInUser(ctx, user.Username, user.Email, user.Phone)
	switch {
	case err != nil:
		return response.DBErrorResponse(fmt.Errorf("error in checking if username exists: %v", err))
	case usernameExists:
		return response.ErrorResponse(http.StatusConflict, respcode.USERNAME_ALREADY_EXISTS, fmt.Errorf("username already exists"))
	case emailExists:
		return response.ErrorResponse(http.StatusConflict, respcode.EMAIL_ALREADY_EXISTS, fmt.Errorf("email already exists"))
	case phoneExists:
		return response.ErrorResponse(http.StatusConflict, respcode.PHONE_ALREADY_EXISTS, fmt.Errorf("phone already exists"))
	}

	// verify otp
	if ok, err := uc.smsOtpClient.VerifyOtp(user.Phone, req.OTP); err != nil {
		return response.InternalServerErrorResponse(fmt.Errorf("error in verifying OTP: %v", err))
	} else if !ok {
		return response.ErrorResponse(http.StatusUnauthorized, respcode.INVALID_OTP, fmt.Errorf("invalid OTP"))
	}

	hashpassword, err := hashpassword.GetHashedPassword(password)
	if err != nil {
		return response.InternalServerErrorResponse(fmt.Errorf("error in hashing password: %v", err))
	} else {
		user.HashedPassword = hashpassword
	}

	//create user
	user.ID = uuid.New()
	err = uc.repo.CreateRecord(ctx, &user)
	if err != nil {
		return response.DBErrorResponse(fmt.Errorf("error in creating user: %v", err))
	}

	//generate token
	token, err := jwttoken.GenerateToken(user.ID, constants.RoleUser, nil, defaultTokenExpiry)
	if err != nil {
		return response.InternalServerErrorResponse(fmt.Errorf("error in generating token: %v", err))
	}

	return response.SuccessResponse(http.StatusOK, respcode.Success, map[string]interface{}{
		"token": token,
		"name":  user.Name,
		"id":    user.ID, //for dev purpose
	})
}

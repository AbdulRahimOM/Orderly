package repo

import (
	"context"
	"fmt"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/dto"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"

	"github.com/gofiber/fiber/v2/log"
)

var (
	ErrRecordNotFound = fmt.Errorf("record not found")
)

func (r *Repo) GetCredential(ctx context.Context, username string, role string) (*dto.Credentials, error) {
	var tableName string
	switch role {
	case constants.RoleSuperAdmin:
		tableName = models.SuperAdmin_TableName
	case "admin":
		tableName = models.Admins_TableName
	case "user":
		tableName = models.Users_TableName
	default:
		log.Error("potential bug: invalid role mentioned in code: %s", role)
		return nil, fmt.Errorf("potential bug: Invalid role mentioned in code: %s", role)
	}

	var credentials dto.Credentials
	result := r.db.Table(tableName).Where("username = ?", username).Find(&credentials)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting credentials: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, ErrRecordNotFound
	}
	return &credentials, nil
}

func (r *Repo) GetAdmins(ctx context.Context, req *request.GetRequest) ([]dto.AdminInList, error) {
	var (
		admins []dto.AdminInList
		err    error
	)

	if req.IsDeleted {
		err = r.db.Table(models.Admins_TableName).Select("id", "name", "phone", "designation", "is_blocked").Where("deleted_at IS NOT NULL").Scan(&admins).Limit(req.Limit).Offset(req.Offset).Error
	} else {
		err = r.db.Table(models.Admins_TableName).Select("id", "name", "phone", "designation", "is_blocked").Where("deleted_at IS NULL").Scan(&admins).Limit(req.Limit).Offset(req.Offset).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error getting admins: %v", err)
	}
	return admins, nil
}

func (r *Repo) GetAdminByID(ctx context.Context, id string) (*dto.Admin, error) {
	var admin dto.Admin
	result := r.db.Raw(`
		SELECT 
			id, name, email, phone, designation, is_blocked, created_at, updated_at, deleted_at, 
			CASE WHEN deleted_at IS NULL THEN false ELSE true END as is_deleted
		FROM admins
		WHERE id = ?
	`, id).Scan(&admin)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting admin: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, ErrRecordNotFound
	}
	return &admin, nil
}

func (r *Repo) UpdateAdminByID(ctx context.Context, id string, req *request.UpdateAdminReq) error {
	result := r.db.Table(models.Admins_TableName).Where("id = ?", id).Save(req)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (r *Repo) MarkUserAsVerified(ctx context.Context, id string) error {
	result := r.db.Table(models.Users_TableName).Where("id = ? AND is_blocked = false AND deleted_at IS NULL", id).Update("is_verified", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (r *Repo) GetUserSignInDetails(ctx context.Context, userID string) (*dto.UserSignInDetails, error) {
	var user dto.UserSignInDetails
	result := r.db.Table(models.Users_TableName).Select("name,phone,is_blocked").Where("id = ? AND deleted_at IS NULL", userID).Scan(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting user details: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, ErrRecordNotFound
	}
	return &user, nil
}

//CheckIfUsernameEmailOrPhoneExists(ctx context.Context, username, email, phone string) (usernameExists, emailExists, phoneExists bool, err error)
func (r *Repo) CheckIfUsernameEmailOrPhoneExistsInUser(ctx context.Context, username, email, phone string) (usernameExists, emailExists, phoneExists bool, err error) {
	var count int64
	result := r.db.Table(models.Users_TableName).Where("username = ?", username).Count(&count)
	if result.Error != nil {
		return false, false, false, fmt.Errorf("error checking username: %v", result.Error)
	}
	usernameExists = count > 0

	result = r.db.Table(models.Users_TableName).Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return false, false, false, fmt.Errorf("error checking email: %v", result.Error)
	}
	emailExists = count > 0

	result = r.db.Table(models.Users_TableName).Where("phone = ?", phone).Count(&count)
	if result.Error != nil {
		return false, false, false, fmt.Errorf("error checking phone: %v", result.Error)
	}
	phoneExists = count > 0

	return usernameExists, emailExists, phoneExists, nil
}
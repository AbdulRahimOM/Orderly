package repo

import (
	"context"
	"fmt"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/dto"
	"orderly/internal/domain/models"
)

var ErrRecordNotFound = fmt.Errorf("record not found")

func (r *Repo) GetCredential(ctx context.Context, username string, role string) (*dto.Credentials, error) {
	var tableName string
	switch role {
	case constants.RoleSuperAdmin:
		tableName = models.SuperAdmin_TableName
		// case "admin":
		// 	tableName = models.Admin_TableName
		// case "user":
		// 	tableName = models.User_TableName
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

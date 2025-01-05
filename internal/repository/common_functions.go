package repo

import (
	"context"
	"fmt"
	"orderly/internal/domain/respcode"
	"strings"
)

func (r *Repo) CreateRecord(ctx context.Context, record interface{}) error {
	if err := r.db.WithContext(ctx).Create(record).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) SoftDeleteRecordByID(ctx context.Context, tableName string, id int) error {
	if err := r.db.WithContext(ctx).Table(tableName).Where("id = ?", id).Update("deleted_at", "now()").Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UndoSoftDeleteRecordByID(ctx context.Context, tableName string, id int) (string, error) {
	err := r.db.Table(tableName).Where("id = ?", id).Update("deleted_at", nil).Error
	if err != nil {
		if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
			return respcode.UniqueFieldViolation, err
		}
		return respcode.DbError, fmt.Errorf("failed to undo soft delete record in table %s: %w", tableName, err)
	}

	return "", nil
}

func (r *Repo) SoftDeleteRecordByUUID(ctx context.Context, tableName string, id string) error {
	result := r.db.WithContext(ctx).Table(tableName).Where("id = ?", id).Update("deleted_at", "now()")
	if result.Error != nil {
		return fmt.Errorf("failed to soft delete record in table %s: %w", tableName, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found in table %s", tableName)
	}
	return nil
}

func (r *Repo) UndoSoftDeleteRecordByUUID(ctx context.Context, tableName string, id string) (string, error) {
	result := r.db.Table(tableName).Where("id = ?", id).Update("deleted_at", nil)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "(SQLSTATE 23505)") {
			return respcode.UniqueFieldViolation, result.Error
		}
		return respcode.DbError, fmt.Errorf("failed to undo soft delete record in table %s: %w", tableName, result.Error)
	}
	if result.RowsAffected == 0 {
		return respcode.DbError, fmt.Errorf("record not found in table %s", tableName)
	}

	return "", nil
}

func (r *Repo) ActivateByUUID(ctx context.Context, tableName string, id string) error {
	result :=  r.db.WithContext(ctx).Table(tableName).Where("id = ?", id).Update("is_active", true)
	if result.Error != nil {
		return fmt.Errorf("failed to activate record in table %s: %w", tableName, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found in table %s", tableName)
	}
	return nil
}

func (r *Repo) DeactivateByUUID(ctx context.Context, tableName string, id string) error {
	result := r.db.WithContext(ctx).Table(tableName).Where("id = ?", id).Update("is_active", false)
	if result.Error != nil {
		return fmt.Errorf("failed to deactivate record in table %s: %w", tableName, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found in table %s", tableName)
	}
	return nil
}

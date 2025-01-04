package repo

import "context"



func (r *Repo) CreateRecord(ctx context.Context, record interface{}) error {
	if err := r.db.WithContext(ctx).Create(record).Error; err != nil {
		return err
	}
	return nil
}

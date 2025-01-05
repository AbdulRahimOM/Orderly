package repo

import (
	"context"
	"fmt"
	"orderly/internal/domain/dto"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
)

func (r *Repo) GetCategories(ctx context.Context, req *request.GetRequest) ([]models.Category, error) {
	var (
		categories     []models.Category
		whereCondition string
	)
	if req.IsDeleted {
		whereCondition = "deleted_at IS NOT NULL"
	} else {
		whereCondition = "deleted_at IS NULL"
	}
	err := r.db.WithContext(ctx).Unscoped().Where(whereCondition).Order("name").Limit(req.Limit).Offset(req.Offset).Find(&categories).Error
	if err != nil {
		return nil, fmt.Errorf("error getting categories: %v", err)
	}

	return categories, nil
}

func (r *Repo) GetCategoryByID(ctx context.Context, id int) (*models.Category, error) {
	var category models.Category
	result := r.db.WithContext(ctx).Where("id = ?", id).Find(&category)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting category: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, ErrRecordNotFound
	}

	return &category, nil
}

func (r *Repo) UpdateCategoryByID(ctx context.Context, id int, req *request.CategoryReq) error {
	err := r.db.WithContext(ctx).Model(&models.Category{}).Where("id = ?", id).Update("name", req.Name).Error
	if err != nil {
		return fmt.Errorf("error updating category: %v", err)
	}

	return nil
}

func (r *Repo) GetProducts(ctx context.Context, req *request.GetRequest) ([]dto.ProductInList, error) {
	var (
		products         []dto.ProductInList
		deletedCondition string
	)
	if req.IsDeleted {
		deletedCondition = "deleted_at IS NOT NULL"
	} else {
		deletedCondition = "deleted_at IS NULL"
	}
	err := r.db.Table(models.Products_TableName).Select("id", "name", "description", "current_sale_price", "max_sale_price", "current_stock").Where(deletedCondition).Order("name").Limit(req.Limit).Offset(req.Offset).Scan(&products).Error
	if err != nil {
		return nil, fmt.Errorf("error getting products: %v", err)
	}

	return products, nil
}

func (r *Repo) GetProductByID(ctx context.Context, id int) (*dto.Product, error) {
	var product dto.Product
	result := r.db.Raw(`
		SELECT 
			p.id, p.name, description, category_id, min_sale_price, max_sale_price, base_price, current_sale_price, optimal_stock, current_stock, created_at, updated_at, p.deleted_at,
			CASE WHEN p.deleted_at IS NULL THEN false ELSE true END as is_deleted,
			c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = ?
	`, id).Scan(&product)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting product: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, ErrRecordNotFound
	}

	return &product, nil
}

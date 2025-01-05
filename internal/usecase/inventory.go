package uc

import (
	"context"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
	"orderly/internal/domain/respcode"
	"orderly/internal/domain/response"
	repo "orderly/internal/repository"
)

func (u *Usecase) CreateCategory(ctx context.Context, req *request.CategoryReq) *response.Response {
	category := &models.Category{
		Name: req.Name,
	}

	err := u.repo.CreateRecord(ctx, category)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.CreatedResponse(category.ID)
}

func (u *Usecase) GetCategories(ctx context.Context, req *request.GetRequest) *response.Response {
	categories, err := u.repo.GetCategories(ctx, req)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, categories)
}

func (u *Usecase) GetCategoryByID(ctx context.Context, id int) *response.Response {
	category, err := u.repo.GetCategoryByID(ctx, id)
	if err != nil {
		if err == repo.ErrRecordNotFound {
			return response.NotFoundResponse("category")
		}
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, category)
}

func (u *Usecase) UpdateCategoryByID(ctx context.Context, id int, req *request.CategoryReq) *response.Response {
	err := u.repo.UpdateCategoryByID(ctx, id, req)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, nil)
}

func (u *Usecase) CreateProduct(ctx context.Context, req *request.ProductReq) *response.Response {
	product := &models.Product{
		Name:             req.Name,
		Description:      req.Description,
		CategoryID:       req.CategoryID,
		MinSalePrice:     req.MinSalePrice,
		MaxSalePrice:     req.MaxSalePrice,
		BasePrice:        req.BasePrice,
		CurrentSalePrice: req.CurrentSalePrice,
		OptimalStock:     req.OptimalStock,
		CurrentStock:     req.CurrentStock,
	}

	err := u.repo.CreateRecord(ctx, product)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.CreatedResponse(product.ID)
}

func (u *Usecase) GetProducts(ctx context.Context, req *request.GetRequest) *response.Response {
	products, err := u.repo.GetProducts(ctx, req)
	if err != nil {
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, products)
}

func (u *Usecase) GetProductByID(ctx context.Context, id int) *response.Response {
	product, err := u.repo.GetProductByID(ctx, id)
	if err != nil {
		if err == repo.ErrRecordNotFound {
			return response.NotFoundResponse("product")
		}
		return response.DBErrorResponse(err)
	}

	return response.SuccessResponse(200, respcode.Success, product)
}

func (u *Usecase) UpdateProductByID(ctx context.Context, id int, req *request.ProductReq) *response.Response {
	// err := u.repo.UpdateProductByID(ctx, id, req)
	// if err != nil {
	// 	return response.DBErrorResponse(err)
	// }

	return response.SuccessResponse(200, respcode.Success, nil)
}
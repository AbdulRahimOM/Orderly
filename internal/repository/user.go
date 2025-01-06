package repo

import (
	"context"
	"fmt"
	"orderly/internal/domain/dto"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
	"orderly/pkg/utils/helper"

	"gorm.io/gorm"
)

func (r *Repo) GetCart(ctx context.Context) ([]dto.Cart, error) {
	var (
		userID = helper.GetUserIdFromContext(ctx)
		cart   = []dto.Cart{}
	)
	err := r.db.WithContext(ctx).Table(models.CartItems_TableName).
		Select("product_id, quantity, price_when_put_in_cart","current_sale_price").
		Joins("JOIN products ON cart_items.product_id = products.id").
		Where("user_id = ?", userID).
		Scan(&cart).Error
	if err != nil {
		return nil, fmt.Errorf("error getting cart: %v", err)
	}

	return cart, nil
}

func (r *Repo) AddToCart(ctx context.Context, req *request.AddToCartReq) error {
	var (
		userID = helper.GetUserIdFromContext(ctx)
		alreadyInCart bool
	)

	//check if product is already in cart
	err := r.db.WithContext(ctx).Raw("SELECT EXISTS(SELECT 1 FROM cart_items WHERE user_id = ? AND product_id = ?)", userID, req.ProductID).Scan(&alreadyInCart).Error
	if err != nil {
		return fmt.Errorf("error checking if product already in cart: %v", err)
	}

	fmt.Println("alreadyInCart: ", alreadyInCart)

	if alreadyInCart {
		//update quantity
		err = r.db.WithContext(ctx).Table(models.CartItems_TableName).
			Where("user_id = ? AND product_id = ?", userID, req.ProductID).
			Update("quantity", gorm.Expr("quantity + ?", 1)).Error
		if err != nil {
			return fmt.Errorf("error updating quantity in cart: %v", err)
		}
	} else {
		//add to cart
		err = r.db.WithContext(ctx).Exec(`
			INSERT INTO cart_items (user_id, product_id, quantity, price_when_put_in_cart)
			VALUES ($1, $2, 1, (SELECT current_sale_price FROM products WHERE id = $2))
			`, userID, req.ProductID).Error
		if err != nil {
			return fmt.Errorf("error adding to cart: %v", err)
		}
	}

	return nil
}

func (r *Repo) RemoveProductFromCart(ctx context.Context, productId int) error {
	var (
		userID = helper.GetUserIdFromContext(ctx)
	)

	err := r.db.WithContext(ctx).Table(models.CartItems_TableName).
		Where("user_id = ? AND product_id = ?", userID, productId).
		Delete(nil).Error
	if err != nil {
		return fmt.Errorf("error removing product from cart: %v", err)
	}

	return nil
}

func (r *Repo) UpdateCartItemQuantity(ctx context.Context, req *request.UpdateCartItemQuantityReq) error {
	var (
		userID = helper.GetUserIdFromContext(ctx)
	)

	err := r.db.WithContext(ctx).Table(models.CartItems_TableName).
		Where("user_id = ? AND product_id = ?", userID, req.ProductID).
		Update("quantity", req.Quantity).Error
	if err != nil {
		return fmt.Errorf("error updating quantity in cart: %v", err)
	}

	return nil
}

func (r *Repo) ClearCart(ctx context.Context) error {
	var (
		userID = helper.GetUserIdFromContext(ctx)
	)

	err := r.db.WithContext(ctx).Table(models.CartItems_TableName).
		Where("user_id = ?", userID).
		Delete(nil).Error
	if err != nil {
		return fmt.Errorf("error clearing cart: %v", err)
	}

	return nil
}
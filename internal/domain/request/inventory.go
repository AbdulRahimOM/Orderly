package request

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CategoryReq struct {
	Name string `json:"name" validate:"required"`
}

type ProductReq struct {
	Name             string  `json:"name" validate:"required"`
	Description      string  `json:"description" validate:"required,max=255"`
	CategoryID       int     `json:"category_id" validate:"required"`
	MinSalePrice     float64 `json:"min_sale_price" validate:"required,gt=0"`
	MaxSalePrice     float64 `json:"max_sale_price" validate:"required,gt=0"`
	BasePrice        float64 `json:"base_price" validate:"required,gt=0"`         //foundational price that serves as a reference for adjustments based on market conditions.
	CurrentSalePrice float64 `json:"current_sale_price" validate:"required,gt=0"` //the price at which the product is currently being sold.
	OptimalStock     int     `json:"optimal_stock" validate:"required,gte=1"`     //the ideal amount of stock to have on hand.
	CurrentStock     int     `json:"current_stock" validate:"required,gte=0"`     //the amount of stock currently on hand.
}

func (p *ProductReq) CustomValidation(ctx *fiber.Ctx) (string, error) {
	const invalidPriceValues = "INVALID_PRICE_VALUES"
	switch {
	case p.MinSalePrice > p.MaxSalePrice:
		return invalidPriceValues, fmt.Errorf("min_sale_price cannot be greater than max_sale_price")
	case p.MinSalePrice > p.BasePrice:
		return invalidPriceValues, fmt.Errorf("min_sale_price cannot be greater than base_price")
	case p.MinSalePrice > p.CurrentSalePrice:
		return invalidPriceValues, fmt.Errorf("min_sale_price cannot be greater than current_sale_price")
	case p.MaxSalePrice < p.BasePrice:
		return invalidPriceValues, fmt.Errorf("max_sale_price cannot be less than base_price")
	case p.MaxSalePrice < p.CurrentSalePrice:
		return invalidPriceValues, fmt.Errorf("max_sale_price cannot be less than current_sale_price")
	default:
		return "", nil
	}
}

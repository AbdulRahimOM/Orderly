package request

type AddToCartReq struct {
	ProductID int `json:"product_id" validate:"required"`
}

type UpdateCartItemQuantityReq struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type RemoveProductFromCartReq struct {
	ProductID int `json:"product_id" validate:"required"`
}

type CreateOrderReq struct {
	PaymentMethod string `json:"payment_method"`
	AddressID string `json:"address_id" validate:"required"`
}

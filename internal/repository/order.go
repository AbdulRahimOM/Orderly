package repo

import (
	"context"
	"fmt"
	"log"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/dto"
	"orderly/internal/domain/models"
	"orderly/internal/domain/request"
	"orderly/pkg/utils/helper"
	"time"

	"gorm.io/gorm"
)

const (
	INSUFFICIENT_STOCK = "INSUFFICIENT_STOCK"
)

func (r *Repo) GetOrders(ctx context.Context, req *request.Pagination) ([]*dto.OrderInListForAdmin, error) {
	var (
		orders []*dto.OrderInListForAdmin
	)

	err := r.db.WithContext(ctx).Raw(`
		SELECT
			o.id,
			o.user_id,
			u.name AS customer_name,
			o.order_time,
			o.total_amount,
			o.payment_method,
			o.order_status
		FROM
			orders o
		JOIN
			users u ON o.user_id = u.id
		ORDER BY
			o.order_time DESC
		LIMIT ?
		OFFSET ?
	`, req.Limit, req.Offset).Scan(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("error getting orders: %v", err)
	}

	return orders, nil
}

func (r *Repo) GetMyOrders(ctx context.Context, req *request.Pagination) ([]*dto.OrderInListForUser, error) {
	var (
		userID = helper.GetUserIdFromContext(ctx)
		orders []*dto.OrderInListForUser
	)

	err := r.db.WithContext(ctx).Raw(`
		SELECT
			o.id,
			o.order_time,
			o.total_amount,
			o.payment_method,
			o.order_status
		FROM
			orders o
		WHERE
			o.user_id = ?
		ORDER BY
			o.order_time DESC
		LIMIT ?
		OFFSET ?
	`, userID, req.Limit, req.Offset).Scan(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("error getting orders: %v", err)
	}

	return orders, nil
}

func (r *Repo) GetOrderDetails(ctx context.Context, id string) (*dto.OrderForAdmin, error) {
	var order dto.OrderForAdmin
	err := r.db.WithContext(ctx).Raw(`
		SELECT
			o.id,
			o.user_id,
			u.name AS customer_name,
			o.order_time,
			o.total_amount,
			o.payment_method,
			o.payment_id,
			o.order_status,
			CASE
				WHEN o.delivered_or_cancelled_at IS NOT NULL AND o.order_status = 'delivered' THEN o.delivered_or_cancelled_at
				ELSE NULL
			END AS delivered_at,
			CASE WHEN o.delivered_or_cancelled_at IS NOT NULL AND o.order_status = 'cancelled' THEN o.delivered_or_cancelled_at
				ELSE NULL
			END AS cancelled_at,

			-- shipping address
			o.house,
			o.street1,
			o.street2,
			o.city,
			o.state,
			o.pincode,
			o.landmark,
			o.country
		FROM
			orders o
		JOIN
			users u ON o.user_id = u.id
		WHERE
			o.id = ?
	`, id).Scan(&order).Error
	if err != nil {
		return nil, fmt.Errorf("error getting order details: %v", err)
	}

	return &order, nil
}

func (r *Repo) CancelOrder(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Table(models.Orders_TableName).Where("id = ?", id).Update("order_status", constants.Order_Cancelled).Error
	if err != nil {
		return fmt.Errorf("error cancelling order: %v", err)
	}

	return nil
}

func (r *Repo) GetCartItemsForOrder(ctx context.Context) ([]*dto.CartItemsForOrder, error) {
	var (
		userID    = helper.GetUserIdFromContext(ctx)
		cartItems []*dto.CartItemsForOrder
	)

	err := r.db.WithContext(ctx).Raw(`
		SELECT
			c.product_id,
			c.quantity,
			p.current_sale_price,
			p.current_stock
		FROM
			cart_items c
		JOIN
			products p ON c.product_id = p.id
		WHERE
			c.user_id = ?
	`, userID).Scan(&cartItems).Error
	if err != nil {
		return nil, fmt.Errorf("error getting cart items: %v", err)
	}

	return cartItems, nil
}

// CreateOrderFromCart(ctx context.Context, req *request.CreateOrderReq) (orderID int, responsecode string, err error)
func (r *Repo) CreateOrder(ctx context.Context, order *models.Order, orderProducts []*models.OrderProduct) (int error) {

	var (
		err error
	)

	tx := r.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback().Error; err != nil && err != gorm.ErrInvalidTransaction {
				log.Printf("rollback failed: %v\n", err)
			}
		}
	}()

	err = tx.WithContext(ctx).Create(order).Error
	if err != nil {
		return fmt.Errorf("error creating order: %v", err)
	}

	for i := range orderProducts {
		orderProducts[i].OrderID = order.ID
		err = tx.WithContext(ctx).Create(orderProducts[i]).Error
		if err != nil {
			return fmt.Errorf("error creating order product: %v", err)
		}
	}

	// delete cart items
	err = tx.WithContext(ctx).Exec("DELETE FROM cart_items WHERE user_id = ?", order.UserID).Error
	if err != nil {
		return fmt.Errorf("error deleting cart items: %v", err)
	}

	//reduce stock
	for _, item := range orderProducts {
		err = tx.WithContext(ctx).Exec(`
			UPDATE products
			SET current_stock = current_stock - ?
			WHERE id = ?
		`, item.Quantity, item.ProductID).Error
		if err != nil {
			return fmt.Errorf("error reducing stock: %v", err)
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (r *Repo) CheckOrderBelongsToUser(ctx context.Context, orderID string) (bool, error) {
	var (
		userID        = helper.GetUserIdFromContext(ctx)
		belongsToUser bool
	)

	err := r.db.WithContext(ctx).
		Raw("SELECT EXISTS(SELECT 1 FROM orders WHERE id = ? AND user_id = ?)", orderID, userID).
		Scan(&belongsToUser).Error
	if err != nil {
		return false, fmt.Errorf("error checking if order belongs to user: %v", err)
	}

	return belongsToUser, nil
}

func (r *Repo) GetOrderStatus(ctx context.Context, orderID string) (string, bool, error) {
	var output struct {
		OrderStatus string `json:"order_status" gorm:"column:order_status"`
		PaymentDone bool   `json:"payment_done" gorm:"column:payment_done"`
	}
	result := r.db.WithContext(ctx).Raw(`
		SELECT
			order_status,
			CASE WHEN payment_id IS NOT NULL THEN true ELSE false END AS payment_done
		FROM orders
		WHERE id = ?
	`, orderID).Scan(&output)
	if result.Error != nil {
		return "", false, fmt.Errorf("error getting order status: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return "", false, ErrRecordNotFound
	}

	return output.OrderStatus, output.PaymentDone, nil
}

func (r *Repo) MarkOrderAsDelivered(ctx context.Context, orderID string) error {
	err := r.db.WithContext(ctx).Model(models.Order{}).Where("id = ?", orderID).Updates(map[string]interface{}{
		"order_status":              constants.Order_Delivered,
		"delivered_or_cancelled_at": time.Now(),
	}).Error
	if err != nil {
		return fmt.Errorf("error marking order as delivered: %v", err)
	}

	return nil
}

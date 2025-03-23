package order

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateOrderItem(orderItem *model.OrderItem) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO order_items (order_details_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING order_item_id",
		orderItem.OrderDetailsID, orderItem.ProductID, orderItem.Quantity,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateOrderItem: %v", err)
	}

	return lastInsertedID, nil
}

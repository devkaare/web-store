package order

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetOrderItemsByOrderDetailsID(orderDetailsID int) ([]model.OrderItem, error) {
	var orderItems []model.OrderItem

	rows, err := r.Client.Query("SELECT order_item_id, user_id, quantity FROM order_items WHERE order_details_id = $1", orderDetailsID)
	if err != nil {
		return orderItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem model.OrderItem
		if err := rows.Scan(&orderItem.OrderItemID, &orderItem.OrderDetailsID, &orderItem.ProductID, &orderItem.Quantity); err != nil {
			return orderItems, fmt.Errorf("GetOrderItemsByOrderDetailsID %d: %v", orderItem.OrderItemID, err)
		}
		orderItems = append(orderItems, orderItem)
	}
	if err := rows.Err(); err != nil {
		return orderItems, fmt.Errorf("GetOrderItemsByOrderDetailsID %v:", err)
	}
	return orderItems, nil
}

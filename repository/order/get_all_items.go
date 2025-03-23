package order

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetAllOrderItems() ([]model.OrderItem, error) {
	var allOrderItems []model.OrderItem

	rows, err := r.Client.Query("SELECT * FROM order_items")
	if err != nil {
		return allOrderItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem model.OrderItem
		if err := rows.Scan(&orderItem.OrderItemID, &orderItem.OrderDetailsID, &orderItem.ProductID, &orderItem.Quantity); err != nil {
			return allOrderItems, fmt.Errorf("GetAllOrderItems %d: %v", orderItem.OrderItemID, err)
		}
		allOrderItems = append(allOrderItems, orderItem)
	}
	if err := rows.Err(); err != nil {
		return allOrderItems, fmt.Errorf("GetAllOrderItems %v:", err)
	}
	return allOrderItems, nil
}

package order

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetOrderItemByOrderItemID(orderItemID int) (*model.OrderItem, error) {
	orderItem := &model.OrderItem{}

	row := r.Client.QueryRow("SELECT * FROM order_items WHERE order_item_id = $1", orderItemID)
	if err := row.Scan(&orderItem.OrderItemID, &orderItem.OrderDetailsID, &orderItem.ProductID, &orderItem.Quantity); err != nil {
		if err == sql.ErrNoRows {
			return orderItem, err
		}
		return orderItem, fmt.Errorf("GetOrderItemByOrderItemID %d: %v", orderItemID, err)
	}
	return orderItem, nil
}

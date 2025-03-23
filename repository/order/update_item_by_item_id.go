package order

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateOrderItemByOrderItemID(orderItem *model.OrderItem) error {
	_, err := r.Client.Exec("UPDATE order_items SET order_details_id = $2, product_id = $3, quantity = $4 WHERE order_item_id = $1", orderItem.OrderItemID, orderItem.OrderDetailsID, orderItem.ProductID, orderItem.Quantity)
	if err != nil {
		return fmt.Errorf("UpdateOrderItemByOrderItemID: %v", err)
	}
	return nil
}

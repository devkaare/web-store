package order

import (
	"fmt"
)

func (r *Repo) DeleteOrderItemByOrderItemID(orderItemID int) error {
	result, err := r.Client.Exec("DELETE FROM order_items WHERE order_item_id = $1", orderItemID)
	if err != nil {
		return fmt.Errorf("DeleteOrderItemByOrderItemID %d: %v", orderItemID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteOrderItemByOrderItemID %d: %v", orderItemID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteOrderItemByOrderItemID %d: order_item not found", orderItemID)
	}
	return nil
}

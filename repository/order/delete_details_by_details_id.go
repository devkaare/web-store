package order

import (
	"fmt"
)

func (r *Repo) DeleteOrderDetailsByOrderDetailsID(orderDetailsID int) error {
	result, err := r.Client.Exec("DELETE FROM order_details WHERE order_details_id = $1", orderDetailsID)
	if err != nil {
		return fmt.Errorf("DeleteOrderDetailsByOrderDetailsID %d: %v", orderDetailsID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteOrderDetailsByOrderDetailsID %d: %v", orderDetailsID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteOrderDetailsByOrderDetailsID %d: order_details not found", orderDetailsID)
	}
	return nil
}

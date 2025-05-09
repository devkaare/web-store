package order

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateOrderDetailsByOrderDetailsID(orderDetails *model.OrderDetails) error {
	_, err := r.Client.Exec("UPDATE order_details SET user_id = $2, payment_id = $3, total = $4, payment_status = $5 WHERE order_details_id = $1", orderDetails.OrderDetailsID, orderDetails.UserID, orderDetails.PaymentID, orderDetails.Total, orderDetails.PaymentStatus)
	if err != nil {
		return fmt.Errorf("UpdateOrderDetailsByOrderDetailsID: %v", err)
	}
	return nil
}

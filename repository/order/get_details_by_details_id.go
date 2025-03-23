package order

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetOrderDetailsByOrderDetailsID(orderDetailsID int) (*model.OrderDetails, error) {
	orderDetails := &model.OrderDetails{}

	row := r.Client.QueryRow("SELECT * FROM order_details WHERE order_details_id = $1", orderDetailsID)
	if err := row.Scan(&orderDetails.OrderDetailsID, &orderDetails.UserID, &orderDetails.PaymentID, &orderDetails.Total); err != nil {
		if err == sql.ErrNoRows {
			return orderDetails, err
		}
		return orderDetails, fmt.Errorf("GetOrderDetailsByOrderDetailsID %d: %v", orderDetailsID, err)
	}
	return orderDetails, nil
}

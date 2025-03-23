package order

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateOrderDetails(orderDetails *model.OrderDetails) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO order_details (user_id, payment_id, total) VALUES ($1, $2, $3) RETURNING order_details_id",
		orderDetails.UserID, orderDetails.PaymentID, orderDetails.Total,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateOrderDetails: %v", err)
	}

	return lastInsertedID, nil
}

package order

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetAllOrderDetails() ([]model.OrderDetails, error) {
	var allOrderDetails []model.OrderDetails

	rows, err := r.Client.Query("SELECT * FROM order_details")
	if err != nil {
		return allOrderDetails, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderDetails model.OrderDetails
		if err := rows.Scan(&orderDetails.OrderDetailsID, &orderDetails.UserID, &orderDetails.PaymentID, &orderDetails.Total); err != nil {
			return allOrderDetails, fmt.Errorf("GetAllOrderDetails %d: %v", orderDetails.OrderDetailsID, err)
		}
		allOrderDetails = append(allOrderDetails, orderDetails)
	}
	if err := rows.Err(); err != nil {
		return allOrderDetails, fmt.Errorf("GetAllOrderDetails %v:", err)
	}
	return allOrderDetails, nil
}

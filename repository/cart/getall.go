package cart

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetAllCartItems() ([]model.CartItem, error) {
	var cartItems []model.CartItem

	rows, err := r.Client.Query("SELECT * FROM cart_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.UserID, &cartItem.ProductID, &cartItem.Size, &cartItem.Quantity); err != nil {
			return nil, fmt.Errorf("GetAllCartItems %d: %v", cartItem.UserID, err)
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllCartItems %v:", err)
	}
	return cartItems, nil
}

package cart

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetCartItemsByUserID(userID int) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	rows, err := r.Client.Query("SELECT * FROM cart_items WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.UserID, &cartItem.ProductID, &cartItem.Size, &cartItem.Quantity); err != nil {
			return nil, fmt.Errorf("GetCartItemsByUserID %d: %v", cartItem.UserID, err)
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetCartItemsByUserID %v:", err)
	}
	return cartItems, nil
}

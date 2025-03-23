package cartitem

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetCartItemsByCartItemID(cartItemID int) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	rows, err := r.Client.Query("SELECT * FROM cart_items WHERE cart_item_id = $1", cartItemID)
	if err != nil {
		return cartItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.CartItemID, &cartItem.ShoppingSessionID, &cartItem.ProductID, &cartItem.Quantity); err != nil {
			return cartItems, fmt.Errorf("GetCartItemsByCartItemID %d: %v", cartItem.CartItemID, err)
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := rows.Err(); err != nil {
		return cartItems, fmt.Errorf("GetCartItemsByCartItemID %v:", err)
	}
	return cartItems, nil
}

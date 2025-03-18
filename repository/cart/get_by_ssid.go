package cart

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetCartItemsByShoppingSessionID(shoppingSessionID int) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	rows, err := r.Client.Query("SELECT * FROM cart_items WHERE shopping_session_id = $1", shoppingSessionID)
	if err != nil {
		return cartItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.CartItemID, &cartItem.ShoppingSessionID, &cartItem.ProductID, &cartItem.Quantity); err != nil {
			return cartItems, fmt.Errorf("GetCartItemsByShoppingSessionID %d: %v", cartItem.ShoppingSessionID, err)
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := rows.Err(); err != nil {
		return cartItems, fmt.Errorf("GetCartItemsByShoppingSessionID %v:", err)
	}
	return cartItems, nil
}

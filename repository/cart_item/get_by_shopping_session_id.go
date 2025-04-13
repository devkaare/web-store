package cartitem

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetCartItemsByShoppingSessionID(shoppingSessionID string) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	rows, err := r.Client.Query("SELECT * FROM cart_items WHERE shopping_session_id = $1", shoppingSessionID)
	if err != nil {
		return cartItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.CartItemID, &cartItem.ShoppingSessionID, &cartItem.ProductID, &cartItem.Quantity); err != nil {
			return cartItems, fmt.Errorf("GetCartItemsByShoppingSessionID %s: %v", cartItem.ShoppingSessionID, err)
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := rows.Err(); err != nil {
		if err == sql.ErrNoRows {
			return cartItems, nil
		}

		return cartItems, fmt.Errorf("GetCartItemsByShoppingSessionID %v:", err)
	}
	return cartItems, nil
}

package cart

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateCartItem(cartItem *model.CartItem) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO cart_items (shopping_session_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING cart_item_id",
		cartItem.ShoppingSessionID, cartItem.ProductID, cartItem.Quantity,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateCartItem: %v", err)
	}

	return lastInsertedID, nil
}

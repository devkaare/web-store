package cartitem

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetCartItemByCartItemID(cartItemID int) (*model.CartItem, error) {
	cartItem := &model.CartItem{}

	row := r.Client.QueryRow("SELECT * FROM cart_items WHERE cart_item_id = $1", cartItemID)
	if err := row.Scan(&cartItem.CartItemID, &cartItem.ShoppingSessionID, &cartItem.ProductID, &cartItem.Quantity); err != nil {
		if err == sql.ErrNoRows {
			return cartItem, err
		}
		return cartItem, fmt.Errorf("GetCartItemByCartItemID %d: %v", cartItemID, err)
	}
	return cartItem, nil
}

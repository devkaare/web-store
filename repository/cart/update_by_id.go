package cart

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateCartItemQuantity(cartItem *model.CartItem) error {
	_, err := r.Client.Exec("UPDATE cart_items SET quantity = $3 WHERE cart_item_id = $1 AND product_id = $2", cartItem.CartItemID, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return fmt.Errorf("UpdateCartItemQuantity: %v", err)
	}
	return nil
}

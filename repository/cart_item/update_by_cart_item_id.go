package cartitem

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateCartItemQuantityByCartItemID(cartItem *model.CartItem) error {
	_, err := r.Client.Exec("UPDATE cart_items SET quantity = $3 WHERE cart_item_id = $1", cartItem.CartItemID, cartItem.Quantity)
	if err != nil {
		return fmt.Errorf("UpdateCartItemQuantity: %v", err)
	}
	return nil
}

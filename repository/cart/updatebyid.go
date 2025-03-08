package cart

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateCartItemQuantity(cartItem *model.CartItem) error {
	_, err := r.Client.Exec("UPDATE cart_items SET quantity = $4 WHERE user_id = $1 AND product_id = $2 AND sizes = $3", cartItem.UserID, cartItem.ProductID, cartItem.Size, cartItem.Quantity)
	if err != nil {
		return fmt.Errorf("UpdateCartItemQuantity: %v", err)
	}
	return nil
}

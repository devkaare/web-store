package cartitem

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) DeleteCartItemByCartItemID(cartItem *model.CartItem) error {
	result, err := r.Client.Exec("DELETE FROM cart_items WHERE cart_item_id = $1", cartItem.CartItemID, cartItem.ShoppingSessionID, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return fmt.Errorf("DeleteCartItemByCartItemID %d, %v", cartItem.CartItemID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteCartItemByCartItemID %d: %v", cartItem.CartItemID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteCartItemByCartItemID %d: cart item not found", cartItem.CartItemID)
	}
	return nil
}

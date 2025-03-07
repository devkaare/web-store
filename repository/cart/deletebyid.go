package cart

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *CartRepo) DeleteCartItem(cartItem *model.CartItem) error {
	result, err := r.Client.Exec("DELETE FROM cart_items WHERE user_id = $1 AND product_id = $2 AND sizes = $3", cartItem.UserID, cartItem.ProductID, cartItem.Size)
	if err != nil {
		return fmt.Errorf("DeleteCartItem %d, %v", cartItem.UserID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteCartItem %d: %v", cartItem.UserID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteCartItem %d: cart item not found", cartItem.UserID)
	}
	return nil
}

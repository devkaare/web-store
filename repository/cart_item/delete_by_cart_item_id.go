package cartitem

import (
	"fmt"
)

func (r *Repo) DeleteCartItemByCartItemID(cartItemID int) error {
	result, err := r.Client.Exec("DELETE FROM cart_items WHERE cart_item_id = $1", cartItemID)
	if err != nil {
		return fmt.Errorf("DeleteCartItemByCartItemID %d: %v", cartItemID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteCartItemByCartItemID %d: %v", cartItemID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteCartItemByCartItemID %d: cart item not found", cartItemID)
	}
	return nil
}

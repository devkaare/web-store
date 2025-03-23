package shoppingsession

import (
	"fmt"
)

func (r *Repo) DeleteShoppingSessionBySessionID(shoppingSessionID int) error {
	result, err := r.Client.Exec("DELETE FROM shopping_sessions WHERE shopping_session_id = $1", shoppingSessionID)
	if err != nil {
		return fmt.Errorf("DeleteShoppingSessionBySessionID %d: %v", shoppingSessionID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteShoppingSessionBySessionID %d: %v", shoppingSessionID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteShoppingSessionBySessionID %d: shoppingSession not found", shoppingSessionID)
	}
	return nil
}

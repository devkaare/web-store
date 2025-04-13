package shoppingsession

import (
	"fmt"
)

func (r *Repo) DeleteShoppingSessionBySessionID(shoppingSessionID string) error {
	result, err := r.Client.Exec("DELETE FROM shopping_sessions WHERE shopping_session_id = $1", shoppingSessionID)
	if err != nil {
		return fmt.Errorf("DeleteShoppingSessionBySessionID %s: %v", shoppingSessionID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteShoppingSessionBySessionID %s: %v", shoppingSessionID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteShoppingSessionBySessionID %s: shoppingSession not found", shoppingSessionID)
	}
	return nil
}

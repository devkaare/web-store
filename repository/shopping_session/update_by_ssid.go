package shoppingsession

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateShoppingSessionByShoppinSessionID(shoppingSession *model.ShoppingSession) error {
	_, err := r.Client.Exec("UPDATE shopping_sessions SET total = $2 WHERE shopping_session_id = $1", shoppingSession.ShoppingSessionID, shoppingSession.Total)
	if err != nil {
		return fmt.Errorf("UpdateShoppingSessionByShoppinSessionID: %v", err)
	}
	return nil
}

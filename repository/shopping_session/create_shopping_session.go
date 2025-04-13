package shoppingsession

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateShoppingSession(shoppingSession *model.ShoppingSession) error {
	err := r.Client.QueryRow(
		"INSERT INTO shopping_sessions (shopping_session_id, session_id, total) VALUES ($1, $2, $3)",
		shoppingSession.ShoppingSessionID, shoppingSession.SessionID, shoppingSession.Total,
	)
	if err != nil {
		return fmt.Errorf("CreateShoppingSession: %v", err)
	}

	return nil
}

package shoppingsession

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateShoppingSession(shoppingSession *model.ShoppingSession) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO shopping_sessions (user_id, total) VALUES ($1, $2) RETURNING shopping_session_id",
		shoppingSession.UserID, shoppingSession.Total,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateShoppingSession: %v", err)
	}

	return lastInsertedID, nil
}

package shoppingsession

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetShoppingSessionBySessionID(sessionID string) (*model.ShoppingSession, error) {
	shoppingSession := &model.ShoppingSession{}

	row := r.Client.QueryRow("SELECT * FROM shopping_sessions WHERE session_id = $1", sessionID)
	if err := row.Scan(&shoppingSession.SessionID, &shoppingSession.ShoppingSessionID, &shoppingSession.Total); err != nil {
		if err == sql.ErrNoRows {
			return shoppingSession, err
		}
		return shoppingSession, fmt.Errorf("GetShoppingSessionBySessionID %s: %v", sessionID, err)
	}
	return shoppingSession, nil

}

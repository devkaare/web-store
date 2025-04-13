package shoppingsession

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetShoppingSessionByShoppingSessionID(shoppingSessionID string) (*model.ShoppingSession, error) {
	shoppingSession := &model.ShoppingSession{}

	row := r.Client.QueryRow("SELECT * FROM shopping_sessions WHERE shopping_session_id = $1", shoppingSessionID)
	if err := row.Scan(&shoppingSession.ShoppingSessionID, &shoppingSession.SessionID, &shoppingSession.Total); err != nil {
		if err == sql.ErrNoRows {
			return shoppingSession, err
		}
		return shoppingSession, fmt.Errorf("GetShoppingSessionByShoppingSessionID %s: %v", shoppingSessionID, err)
	}
	return shoppingSession, nil

}

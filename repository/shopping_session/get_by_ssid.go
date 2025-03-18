package shoppingsession

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetShoppingSessionBySessionID(shoppingSessionID int) (*model.ShoppingSession, error) {
	shoppingSession := &model.ShoppingSession{}

	row := r.Client.QueryRow("SELECT * FROM shopping_sessions WHERE shopping_session_id = $1", shoppingSessionID)
	if err := row.Scan(&shoppingSession.ShoppingSessionID, &shoppingSession.UserID, &shoppingSession.Total); err != nil {
		if err == sql.ErrNoRows {
			return shoppingSession, err
		}
		return shoppingSession, fmt.Errorf("GetShoppingSessionBySessionID %d: %v", shoppingSessionID, err)
	}
	return shoppingSession, nil

}

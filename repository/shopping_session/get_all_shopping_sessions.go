package shoppingsession

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetAllShoppingSessions() ([]model.ShoppingSession, error) {
	var shoppingSessions []model.ShoppingSession

	rows, err := r.Client.Query("SELECT * FROM shopping_sessions")
	if err != nil {
		return shoppingSessions, err
	}
	defer rows.Close()

	for rows.Next() {
		var shoppingSession model.ShoppingSession
		if err := rows.Scan(&shoppingSession.ShoppingSessionID, &shoppingSession.SessionID, &shoppingSession.Total); err != nil {
			return shoppingSessions, fmt.Errorf("GetAllShoppingSessions %d: %v", shoppingSession.ShoppingSessionID, err)
		}
		shoppingSessions = append(shoppingSessions, shoppingSession)
	}
	if err := rows.Err(); err != nil {
		return shoppingSessions, fmt.Errorf("GetAllShoppingSessions %v:", err)
	}
	return shoppingSessions, nil
}

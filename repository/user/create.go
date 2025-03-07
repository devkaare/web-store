package user

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *UserRepo) CreateUser(user *model.User) (int, error) {
	var lastInsertedID int = 0
	err := r.Client.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING user_id",
		user.Email, user.Password,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateUser: %v", err)
	}

	return lastInsertedID, nil
}

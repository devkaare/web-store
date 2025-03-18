package user

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateUser(user *model.User) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING user_id",
		user.FirstName, user.LastName, user.Email, user.Password,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateUser: %v", err)
	}

	return lastInsertedID, nil
}

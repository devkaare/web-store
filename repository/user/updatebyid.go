package user

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateUserByUserID(user *model.User) error {
	_, err := r.Client.Exec("UPDATE users SET first_name = $2, last_name = $3, email = $4, password = $5 WHERE user_id = $1", user.UserID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("UpdateUserByUserID: %v", err)
	}
	return nil
}

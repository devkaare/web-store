package user

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *UserRepo) UpdateUserByUserID(user *model.User) error {
	_, err := r.Client.Exec("UPDATE users SET email = $2, password = $3 WHERE user_id = $1", user.UserID, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("UpdateUserByUserID: %v", err)
	}
	return nil
}

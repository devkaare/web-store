package user

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *UserRepo) GetUserByUserID(userID int) (*model.User, error) {
	user := &model.User{}

	row := r.Client.QueryRow("SELECT * FROM users WHERE user_id = $1", userID)
	if err := row.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, fmt.Errorf("GetUserByUserID %d: %v", userID, err)
	}
	return user, nil

}

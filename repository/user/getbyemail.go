package user

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}

	row := r.Client.QueryRow("SELECT * FROM users WHERE email = $1", email)
	if err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("GetUserByEmail %s: %v", email, err)
	}
	return user, nil

}

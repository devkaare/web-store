package user

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetAllUsers() ([]model.User, error) {
	var users []model.User

	rows, err := r.Client.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
			return users, fmt.Errorf("GetAllUsers %d: %v", user.UserID, err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return users, fmt.Errorf("GetAllUsers %v:", err)
	}
	return users, nil
}

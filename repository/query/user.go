package query

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *PostgresRepo) CreateUser(user *model.User) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING user_id",
		user.Email, user.Password,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateUser: %v", err)
	}

	return lastInsertedID, nil
}

func (r *PostgresRepo) GetUsers() ([]model.User, error) {
	var users []model.User

	rows, err := r.Client.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("GetUsers %d: %v", user.UserID, err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetUsers %v:", err)
	}
	return users, nil
}

func (r *PostgresRepo) GetUserByUserID(userID uint32) (*model.User, bool, error) {
	user := &model.User{}

	row := r.Client.QueryRow("SELECT * FROM users WHERE user_id = $1", userID)
	if err := row.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, false, nil
		}
		return user, false, fmt.Errorf("GetUserByUserID %d: %v", userID, err)
	}
	return user, true, nil

}

func (r *PostgresRepo) UpdateUserByUserID(user *model.User) error {
	_, err := r.Client.Exec("UPDATE users SET email = $2, password = $3 WHERE user_id = $1", user.UserID, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("UpdateUserByUserID: %v", err)
	}
	return nil
}

func (r *PostgresRepo) DeleteUserByUserID(userID uint32) error {
	result, err := r.Client.Exec("DELETE FROM users WHERE user_id = $1", userID)
	if err != nil {
		return fmt.Errorf("DeleteUserByUserID %d, %v", userID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteUserByUserID %d: %v", userID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteUserByUserID %d: no such user", userID)
	}
	return nil
}

func (r *PostgresRepo) GetUserByEmail(email string) (*model.User, bool, error) {
	user := &model.User{}

	row := r.Client.QueryRow("SELECT * FROM users WHERE email = $2", email)
	if err := row.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, false, nil
		}
		return user, false, fmt.Errorf("GetUserByEmail %s: %v", email, err)
	}
	return user, true, nil

}

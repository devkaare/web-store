package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/devkaare/web-store/model"
)

type PostgresRepo struct {
	Client *sql.DB
}

func (r *PostgresRepo) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := r.Client.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err)
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	return stats
}

func (r *PostgresRepo) Close() error {
	log.Println("Disconnected from database")
	return r.Client.Close()
}

func (r *PostgresRepo) CreateUser(user *model.User) error {
	_, err := r.Client.Exec(
		"INSERT INTO users (user_id, email, password) VALUES ($1, $2, $3)",
		user.UserID, user.Email, user.Password,
	)
	if err != nil {
		return fmt.Errorf("CreateUser: %v", err)
	}

	return nil
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
		return nil, fmt.Errorf("GetTodoList %v:", err)
	}
	return users, nil
}

func (r *PostgresRepo) GetUserByID(userID uint32) (model.User, error) {
	var user model.User
	row := r.Client.QueryRow("SELECT * FROM users WHERE user_id = $1", userID)
	if err := row.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, fmt.Errorf("GetTodoByID %d: %v", userID, err)
	}
	return user, nil

}

// func (r *PostgresRepo) UpdateTodoByID(todo *model.Todo) error {
// 	_, err := r.Client.Exec("UPDATE todo SET title = $2, description = $3 WHERE id = $1", todo.ID, todo.Title, todo.Description)
// 	if err != nil {
// 		return fmt.Errorf("UpdateTodoByID: %v", err)
// 	}
// 	return nil
// }
//
// func (r *PostgresRepo) DeleteTodoByID(id uint32) error {
// 	result, err := r.Client.Exec("DELETE FROM todo WHERE id = $1", id)
// 	if err != nil {
// 		return fmt.Errorf("DeleteTodoByID %d, %v", id, err)
// 	}
// 	count, err := result.RowsAffected()
// 	if err != nil {
// 		return fmt.Errorf("DeleteTodoByID %d: %v", id, err)
// 	}
// 	if count < 1 {
// 		return fmt.Errorf("DeleteTodoByID %d: no such todo", id)
// 	}
// 	return nil
// }

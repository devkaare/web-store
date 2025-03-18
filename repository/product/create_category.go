package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateCategory(category *model.Category) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO categories (name) VALUES ($1) RETURNING category_id",
		category.Name,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateCategory: %v", err)
	}

	return lastInsertedID, nil
}

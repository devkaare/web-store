package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category

	rows, err := r.Client.Query("SELECT * FROM categories")
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.CategoryID, &category.Name); err != nil {
			return categories, fmt.Errorf("GetAllCategories %d: %v", category.CategoryID, err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return categories, fmt.Errorf("GetAllCategories %v:", err)
	}
	return categories, nil
}

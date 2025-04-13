package product

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetCategoryByCategoryName(categoryName string) (*model.Category, error) {
	category := &model.Category{}

	row := r.Client.QueryRow("SELECT * FROM categories WHERE name = $1", categoryName)
	if err := row.Scan(&category.CategoryID, &category.Name); err != nil {
		if err == sql.ErrNoRows {
			return category, err
		}
		return category, fmt.Errorf("GetCategoryByCategoryID %s: %v", categoryName, err)
	}
	return category, nil
}

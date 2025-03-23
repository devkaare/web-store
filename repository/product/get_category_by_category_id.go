package product

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetCategoryByCategoryID(categoryID int) (*model.Category, error) {
	category := &model.Category{}

	row := r.Client.QueryRow("SELECT * FROM categories WHERE category_id = $1", categoryID)
	if err := row.Scan(&category.CategoryID, &category.Name); err != nil {
		if err == sql.ErrNoRows {
			return category, err
		}
		return category, fmt.Errorf("GetCategoryByCategoryID %d: %v", categoryID, err)
	}
	return category, nil
}

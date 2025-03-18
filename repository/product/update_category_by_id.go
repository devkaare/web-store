package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateCategoryByCategoryID(category *model.Category) error {
	_, err := r.Client.Exec("UPDATE categories SET name = $2 WHERE category_id = $1", category.CategoryID, category.Name)
	if err != nil {
		return fmt.Errorf("UpdateCategoryByCategoryID: %v", err)
	}
	return nil
}

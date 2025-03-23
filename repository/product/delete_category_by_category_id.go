package product

import (
	"fmt"
)

func (r *Repo) DeleteCategoryByCategoryID(categoryID int) error {
	result, err := r.Client.Exec("DELETE FROM categories WHERE category_id = $1", categoryID)
	if err != nil {
		return fmt.Errorf("DeleteCategoryByCategoryID %d: %v", categoryID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteCategoryByCategoryID %d: %v", categoryID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteCategoryByCategoryID %d: category not found", categoryID)
	}
	return nil
}

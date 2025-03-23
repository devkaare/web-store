package product

import (
	"fmt"
)

func (r *Repo) DeleteProductByProductID(productID int) error {
	result, err := r.Client.Exec("DELETE FROM products WHERE product_id = $1", productID)
	if err != nil {
		return fmt.Errorf("DeleteProductByProductID %d: %v", productID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteProductByProductID %d: %v", productID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteProductByProductID %d: product not found", productID)
	}
	return nil
}

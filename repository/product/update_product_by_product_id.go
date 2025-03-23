package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) UpdateProductByProductID(product *model.Product) error {
	_, err := r.Client.Exec("UPDATE products SET category_id = $2, name = $3, description = $4, price = $5, size = $6, color = $7, image_path = $8 WHERE product_id = $1", product.ProductID, product.CategoryID, product.Name, product.Description, product.Price, product.Size, product.Color, product.ImagePath)
	if err != nil {
		return fmt.Errorf("UpdateProductByProductID: %v", err)
	}
	return nil
}

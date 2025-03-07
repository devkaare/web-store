package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *ProductRepo) UpdateProductByProductID(product *model.Product) error {
	_, err := r.Client.Exec("UPDATE products SET name = $2, price = $3, sizes = $4, image_path = $5 WHERE product_id = $1", product.ProductID, product.Name, product.Price, product.Sizes, product.ImagePath)
	if err != nil {
		return fmt.Errorf("UpdateProductByProductID: %v", err)
	}
	return nil
}

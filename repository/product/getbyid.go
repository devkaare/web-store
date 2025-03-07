package product

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *ProductRepo) GetProductByProductID(productID uint32) (*model.Product, error) {
	product := &model.Product{}

	row := r.Client.QueryRow("SELECT * FROM products WHERE product_id = $1", productID)
	if err := row.Scan(&product.ProductID, &product.Name, &product.Price, &product.Sizes, &product.ImagePath); err != nil {
		if err == sql.ErrNoRows {
			return product, err
		}
		return product, fmt.Errorf("GetProductByProductID %d: %v", productID, err)
	}
	return product, nil
}

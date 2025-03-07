package product

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *ProductRepo) GetProductByName(productName string) (*model.Product, error) {
	product := &model.Product{}

	row := r.Client.QueryRow("SELECT * FROM products WHERE name = $1", productName)
	if err := row.Scan(&product.ProductID, &product.Name, &product.Price, &product.Sizes, &product.ImagePath); err != nil {
		if err == sql.ErrNoRows {
			return product, err
		}
		return product, fmt.Errorf("GetProductByName %s: %v", productName, err)
	}
	return product, nil
}

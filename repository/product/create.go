package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *ProductRepo) CreateProduct(product *model.Product) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO products (name, price, sizes, image_path) VALUES ($1, $2, $3, $4) RETURNING product_id",
		product.Name, product.Price, product.Sizes, product.ImagePath,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateProduct: %v", err)
	}

	return lastInsertedID, nil
}

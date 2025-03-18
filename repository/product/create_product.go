package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateProduct(product *model.Product) (int, error) {
	lastInsertedID := 0
	err := r.Client.QueryRow(
		"INSERT INTO products (category_id, name, description, price, size, color, image_path) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING product_id",
		product.CategoryID, product.Name, product.Description, product.Price, product.Size, product.Color, product.ImagePath,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateProduct: %v", err)
	}

	return lastInsertedID, nil
}

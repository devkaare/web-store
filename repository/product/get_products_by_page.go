package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetProductsByPage(page int) ([]model.Product, error) {
	var products []model.Product

	limit := 4
	offset := (page - 1) * limit

	rows, err := r.Client.Query("SELECT product_id, category_id, name, description, price, size, color, image_path FROM products LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ProductID, &product.CategoryID, &product.Name, &product.Description, &product.Price, &product.Size, &product.Color, &product.ImagePath); err != nil {
			return products, fmt.Errorf("GetProductsByPage %d: %v", product.ProductID, err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return products, fmt.Errorf("GetProductsByPage %v:", err)
	}
	return products, nil
}

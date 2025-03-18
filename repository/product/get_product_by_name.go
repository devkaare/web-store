package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetProductsBySearch(search string) ([]model.Product, error) {
	var products []model.Product

	rows, err := r.Client.Query("SELECT product_id, name, price, sizes, image_path FROM products WHERE name ~* '\\b$1\\b'", search)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ProductID, &product.Name, &product.Price, &product.Sizes, &product.ImagePath); err != nil {
			return products, fmt.Errorf("GetProductsBySearch %d: %v", product.ProductID, err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return products, fmt.Errorf("GetProductsBySearch %v:", err)
	}
	return products, nil
}

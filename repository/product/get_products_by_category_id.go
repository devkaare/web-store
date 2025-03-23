package product

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetProductsByCategoryID(categoryID int) ([]model.Product, error) {
	var products []model.Product

	rows, err := r.Client.Query("SELECT product_id, category_id, name, description, price, image_path FROM products WHERE category_id = $1", categoryID)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ProductID, &product.CategoryID, &product.Name, &product.Description, &product.Price, &product.ImagePath); err != nil {
			return products, fmt.Errorf("GetProductsByCategoryID %d: %v", product.ProductID, err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return products, fmt.Errorf("GetProductsByCategoryID %v:", err)
	}
	return products, nil
}

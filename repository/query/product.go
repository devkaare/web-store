package query

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *PostgresRepo) CreateProduct(product *model.Product) (uint32, error) {
	var lastInsertedID uint32 = 0
	err := r.Client.QueryRow(
		"INSERT INTO products (name, price, sizes, image_path) VALUES ($1, $2, $3, $4) RETURNING product_id",
		product.Name, product.Price, product.Sizes, product.ImagePath,
	).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, fmt.Errorf("CreateProduct: %v", err)
	}

	return lastInsertedID, nil
}

func (r *PostgresRepo) GetProducts() ([]model.Product, error) {
	var products []model.Product

	rows, err := r.Client.Query("SELECT * FROM products")
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ProductID, &product.Name, &product.Price, &product.Sizes, &product.ImagePath); err != nil {
			return products, fmt.Errorf("GetProducts %d: %v", product.ProductID, err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return products, fmt.Errorf("GetProducts %v:", err)
	}
	return products, nil
}

func (r *PostgresRepo) GetProductByProductID(productID uint32) (*model.Product, error) {
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

func (r *PostgresRepo) GetProductsByPage(page uint32) ([]model.Product, error) {
	var products []model.Product

	limit := 10
	offset := (int(page) - 1) * limit

	rows, err := r.Client.Query("SELECT product_id, name, price, sizes, image_path FROM products LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ProductID, &product.Name, &product.Price, &product.Sizes, &product.ImagePath); err != nil {
			return products, fmt.Errorf("GetProductsByPage %d: %v", product.ProductID, err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return products, fmt.Errorf("GetProductsByPage %v:", err)
	}
	return products, nil
}

func (r *PostgresRepo) UpdateProductByProductID(product *model.Product) error {
	_, err := r.Client.Exec("UPDATE products SET name = $2, price = $3, sizes = $4, image_path = $5 WHERE product_id = $1", product.ProductID, product.Name, product.Price, product.Sizes, product.ImagePath)
	if err != nil {
		return fmt.Errorf("UpdateProductByProductID: %v", err)
	}
	return nil
}

func (r *PostgresRepo) DeleteProductByProductID(productID uint32) error {
	result, err := r.Client.Exec("DELETE FROM products WHERE product_id = $1", productID)
	if err != nil {
		return fmt.Errorf("DeleteProductByProductID %d, %v", productID, err)
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

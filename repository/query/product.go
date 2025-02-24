package query

import (
	// "database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *PostgresRepo) CreateProduct(product *model.Product) (int, error) {
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

func (r *PostgresRepo) GetProducts() ([]model.Product, error) {
	var products []model.Product

	rows, err := r.Client.Query("SELECT * FROM products")
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ProductID, &product.Name, &product.Price, product.Sizes, product.ImagePath); err != nil {
			return nil, fmt.Errorf("GetProducts %d: %v", product.ProductID, err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetProducts %v:", err)
	}
	return products, nil
}

// func (r *PostgresRepo) GetProductByProductID(userID uint32) (*model.Product, bool, error) {
// 	user := &model.Product{}
//
// 	row := r.Client.QueryRow("SELECT * FROM users WHERE user_id = $1", userID)
// 	if err := row.Scan(&product.ProductID, &product.Email, &product.Password); err != nil {
// 		if err == sql.ErrNoRows {
// 			return user, false, err
// 		}
// 		return user, false, fmt.Errorf("GetProductByProductID %d: %v", userID, err)
// 	}
// 	return user, true, nil
//
// }
//
// func (r *PostgresRepo) UpdateProductByProductID(user *model.Product) error {
// 	_, err := r.Client.Exec("UPDATE user SET email = $2, password = $3 WHERE user_id = $1", product.ProductID, product.Email, product.Password)
// 	if err != nil {
// 		return fmt.Errorf("UpdateProductByProductID: %v", err)
// 	}
// 	return nil
// }
//
// func (r *PostgresRepo) DeleteProductByProductID(userID uint32) error {
// 	result, err := r.Client.Exec("DELETE FROM users WHERE user_id = $1", userID)
// 	if err != nil {
// 		return fmt.Errorf("DeleteProductByProductID %d, %v", userID, err)
// 	}
// 	count, err := result.RowsAffected()
// 	if err != nil {
// 		return fmt.Errorf("DeleteProductByProductID %d: %v", userID, err)
// 	}
// 	if count < 1 {
// 		return fmt.Errorf("DeleteProductByProductID %d: no such user", userID)
// 	}
// 	return nil
// }

package query

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *PostgresRepo) CreateCartItem(cartItem *model.CartItem) error {
	_, err := r.Client.Exec(
		"INSERT INTO cart_items (user_id, product_id, size, quantity) VALUES ($1, $2, $3, $4)",
		cartItem.UserID, cartItem.ProductID, cartItem.Size, cartItem.Quantity,
	)
	if err != nil {
		return fmt.Errorf("CreateCartItem: %v", err)
	}

	return nil
}

func (r *PostgresRepo) GetCartItems() ([]model.CartItem, error) {
	var cartItems []model.CartItem

	rows, err := r.Client.Query("SELECT * FROM cart_items")
	if err != nil {
		return cartItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.UserID, &cartItem.ProductID, &cartItem.Size, cartItem.Quantity); err != nil {
			return nil, fmt.Errorf("GetCartItems %d: %v", cartItem.UserID, err)
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetCartItems %v:", err)
	}
	return cartItems, nil
}

func (r *PostgresRepo) GetCartItemsByUserID(userID uint32) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	rows, err := r.Client.Query("SELECT * FROM cart_items WHERE user_id = $1", userID)
	if err != nil {
		return cartItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.UserID, &cartItem.ProductID, &cartItem.Size, cartItem.Quantity); err != nil {
			return nil, fmt.Errorf("GetCartItemsByUserID %d: %v", cartItem.UserID, err)
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetCartItemsByUserID %v:", err)
	}
	return cartItems, nil
}

func (r *PostgresRepo) UpdateCartItemQuantity(cartItem *model.CartItem) error {
	_, err := r.Client.Exec("UPDATE cart_items SET quantity = $4 WHERE user_id = $1 AND product_id = $2 AND size = $3", cartItem.UserID, cartItem.ProductID, cartItem.Size, cartItem.Quantity)
	if err != nil {
		return fmt.Errorf("UpdateCartItemQuantity: %v", err)
	}
	return nil
}

func (r *PostgresRepo) DeleteCartItem(cartItem *model.CartItem) error {
	result, err := r.Client.Exec("DELETE FROM cart_items WHERE user_id = $1 AND product_id = $2 AND size = $3", cartItem.UserID, cartItem.ProductID, cartItem.Size)
	if err != nil {
		return fmt.Errorf("DeleteCartItem %d, %v", cartItem.UserID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteCartItem %d: %v", cartItem.UserID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteCartItem %d: no such cart item", cartItem.UserID)
	}
	return nil
}

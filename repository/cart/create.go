package cart

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *CartRepo) CreateCartItem(cartItem *model.CartItem) error {
	_, err := r.Client.Exec(
		"INSERT INTO cart_items (user_id, product_id, sizes, quantity) VALUES ($1, $2, $3, $4)",
		cartItem.UserID, cartItem.ProductID, cartItem.Size, cartItem.Quantity,
	)
	if err != nil {
		return fmt.Errorf("CreateCartItem: %v", err)
	}

	return nil
}

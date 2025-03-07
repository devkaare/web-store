package model

type CartItem struct {
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Size      string `json:"size"`
	Quantity  int    `json:"quantity"`
}

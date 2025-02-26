package model

type CartItem struct {
	UserID    uint32 `json:"user_id"`
	ProductID uint32 `json:"product_id"`
	Size      string `json:"size"`
	Quantity  uint32 `json:"quantity"`
}

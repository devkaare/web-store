package model

type CartItem struct {
	ProductID uint32
	Size      string
	Quantity  uint32
}

type Cart struct {
	UserID uint32
	Items  []CartItem
}

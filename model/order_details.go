package model

type OrderDetails struct {
	OrderDetailsID int
	UserID         int
	PaymentID      string
	Total          int
	PaymentStatus  string
}

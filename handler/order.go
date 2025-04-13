package handler

import (
	"database/sql"
	"net/http"

	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/cart_item"
	"github.com/devkaare/web-store/repository/order"
	"github.com/devkaare/web-store/repository/session"
	"github.com/devkaare/web-store/repository/shopping_session"
	"github.com/google/uuid"
)

type Order struct {
	OrderRepo       *order.Repo
	CartItemRepo    *cartitem.Repo
	ShoppingSession *shoppingsession.Repo
	Session         *session.Repo
}

var orderHandler = &Order{
	OrderRepo: &order.Repo{},
}

func NewOrderHandler(db *sql.DB) *Order {
	orderHandler.OrderRepo = repository.GetOrder(func() *sql.DB {
		return db
	})
	return orderHandler
}

func (o *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	shoppingSessionID := ""

	shoppingSession, err := o.ShoppingSession.GetShoppingSessionByShoppingSessionID(shoppingSessionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := o.Session.GetSessionBySessionID(shoppingSession.SessionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	orderDetails := &model.OrderDetails{
		UserID:    user.UserID,
		PaymentID: uuid.New().String(),
		Total:     shoppingSession.Total,
	}

	orderDetailsID, err := o.OrderRepo.CreateOrderDetails(orderDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartItems, err := o.CartItemRepo.GetCartItemsByShoppingSessionID(shoppingSessionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, ci := range cartItems {
		orderItem := &model.OrderItem{}

		orderItem.OrderItemID = ci.CartItemID
		orderItem.ProductID = ci.ProductID
		orderItem.Quantity = ci.Quantity

		orderItem.OrderDetailsID = orderDetailsID

		_, err := o.OrderRepo.CreateOrderItem(orderItem)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}

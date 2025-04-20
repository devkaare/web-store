package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/cart_item"
	"github.com/devkaare/web-store/repository/order"
	"github.com/devkaare/web-store/repository/product"
	"github.com/devkaare/web-store/repository/session"
	"github.com/devkaare/web-store/repository/shopping_session"
	"github.com/devkaare/web-store/views"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stripe/stripe-go/v81"
	payment "github.com/stripe/stripe-go/v81/checkout/session"
)

type Order struct {
	OrderRepo       *order.Repo
	CartItemRepo    *cartitem.Repo
	ShoppingSession *shoppingsession.Repo
	Session         *session.Repo
	ProductRepo     *product.Repo
}

var orderHandler = &Order{
	OrderRepo:       &order.Repo{},
	CartItemRepo:    &cartitem.Repo{},
	ShoppingSession: &shoppingsession.Repo{},
	Session:         &session.Repo{},
	ProductRepo:     &product.Repo{},
}

func NewOrderHandler(db *sql.DB) *Order {
	orderHandler.OrderRepo = repository.GetOrder(func() *sql.DB {
		return db
	})
	orderHandler.CartItemRepo = repository.GetCartItem(func() *sql.DB {
		return db
	})
	orderHandler.ShoppingSession = repository.GetShoppingSession(func() *sql.DB {
		return db
	})
	orderHandler.Session = repository.GetSession(func() *sql.DB {
		return db
	})
	orderHandler.ProductRepo = repository.GetProduct(func() *sql.DB {
		return db
	})
	return orderHandler
}

func (o *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	shoppingSessionID := r.Context().Value("shopping_session_id").(string)

	shoppingSession, err := o.ShoppingSession.GetShoppingSessionByShoppingSessionID(shoppingSessionID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := o.Session.GetSessionBySessionID(shoppingSession.SessionID)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartItems, err := o.CartItemRepo.GetCartItemsByShoppingSessionID(shoppingSessionID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var items []*stripe.CheckoutSessionLineItemParams

	for _, ci := range cartItems {
		item := &model.OrderItem{}

		item.OrderItemID = ci.CartItemID
		item.ProductID = ci.ProductID
		item.Quantity = ci.Quantity
		item.OrderDetailsID = orderDetailsID

		_, err := o.OrderRepo.CreateOrderItem(item)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		product, err := o.ProductRepo.GetProductByProductID(item.ProductID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		items = append(items, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(product.Name),
				},
				UnitAmount: stripe.Int64(int64(product.Price * 100)),
			},
			Quantity: stripe.Int64(int64(ci.Quantity)),
		})
	}

	params := &stripe.CheckoutSessionParams{
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems:  items,
		SuccessURL: stripe.String("http://localhost:3000/checkout/success"),
		CancelURL:  stripe.String("http://localhost:3000/checkout/cancel"),
	}

	s, err := payment.New(params)
	if err != nil {
		log.Println(err)
	}

	// err = o.ShoppingSession.DeleteShoppingSessionBySessionID(shoppingSessionID)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	log.Println(orderDetailsID)

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}

func (o *Order) GetOrderByOrderDetailsID(w http.ResponseWriter, r *http.Request) {
	orderDetailsID, _ := strconv.Atoi(chi.URLParam(r, "order_details_id"))

	orderDetails, err := o.OrderRepo.GetOrderDetailsByOrderDetailsID(orderDetailsID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	items, err := o.OrderRepo.GetOrderItemsByOrderDetailsID(orderDetailsID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var itemProps []views.OrderItemProp

	for _, item := range items {
		var itemProp views.OrderItemProp

		product, err := o.ProductRepo.GetProductByProductID(item.ProductID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		itemProp.OrderDetailsID = item.OrderDetailsID
		itemProp.OrderItemID = item.OrderItemID
		itemProp.ProductID = item.ProductID
		itemProp.Quantity = item.Quantity
		itemProp.Name = product.Name
		itemProp.Price = product.Price
		itemProp.ImagePath = product.ImagePath

		itemProps = append(itemProps, itemProp)
	}

	templ.Handler(views.OrderPage(itemProps, orderDetails)).ServeHTTP(w, r)
}

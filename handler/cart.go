package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/cart"
	"github.com/go-chi/chi/v5"
)

type Cart struct {
	Repo *cart.Repo
}

var cartHandler = &Cart{
	Repo: &cart.Repo{},
}

type userID int

func NewCartHandler(db *sql.DB) *Cart {
	cartHandler.Repo = repository.GetCart(func() *sql.DB {
		return db
	})
	return cartHandler
}

func (c *Cart) GetAllCartItems(w http.ResponseWriter, r *http.Request) {
	cartItems, err := c.Repo.GetAllCartItems()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(cartItems)
	_, _ = w.Write(jsonResp)
}

func (c *Cart) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	productID, _ := strconv.Atoi(r.FormValue("productID"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	size := r.Form["sizes"][0]

	product := &model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Size:      size,
		Quantity:  quantity,
	}

	if err := c.Repo.CreateCartItem(product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Cart) GetCartItemsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	cartItems, err := c.Repo.GetCartItemsByUserID(userID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(cartItems)
	_, _ = w.Write(jsonResp)
}

func (c *Cart) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	productID, _ := strconv.Atoi(chi.URLParam(r, "product_id"))
	size := r.URL.Query().Get("size")

	cartItem := &model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Size:      size,
	}

	if err := c.Repo.DeleteCartItem(cartItem); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Cart) UpdateCartItemQuantity(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	productID, _ := strconv.Atoi(chi.URLParam(r, "product_id"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	size := r.FormValue("size")

	cartItem := &model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Size:      size,
		Quantity:  quantity,
	}

	if err := c.Repo.UpdateCartItemQuantity(cartItem); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

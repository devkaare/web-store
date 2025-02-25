package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository/query"
	"github.com/go-chi/chi/v5"
)

type CartItem struct {
	Repo *query.PostgresRepo
}

func (c *CartItem) GetCartItems(w http.ResponseWriter, r *http.Request) {
	cartItems, err := c.Repo.GetCartItems()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, cartItems)
}

// curl -X POST localhost:3000/carts -d "userID=1&productID=2&size=s&quantity=1"
func (c *CartItem) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.FormValue("userID"))
	productID, _ := strconv.Atoi(r.FormValue("productID"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	size := r.FormValue("size")

	product := &model.CartItem{
		UserID:    uint32(userID),
		ProductID: uint32(productID),
		Size:      size,
		Quantity:  uint32(quantity),
	}

	if err := c.Repo.CreateCartItem(product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *CartItem) GetCartItemsByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	cartItems, err := c.Repo.GetCartItemsByUserID(uint32(userID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, cartItems)
}

// curl -X DELETE localhost:3000/carts/1/2\?size=s
func (c *CartItem) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	productID, _ := strconv.Atoi(chi.URLParam(r, "productID"))
	size := r.URL.Query().Get("size")

	cartItem := &model.CartItem{
		UserID:    uint32(userID),
		ProductID: uint32(productID),
		Size:      size,
	}

	if err := c.Repo.DeleteCartItem(cartItem); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// curl -X PUT localhost:3000/carts/1/2 -d "size=s&quantity=9"
func (c *CartItem) UpdateCartItemQuantity(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	productID, _ := strconv.Atoi(chi.URLParam(r, "productID"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	size := r.FormValue("size")

	cartItem := &model.CartItem{
		UserID:    uint32(userID),
		ProductID: uint32(productID),
		Size:      size,
		Quantity:  uint32(quantity),
	}

	if err := c.Repo.UpdateCartItemQuantity(cartItem); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

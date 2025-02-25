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

func (c *CartItem) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	URLParam := chi.URLParam(r, "ID")
	userID, err := strconv.Atoi(URLParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	size := r.FormValue("size")
	rawQuantity := r.FormValue("quantity")
	rawProductID := r.FormValue("productID")

	quantity, err := strconv.Atoi(rawQuantity)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	productID, err := strconv.Atoi(rawProductID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	URLParam := chi.URLParam(r, "ID")
	userID, err := strconv.Atoi(URLParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartItems, err := c.Repo.GetCartItemsByUserID(uint32(userID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, cartItems)
}

func (c *CartItem) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	size := chi.URLParam(r, "size")
	rawUserID := chi.URLParam(r, "userID")
	rawProductID := chi.URLParam(r, "productID")

	userID, err := strconv.Atoi(rawUserID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	productID, err := strconv.Atoi(rawProductID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

func (c *CartItem) UpdateCartItemQuantity(w http.ResponseWriter, r *http.Request) {
	size := chi.URLParam(r, "size")
	rawQuantity := chi.URLParam(r, "quantity")
	rawUserID := chi.URLParam(r, "userID")
	rawProductID := chi.URLParam(r, "productID")

	quantity, err := strconv.Atoi(rawQuantity)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID, err := strconv.Atoi(rawUserID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	productID, err := strconv.Atoi(rawProductID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

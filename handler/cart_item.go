package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	// "strconv"

	// "github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/cart_item"
	// "github.com/go-chi/chi/v5"
)

type CartItem struct {
	Repo *cartitem.Repo
}

var cartItemHandler = &CartItem{
	Repo: &cartitem.Repo{},
}

type userID int

func NewCartItemHandler(db *sql.DB) *CartItem {
	cartItemHandler.Repo = repository.GetCartItem(func() *sql.DB {
		return db
	})
	return cartItemHandler
}

func (c *CartItem) GetAllCartItems(w http.ResponseWriter, r *http.Request) {
	cartItems, err := c.Repo.GetAllCartItems()
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(cartItems)
	_, _ = w.Write(jsonResp)
}

func (c *CartItem) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	// userID := r.Context().Value("user_id").(int)
	// productID, _ := strconv.Atoi(r.FormValue("productID"))
	// quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	// size := r.Form["sizes"][0]
	//
	// product := &model.CartItem{
	// 	UserID:    userID,
	// 	ProductID: productID,
	// 	Size:      size,
	// 	Quantity:  quantity,
	// }
	//
	// err := c.Repo.CreateCartItem(product)
	// if check(err) {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
}

func (c *CartItem) GetCartItemsByUserID(w http.ResponseWriter, r *http.Request) {
	// userID := r.Context().Value("user_id").(int)
	//
	// cartItems, err := c.Repo.GetCartItemsByUserID(userID)
	// if check(err) {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	//
	// w.Header().Set("Content-Type", "application/json")
	// jsonResp, _ := json.Marshal(cartItems)
	// _, _ = w.Write(jsonResp)
}

func (c *CartItem) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	// userID := r.Context().Value("user_id").(int)
	// productID, _ := strconv.Atoi(chi.URLParam(r, "product_id"))
	// size := r.URL.Query().Get("size")
	//
	// cartItem := &model.CartItem{
	// 	UserID:    userID,
	// 	ProductID: productID,
	// 	Size:      size,
	// }
	//
	// err := c.Repo.DeleteCartItemByCartItemID(cartItem)
	// if check(err) {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
}

func (c *CartItem) UpdateCartItemQuantity(w http.ResponseWriter, r *http.Request) {
	// userID := r.Context().Value("user_id").(int)
	// productID, _ := strconv.Atoi(chi.URLParam(r, "product_id"))
	// quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	// size := r.FormValue("size")
	//
	// cartItem := &model.CartItem{
	// 	UserID:    userID,
	// 	ProductID: productID,
	// 	Size:      size,
	// 	Quantity:  quantity,
	// }
	//
	// err := c.Repo.UpdateCartItemQuantity(cartItem)
	// if check(err) {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
}

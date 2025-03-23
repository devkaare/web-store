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
	"github.com/devkaare/web-store/views"
	"github.com/go-chi/chi/v5"
)

type CartItem struct {
	Repo *cartitem.Repo
}

var cartItemHandler = &CartItem{
	Repo: &cartitem.Repo{},
}

func NewCartItemHandler(db *sql.DB) *CartItem {
	cartItemHandler.Repo = repository.GetCartItem(func() *sql.DB {
		return db
	})
	return cartItemHandler
}

// func (c *CartItem) GetAllCartItems(w http.ResponseWriter, r *http.Request) {
// 	cartItems, err := c.Repo.GetAllCartItems()
// 	if err != nil {
// 		log.Printf("GetAllCartItems: error fetching cart items: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	// templ.Handler(views.CartPage(cartItems)).ServeHTTP(w, r)
// }

func (c *CartItem) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	shoppingSessionID := r.Context().Value("shopping_session_id").(int)
	productID, _ := strconv.Atoi(r.FormValue("productID"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))

	product := &model.CartItem{
		ShoppingSessionID: shoppingSessionID,
		ProductID:         productID,
		Quantity:          quantity,
	}

	_, err := cartItemHandler.Repo.CreateCartItem(product)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error fetching cart items by shopping session ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *CartItem) GetCartItemsByShoppingSessionID(w http.ResponseWriter, r *http.Request) {
	shoppingSessionID := r.Context().Value("shopping_session_id").(int)

	cartItems, err := c.Repo.GetCartItemsByShoppingSessionID(shoppingSessionID)
	if err != nil {
		log.Printf("GetCartItemsByShoppingSessionID: error fetching cart items by shopping session ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var cartProps []views.CartItemProp
	for _, ci := range cartItems {
		product, err := productHandler.Repo.GetProductByProductID(ci.ProductID)
		if err != nil {
			log.Printf("GetCartItemsByShoppingSessionID: error fetching product by product ID: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var cartProp views.CartItemProp

		cartProp.CartItemID = ci.CartItemID
		cartProp.ShoppingSessionID = ci.ShoppingSessionID
		cartProp.ProductID = ci.ProductID
		cartProp.Quantity = ci.Quantity
		cartProp.Name = product.Name
		cartProp.Price = product.Price
		cartProp.ImagePath = product.ImagePath

		cartProps = append(cartProps, cartProp)
	}

	templ.Handler(views.CartPage(cartProps)).ServeHTTP(w, r)
}

func (c *CartItem) DeleteCartItemByCartItemID(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(chi.URLParam(r, "cart_item_id"))

	err := c.Repo.DeleteCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("DeleteCartItemByCartItemID: error deleting cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (c *CartItem) IncreaseCartItemQuantityByCartItemID(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(chi.URLParam(r, "cart_item_id"))

	cartItem, err := cartItemHandler.Repo.GetCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("IncreaseCartItemQuantityByCartItemID: error fetching cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartItem.Quantity = cartItem.Quantity + 1

	err = cartItemHandler.Repo.UpdateCartItemQuantityByCartItemID(cartItem)
	if err != nil {
		log.Printf("IncreaseCartItemQuantityByCartItemID: error updating cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (c *CartItem) DecreaseCartItemQuantityByCartItemID(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(chi.URLParam(r, "cart_item_id"))

	cartItem, err := cartItemHandler.Repo.GetCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("DecreaseCartItemQuantityByCartItemID: error fetching cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartItem.Quantity = cartItem.Quantity - 1

	err = cartItemHandler.Repo.UpdateCartItemQuantityByCartItemID(cartItem)
	if err != nil {
		log.Printf("DecreaseCartItemQuantityByCartItemID: error updating cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (c *CartItem) UpdateCartItemQuantityByCartItemID(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(chi.URLParam(r, "cart_item_id"))
	// quantity, _ := strconv.Atoi(r.URL.Query().Get("quantity"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))

	cartItem, err := cartItemHandler.Repo.GetCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("UpdateCartItemQuantityByCartItemID: error fetching cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartItem.Quantity = quantity

	err = cartItemHandler.Repo.UpdateCartItemQuantityByCartItemID(cartItem)
	if err != nil {
		log.Printf("UpdateCartItemQuantityByCartItemID: error updating cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

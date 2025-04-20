package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/cart_item"
	"github.com/devkaare/web-store/repository/product"
	"github.com/devkaare/web-store/repository/shopping_session"
	"github.com/devkaare/web-store/views"
)

type CartItem struct {
	CartItemRepo        *cartitem.Repo
	ShoppingSessionRepo *shoppingsession.Repo
	ProductRepo         *product.Repo
}

var cartItemHandler = &CartItem{
	CartItemRepo:        &cartitem.Repo{},
	ShoppingSessionRepo: &shoppingsession.Repo{},
	ProductRepo:         &product.Repo{},
}

func NewCartItemHandler(db *sql.DB) *CartItem {
	cartItemHandler.CartItemRepo = repository.GetCartItem(func() *sql.DB {
		return db
	})
	cartItemHandler.ShoppingSessionRepo = repository.GetShoppingSession(func() *sql.DB {
		return db
	})
	cartItemHandler.ProductRepo = repository.GetProduct(func() *sql.DB {
		return db
	})
	return cartItemHandler
}

// func (c *CartItem) GetAllCartItems(w http.ResponseWriter, r *http.Request) {
// 	cartItems, err := c.CartItemRepo.GetAllCartItems()
// 	if err != nil {
// 		log.Printf("GetAllCartItems: error fetching cart items: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	// templ.Handler(views.CartPage(cartItems)).ServeHTTP(w, r)
// }

func (c *CartItem) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	shoppingSessionID := r.Context().Value("shopping_session_id").(string)
	productID, _ := strconv.Atoi(r.FormValue("product_id"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	// log.Printf("Found shopping session id: %d, product id: %d, quantity: %d\n", shoppingSessionID, productID, quantity)

	if quantity < 1 {
		return
	}

	item := &model.CartItem{
		ShoppingSessionID: shoppingSessionID,
		ProductID:         productID,
		Quantity:          quantity,
	}

	_, err := c.CartItemRepo.CreateCartItem(item)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error creating cart item: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := c.ProductRepo.GetProductByProductID(productID)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error fetching product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shoppingSession, err := c.ShoppingSessionRepo.GetShoppingSessionByShoppingSessionID(shoppingSessionID)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error fetching shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i := 0; i < item.Quantity; i++ {
		shoppingSession.Total = shoppingSession.Total + product.Price
	}

	err = c.ShoppingSessionRepo.UpdateShoppingSessionByShoppinSessionID(shoppingSession)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error updating shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/products/%d", productID), http.StatusSeeOther)
}

func (c *CartItem) GetCartItemsByShoppingSessionID(w http.ResponseWriter, r *http.Request) {
	shoppingSessionID := r.Context().Value("shopping_session_id").(string)

	shoppingSession, err := c.ShoppingSessionRepo.GetShoppingSessionByShoppingSessionID(shoppingSessionID)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error fetching shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var itemProps []views.CartItemProp

	items, err := c.CartItemRepo.GetCartItemsByShoppingSessionID(shoppingSessionID)
	if err == sql.ErrNoRows {
		templ.Handler(views.CartPage(itemProps, shoppingSession.Total)).ServeHTTP(w, r)
		return
	}
	if err != nil && err != sql.ErrNoRows {
		log.Printf("GetCartItemsByShoppingSessionID: error fetching cart items by shopping session ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, ci := range items {
		product, err := c.ProductRepo.GetProductByProductID(ci.ProductID)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("GetCartItemsByShoppingSessionID: error fetching product by product ID: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var itemProp views.CartItemProp

		itemProp.CartItemID = ci.CartItemID
		itemProp.ShoppingSessionID = ci.ShoppingSessionID
		itemProp.ProductID = ci.ProductID
		itemProp.Quantity = ci.Quantity
		itemProp.Name = product.Name
		itemProp.Price = product.Price
		itemProp.ImagePath = product.ImagePath

		itemProps = append(itemProps, itemProp)
	}

	templ.Handler(views.CartPage(itemProps, shoppingSession.Total)).ServeHTTP(w, r)
}

func (c *CartItem) DeleteCartItemByCartItemID(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(r.FormValue("cart_item_id"))
	shoppingSessionID := r.Context().Value("shopping_session_id").(string)

	item, err := c.CartItemRepo.GetCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("DeleteCartItemByCartItemID: error fetching cart item: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := c.ProductRepo.GetProductByProductID(item.ProductID)
	if err != nil {
		log.Printf("DeleteCartItemByCartItemID: error fetching product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shoppingSession, err := c.ShoppingSessionRepo.GetShoppingSessionByShoppingSessionID(shoppingSessionID)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error fetching shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i := 0; i < item.Quantity; i++ {
		shoppingSession.Total = shoppingSession.Total - product.Price
	}

	err = c.ShoppingSessionRepo.UpdateShoppingSessionByShoppinSessionID(shoppingSession)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error updating shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.CartItemRepo.DeleteCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("DeleteCartItemByCartItemID: error deleting cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (c *CartItem) IncreaseCartItemQuantityByCartItemID(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(r.FormValue("cart_item_id"))
	shoppingSessionID := r.Context().Value("shopping_session_id").(string)

	item, err := c.CartItemRepo.GetCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("IncreaseCartItemQuantityByCartItemID: error fetching cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	item.Quantity = item.Quantity + 1

	err = c.CartItemRepo.UpdateCartItemQuantityByCartItemID(item)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("IncreaseCartItemQuantityByCartItemID: error updating cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	existingItem, err := c.CartItemRepo.GetCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("DeleteCartItemByCartItemID: error fetching cart item: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := c.ProductRepo.GetProductByProductID(existingItem.ProductID)
	if err != nil {
		log.Printf("DeleteCartItemByCartItemID: error fetching product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shoppingSession, err := c.ShoppingSessionRepo.GetShoppingSessionByShoppingSessionID(shoppingSessionID)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error fetching shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shoppingSession.Total = shoppingSession.Total + product.Price

	err = c.ShoppingSessionRepo.UpdateShoppingSessionByShoppinSessionID(shoppingSession)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error updating shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (c *CartItem) DecreaseCartItemQuantityByCartItemID(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(r.FormValue("cart_item_id"))
	shoppingSessionID := r.Context().Value("shopping_session_id").(string)

	item, err := c.CartItemRepo.GetCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("DecreaseCartItemQuantityByCartItemID: error fetching cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	item.Quantity = item.Quantity - 1

	if item.Quantity < 1 {
		return
	}

	err = c.CartItemRepo.UpdateCartItemQuantityByCartItemID(item)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("DecreaseCartItemQuantityByCartItemID: error updating cart item by cart item ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	existingItem, err := c.CartItemRepo.GetCartItemByCartItemID(cartItemID)
	if err != nil {
		log.Printf("DeleteCartItemByCartItemID: error fetching cart item: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := c.ProductRepo.GetProductByProductID(existingItem.ProductID)
	if err != nil {
		log.Printf("DeleteCartItemByCartItemID: error fetching product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shoppingSession, err := c.ShoppingSessionRepo.GetShoppingSessionByShoppingSessionID(shoppingSessionID)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error fetching shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shoppingSession.Total = shoppingSession.Total - product.Price

	err = c.ShoppingSessionRepo.UpdateShoppingSessionByShoppinSessionID(shoppingSession)
	if err != nil {
		log.Printf("CreateCartItemByShoppingSessionID: error updating shopping session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

// func (c *CartItem) UpdateCartItemQuantityByCartItemID(w http.ResponseWriter, r *http.Request) {
// 	cartItemID, _ := strconv.Atoi(r.URL.Query().Get("cart_item_id"))
// 	quantity, _ := strconv.Atoi(r.URL.Query().Get("quantity"))
//
// 	cartItem, err := c.CartItemRepo.GetCartItemByCartItemID(cartItemID)
// 	if err != nil {
// 		log.Printf("UpdateCartItemQuantityByCartItemID: error fetching cart item by cart item ID: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	cartItem.Quantity = quantity
//
// 	err = c.CartItemRepo.UpdateCartItemQuantityByCartItemID(cartItem)
// 	if err != nil {
// 		log.Printf("UpdateCartItemQuantityByCartItemID: error updating cart item by cart item ID: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	fmt.Fprintf(w, "<p id=\"quantity-%d\">%d</p>", cartItem.CartItemID, cartItem.Quantity)
// 	return
// }

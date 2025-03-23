package repository

import (
	"database/sql"

	"github.com/devkaare/web-store/repository/cart_item"
	"github.com/devkaare/web-store/repository/order"
	"github.com/devkaare/web-store/repository/product"
	"github.com/devkaare/web-store/repository/session"
	"github.com/devkaare/web-store/repository/shopping_session"
	"github.com/devkaare/web-store/repository/user"
	"github.com/devkaare/web-store/repository/utils"
)

type Repository struct {
	Client *sql.DB
}

func GetUtils(utilsGetter func() *sql.DB) *utils.Repo {
	db := utilsGetter()
	return &utils.Repo{Client: db}
}

func GetUser(userGetter func() *sql.DB) *user.Repo {
	db := userGetter()
	return &user.Repo{Client: db}
}

func GetSession(sessionGetter func() *sql.DB) *session.Repo {
	db := sessionGetter()
	return &session.Repo{Client: db}
}

func GetProduct(productGetter func() *sql.DB) *product.Repo {
	db := productGetter()
	return &product.Repo{Client: db}
}

func GetShoppingSession(shoppingSessionGetter func() *sql.DB) *shoppingsession.Repo {
	db := shoppingSessionGetter()
	return &shoppingsession.Repo{Client: db}
}

func GetCartItem(cartItemGetter func() *sql.DB) *cartitem.Repo {
	db := cartItemGetter()
	return &cartitem.Repo{Client: db}
}

func GetOrder(orderGetter func() *sql.DB) *order.Repo {
	db := orderGetter()
	return &order.Repo{Client: db}
}

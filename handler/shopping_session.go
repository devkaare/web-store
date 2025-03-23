package handler

import (
	"database/sql"

	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/shopping_session"
)

type ShoppingSession struct {
	Repo *shoppingsession.Repo
}

var shoppingSessionHandler = &ShoppingSession{
	Repo: &shoppingsession.Repo{},
}

func NewShoppingSessionHandler(db *sql.DB) *ShoppingSession {
	shoppingSessionHandler.Repo = repository.GetShoppingSession(func() *sql.DB {
		return db
	})
	return shoppingSessionHandler
}

package handler

import (
	"database/sql"

	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/order"
)

type Order struct {
	Repo *order.Repo
}

var orderHandler = &Order{
	Repo: &order.Repo{},
}

func NewOrderHandler(db *sql.DB) *Order {
	orderHandler.Repo = repository.GetOrder(func() *sql.DB {
		return db
	})
	return orderHandler
}

package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devkaare/web-store/repository/query"
)

type Product struct {
	Repo *query.PostgresRepo
}

func (p Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	users, err := p.Repo.GetProducts()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, users)
}

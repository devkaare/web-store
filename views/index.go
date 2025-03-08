package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
)

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/products/listings?page=%d", page))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var products []model.Product

	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&products); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(products) < 1 {
		w.Write([]byte("No more products available!"))
		return
	}

	templ.Handler(index(page, len(products), products)).ServeHTTP(w, r)
}

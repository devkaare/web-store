package views

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
)

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:3000/products/listings?page=1")
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

	templ.Handler(index(products)).ServeHTTP(w, r)
}

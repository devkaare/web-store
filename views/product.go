package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
	"github.com/go-chi/chi/v5"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "ID"))
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/products/%d", productID))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var result model.Product

	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&result); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ.Handler(product(&result)).ServeHTTP(w, r)
}

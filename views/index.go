package views

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/views/components"
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

	templ.Handler(index(page, len(products), products)).ServeHTTP(w, r)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	if search == "" {
		return
	}

	rawData := url.Values{}

	rawData.Add("search", search)

	urlStr := fmt.Sprintf("http://localhost:3000/products/search%v", rawData.Encode())

	req, err := http.NewRequest("POST", urlStr, nil)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var products []model.Product

	d := json.NewDecoder(req.Response.Request.Body)
	if err := d.Decode(&products); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(products) < 1 {
		return
	}

	searchResults := components.SearchResults(products)
	searchResults.Render(context.Background(), w)
}

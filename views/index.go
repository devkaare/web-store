package views

import (
	"encoding/json"
	"fmt"
	"io"
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

	result := resp.Body
	data, err := io.ReadAll(result)
	fmt.Println(string(data))

	var resProducts []model.Product
	if err := json.Unmarshal(data, &resProducts); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ.Handler(index(resProducts)).ServeHTTP(w, r)
}

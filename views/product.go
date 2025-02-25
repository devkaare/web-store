package views

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
	"github.com/go-chi/chi/v5"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	_, productID := strconv.Atoi(chi.URLParam(r, "ID"))
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/products/%d", productID))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := resp.Body
	data, err := io.ReadAll(result)
	fmt.Println(string(data))

	resProduct := &model.Product{}
	if err := json.Unmarshal(data, resProduct); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ.Handler(product(resProduct)).ServeHTTP(w, r)
}

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
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/products/%d", page))
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

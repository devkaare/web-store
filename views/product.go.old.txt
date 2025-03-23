package views

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
//
// 	"github.com/a-h/templ"
// 	"github.com/devkaare/web-store/model"
// 	"github.com/go-chi/chi/v5"
// )
//
// func ProductHandler(w http.ResponseWriter, r *http.Request) {
// 	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))
// 	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/products/%d", productID))
// 	if err != nil {
// 		log.Fatal(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	var rawProduct model.Product
//
// 	d := json.NewDecoder(resp.Body)
// 	if err := d.Decode(&rawProduct); err != nil {
// 		log.Fatal(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	sizes := getProductSizes(&rawProduct)
//
// 	productProp := &productProp{
// 		ProductID: rawProduct.ProductID,
// 		Name:      rawProduct.Name,
// 		Price:     rawProduct.Price,
// 		Sizes:     sizes,
// 		ImagePath: rawProduct.ImagePath,
// 	}
//
// 	templ.Handler(product(productProp)).ServeHTTP(w, r)
// }

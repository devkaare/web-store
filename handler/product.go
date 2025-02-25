package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository/query"
	"github.com/go-chi/chi/v5"
)

type Product struct {
	Repo *query.PostgresRepo
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := p.Repo.GetProducts()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(products)
	_, _ = w.Write(jsonResp)
}

// curl -X POST localhost:3000/products -d "name=shirt&price=10&sizes=s,m,l,xl&imagePath=shirt.png"
func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	sizes := r.FormValue("sizes")
	imagePath := r.FormValue("imagePath")
	price, _ := strconv.Atoi(r.FormValue("price"))

	product := &model.Product{
		Name:      name,
		Price:     uint32(price),
		Sizes:     sizes,
		ImagePath: imagePath,
	}

	if _, err := p.Repo.CreateProduct(product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *Product) GetProductsByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	product, ok, err := p.Repo.GetProductByProductID(uint32(productID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		fmt.Fprintf(w, "product with product_id: %d does not exist", productID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(product)
	_, _ = w.Write(jsonResp)
}

func (p *Product) GetProductsByPage(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	products, err := p.Repo.GetProductsByPage(uint32(page))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(products)
	_, _ = w.Write(jsonResp)
}

func (p *Product) DeleteProductByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	_, ok, err := p.Repo.GetProductByProductID(uint32(productID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		fmt.Fprintf(w, "product with product_id: %d does not exist", productID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := p.Repo.DeleteProductByProductID(uint32(productID)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// curl -X PUT localhost:3000/products/1 -d "name=updatedShirt&price=12&sizes=s,m,l,xl&imagePath=updated_shirt.png"
func (p *Product) UpdateProductByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	_, ok, err := p.Repo.GetProductByProductID(uint32(productID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		fmt.Fprintf(w, "product with product_id: %d does not exist", productID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	name := r.FormValue("name")
	sizes := r.FormValue("sizes")
	imagePath := r.FormValue("imagePath")
	price, _ := strconv.Atoi(r.FormValue("price"))

	product := &model.Product{
		ProductID: uint32(productID),
		Name:      name,
		Price:     uint32(price),
		Sizes:     sizes,
		ImagePath: imagePath,
	}

	if err := p.Repo.UpdateProductByProductID(product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

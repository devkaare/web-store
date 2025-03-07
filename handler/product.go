package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository/product"
	"github.com/go-chi/chi/v5"
)

type Product struct {
	ProductRepo *product.ProductRepo
}

func (p *Product) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := p.ProductRepo.GetAllProducts()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(products)
	_, _ = w.Write(jsonResp)
}

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	reqApiKey := r.URL.Query().Get("api_key")
	if reqApiKey != apiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	name := r.FormValue("name")
	sizes := r.FormValue("sizes")
	imagePath := r.FormValue("image_path")
	price, _ := strconv.Atoi(r.FormValue("price"))

	product := &model.Product{
		Name:      name,
		Price:     price,
		Sizes:     sizes,
		ImagePath: imagePath,
	}

	productID, err := p.ProductRepo.CreateProduct(product)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product.ProductID = productID

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(product)
	_, _ = w.Write(jsonResp)
}

func (p *Product) GetProductsByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	product, err := p.ProductRepo.GetProductByProductID(productID)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(product)
	_, _ = w.Write(jsonResp)
}

func (p *Product) GetProductsByPage(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	products, err := p.ProductRepo.GetProductsByPage(page)
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
	reqApiKey := r.URL.Query().Get("api_key")
	if reqApiKey != apiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if _, err := p.ProductRepo.GetProductByProductID(productID); err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := p.ProductRepo.DeleteProductByProductID(productID); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *Product) UpdateProductByProductID(w http.ResponseWriter, r *http.Request) {
	reqApiKey := r.URL.Query().Get("api_key")
	if reqApiKey != apiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if _, err := p.ProductRepo.GetProductByProductID(productID); err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	sizes := r.FormValue("sizes")
	imagePath := r.FormValue("image_path")
	price, _ := strconv.Atoi(r.FormValue("price"))

	product := &model.Product{
		ProductID: productID,
		Name:      name,
		Price:     price,
		Sizes:     sizes,
		ImagePath: imagePath,
	}

	if err := p.ProductRepo.UpdateProductByProductID(product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

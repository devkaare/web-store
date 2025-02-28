package handler

import (
	"database/sql"
	"encoding/json"
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

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	sizes := r.FormValue("sizes")
	imagePath := r.FormValue("image_path")
	price, _ := strconv.Atoi(r.FormValue("price"))

	product := &model.Product{
		Name:      name,
		Price:     uint32(price),
		Sizes:     sizes,
		ImagePath: imagePath,
	}

	productID, err := p.Repo.CreateProduct(product)
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
	productID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	product, err := p.Repo.GetProductByProductID(uint32(productID))
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

	products, err := p.Repo.GetProductsByPage(page)
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

	if _, err := p.Repo.GetProductByProductID(uint32(productID)); err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := p.Repo.DeleteProductByProductID(uint32(productID)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *Product) UpdateProductByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	if _, err := p.Repo.GetProductByProductID(uint32(productID)); err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	sizes := r.FormValue("sizes")
	imagePath := r.FormValue("image_path")
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

package handler

import (
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

	fmt.Fprintln(w, products)
}

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	sizes := r.FormValue("sizes")
	imagePath := r.FormValue("imagePath")
	rawPrice := r.FormValue("price")

	price, err := strconv.Atoi(rawPrice)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	URLParam := chi.URLParam(r, "ID")
	productID, err := strconv.Atoi(URLParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	fmt.Fprintln(w, product)
}

func (p *Product) GetProductsByPage(w http.ResponseWriter, r *http.Request) {
	rawPage := r.URL.Query().Get("page")
	page, err := strconv.Atoi(rawPage)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	products, err := p.Repo.GetProductsByPage(uint32(page))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, products)
}

func (p *Product) DeleteProductByProductID(w http.ResponseWriter, r *http.Request) {
	URLParam := chi.URLParam(r, "ID")
	productID, err := strconv.Atoi(URLParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

func (p *Product) UpdateProductByProductID(w http.ResponseWriter, r *http.Request) {
	URLParam := chi.URLParam(r, "ID")
	productID, err := strconv.Atoi(URLParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	rawPrice := r.FormValue("price")

	price, err := strconv.Atoi(rawPrice)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

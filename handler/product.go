package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/product"
	"github.com/go-chi/chi/v5"
)

type Product struct {
	Repo *product.Repo
}

var productHandler = &Product{
	Repo: &product.Repo{},
}

func NewProductHandler(db *sql.DB) *Product {
	productHandler.Repo = repository.GetProduct(func() *sql.DB {
		return db
	})
	return productHandler
}

func (p *Product) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := p.Repo.GetAllProducts()
	if check(err) {
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
	price, _ := strconv.Atoi(r.FormValue("price"))

	file, _, _ := r.FormFile("image")
	defer file.Close()

	imagePath := filepath.Join("../views/assets/product-imgs/", name, ".png")

	dst, err := os.Create(imagePath)
	check(err)

	defer dst.Close()

	_, err = io.Copy(dst, file)
	check(err)

	product := &model.Product{
		Name:      name,
		Price:     price,
		Sizes:     sizes,
		ImagePath: imagePath,
	}

	_, err = p.Repo.CreateProduct(product)
	check(err)

	/* product.ProductID = productID

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(product)
	_, _ = w.Write(jsonResp) */

	w.Write([]byte("<p>Successfully created product!</p>"))
}

func (p *Product) GetProductsByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	product, err := p.Repo.GetProductByProductID(productID)
	if checkIfNotErrNoRows(err) {
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
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(products)
	_, _ = w.Write(jsonResp)
}

func (p *Product) GetProductsBySearch(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	if search == "" {
		return
	}

	products, err := p.Repo.GetProductsBySearch(search)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(products)
	_, _ = w.Write(jsonResp)
}

func (p *Product) DeleteProductByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	// productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := p.Repo.GetProductByProductID(productID)
	if checkIfNotErrNoRows(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = p.Repo.DeleteProductByProductID(productID)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *Product) UpdateProductByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := p.Repo.GetProductByProductID(productID)
	if checkIfNotErrNoRows(err) {
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

	err = p.Repo.UpdateProductByProductID(product)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

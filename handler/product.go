package handler

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/product"
	"github.com/devkaare/web-store/slug"
	"github.com/devkaare/web-store/views"
	"github.com/devkaare/web-store/views/components"
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
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	products, err := p.Repo.GetAllProducts(page)
	if err != nil {
		log.Printf("GetAllProducts: error fetching listingProps: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ.Handler(views.ProductListingsPage(products, page, len(products))).ServeHTTP(w, r)
}

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	productName := r.FormValue("product_name")
	categoryName := r.FormValue("category_name")
	description := r.FormValue("description")
	price, _ := strconv.Atoi(r.FormValue("price"))

	file, _, _ := r.FormFile("image")
	defer file.Close()

	imagePath := fmt.Sprintf("%s.png", slug.CreateSlug(productName))

	workingDir, _ := os.Getwd()

	dst, err := os.Create(filepath.Join(workingDir, "views/assets/product-imgs", imagePath))
	if err != nil {
		log.Printf("CreateProduct: error creating file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		log.Printf("CreateProduct: error copying file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	category := &model.Category{
		Name: categoryName,
	}

	categoryID, err := p.Repo.CreateCategory(category)
	if err != nil {
		log.Printf("CreateProduct: error creating category: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product := &model.Product{
		CategoryID:  categoryID,
		Name:        productName,
		Description: description,
		Price:       price,
		ImagePath:   imagePath,
	}

	_, err = p.Repo.CreateProduct(product)
	if err != nil {
		log.Printf("CreateProduct: error creating product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("<p>Successfully created product!</p>"))
}

func (p *Product) GetProductByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "product_id"))

	product, err := p.Repo.GetProductByProductID(productID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Product doesn't exist</p>"))
		return
	}
	if err != nil && err != sql.ErrNoRows {
		log.Printf("GetProductsByProductID: error fetching product by product ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ.Handler(views.ProductPage(product)).ServeHTTP(w, r)
}

func (p *Product) GetProductsBySearch(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	if search == "" {
		return
	}

	products, err := p.Repo.GetProductsBySearch(search)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Couldn't find the product you were looking for</p>"))
		return
	}
	if err != nil && err != sql.ErrNoRows {
		log.Printf("GetProductsBySearch: error searching for products: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	searchResults := components.SearchResults(products)
	searchResults.Render(context.Background(), w)
}

func (p *Product) GetProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID, _ := strconv.Atoi(chi.URLParam(r, "category_id"))
	if categoryID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Category doesn't exist</p>"))
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	products, err := p.Repo.GetProductsByCategoryID(categoryID, page)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil && err != sql.ErrNoRows {
		log.Printf("GetProductsByCategoryID: error fetching products: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ.Handler(views.ProductListingsPage(products, page, len(products))).ServeHTTP(w, r)
}

func (p *Product) DeleteProductByProductID(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "product_id"))

	_, err := p.Repo.GetProductByProductID(productID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("DeleteProductByProductID: error fetching product by product ID: %v", err)
		return
	}

	err = p.Repo.DeleteProductByProductID(productID)
	if err != nil {
		log.Printf("DeleteProductByProductID: error deleting product by product ID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("<p>Successfully deleted product!</p>"))
}

// func (p *Product) UpdateProductByProductID(w http.ResponseWriter, r *http.Request) {
// 	productID, _ := strconv.Atoi(chi.URLParam(r, "product_id"))
//
// 	_, err := p.Repo.GetProductByProductID(productID)
// 	if checkIfNotErrNoRows(err) {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	productName := r.FormValue("product_name")
// 	imagePath := r.FormValue("image_path")
// 	price, _ := strconv.Atoi(r.FormValue("price"))
//
// 	product := &model.Product{
// 		ProductID: productID,
// 		Name:      productName,
// 		Price:     price,
// 		ImagePath: imagePath,
// 	}
//
// 	err = p.Repo.UpdateProductByProductID(product)
//      if err == sql.ErrNoRows {
//      	w.WriteHeader(http.StatusBadRequest)
//		return
//      }
//      if err != nil && err != sql.ErrNoRows {
//      	log.Printf("UpdateProductByProductID: error updating product: %v", err)
//      	w.WriteHeader(http.StatusInternalServerError)
//      	return
//      }
// }

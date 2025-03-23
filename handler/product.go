package handler

import (
	"database/sql"
	"encoding/json"
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
	"github.com/devkaare/web-store/views"
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

func getListingProps() ([]views.ListingProp, error) {
	var listingProps []views.ListingProp

	categories, err := productHandler.Repo.GetAllCategories()
	if err != nil {
		return listingProps, fmt.Errorf("getListingProps: error fetching categories: %v", err)
	}

	for _, category := range categories {
		products, err := productHandler.Repo.GetProductsByCategoryID(category.CategoryID)
		if err != nil {
			return listingProps, fmt.Errorf("getListingProps: error fetching products by categoryID: %v", err)
		}

		for _, product := range products {
			var listingProp views.ListingProp

			listingProp.CategoryID = product.CategoryID
			listingProp.Name = product.Name
			listingProp.Price = product.Price
			listingProp.ImagePath = product.ImagePath

			listingProps = append(listingProps, listingProp)
		}
	}

	return listingProps, nil
}

func (p *Product) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	categoryID, _ := strconv.Atoi(chi.URLParam(r, "category_id"))
	if categoryID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	listingProps, err := getListingProps()
	if err != nil {
		log.Printf("GetAllProducts: error fetching listingProps: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ.Handler(views.ProductListingsPage(listingProps)).ServeHTTP(w, r)
}

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	productName := r.FormValue("product_name")
	categoryName := r.FormValue("category_name")
	description := r.FormValue("description")
	size := r.FormValue("size")
	color := r.FormValue("color")
	price, _ := strconv.Atoi(r.FormValue("price"))

	file, _, _ := r.FormFile("image")
	defer file.Close()

	imagePath := filepath.Join("../views/assets/product-imgs/", productName, ".png")

	dst, err := os.Create(imagePath)
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

	categoryID, err := productHandler.Repo.CreateCategory(category)
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
		Size:        size,
		Color:       color,
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
	if err != nil {
		log.Printf("GetProductsByProductID: error fetching product by productID: %v", err)
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
	size := r.FormValue("sizes")
	imagePath := r.FormValue("image_path")
	price, _ := strconv.Atoi(r.FormValue("price"))

	product := &model.Product{
		ProductID: productID,
		Name:      name,
		Price:     price,
		Size:      size,
		ImagePath: imagePath,
	}

	err = p.Repo.UpdateProductByProductID(product)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

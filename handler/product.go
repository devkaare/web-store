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
	users, err := p.Repo.GetProducts()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, users)
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

func (p *Product) GetProductByProductID(w http.ResponseWriter, r *http.Request) {
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

func (u *Product) DeleteProductByProductID(w http.ResponseWriter, r *http.Request) {
	URLParam := chi.URLParam(r, "ID")
	productID, err := strconv.Atoi(URLParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, ok, err := u.Repo.GetProductByProductID(uint32(productID))
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

	if err := u.Repo.DeleteProductByProductID(uint32(productID)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// func (u *User) UpdateUserByUserID(w http.ResponseWriter, r *http.Request) {
// 	URLParam := chi.URLParam(r, "ID")
// 	userID, err := strconv.Atoi(URLParam)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	_, ok, err := u.Repo.GetUserByUserID(uint32(userID))
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if !ok {
// 		fmt.Fprintf(w, "user with user_id: %d does not exist", userID)
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
//
// 	email := r.FormValue("email")
// 	password := r.FormValue("password")
//
// 	user := &model.User{
// 		UserID:   uint32(userID),
// 		Email:    email,
// 		Password: password,
// 	}
//
// 	if err := u.Repo.UpdateUserByUserID(user); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }

package handler

import (
	"net/http"

	"github.com/a-h/templ"
	// "github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/views"
)

type Store struct{}

func (s *Store) SignInPageHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.SignInPage()).ServeHTTP(w, r)
}

func (s *Store) SignUpPageHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.SignUpPage()).ServeHTTP(w, r)
}

func (s *Store) CartPageHandler(w http.ResponseWriter, r *http.Request) {
	// templ.Handler(views.CartPage([]views.CartProp{})).ServeHTTP(w, r)
}

func (s *Store) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// templ.Handler(views.HomePage(0, 0, []model.Product{})).ServeHTTP(w, r)
}

func (s *Store) ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	// templ.Handler(views.ProductPage(&views.ProductProp{})).ServeHTTP(w, r)
}

func (s *Store) SuccessPageHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.SuccessPage()).ServeHTTP(w, r)
}

func (s *Store) CancelPageHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.CancelPage()).ServeHTTP(w, r)
}

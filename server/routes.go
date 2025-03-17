package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.FS(views.Files))
	r.Handle("/assets/*", fileServer)

	r.Route("/utils", s.registerUtilsRoutes)
	r.Route("/users", s.registerUserRoutes)
	r.Route("/products", s.registerProductRoutes)
	r.Route("/carts", s.registerCartRoutes)
	r.Route("/sessions", s.registerSessionRoutes)
	r.Route("/auth", s.registerAuthRoutes)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/listings", http.StatusSeeOther)
	})

	r.Get("/signin", func(w http.ResponseWriter, r *http.Request) { templ.Handler(views.SignInPage()).ServeHTTP(w, r) })
	r.Get("/signup", func(w http.ResponseWriter, r *http.Request) { templ.Handler(views.SignUpPage()).ServeHTTP(w, r) })

	r.Get("/cart", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.CartPage([]views.CartProp{})).ServeHTTP(w, r)
	})

	r.Get("/listings", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.IndexPage(0, 0, []model.Product{})).ServeHTTP(w, r)
	})
	r.Get("/listings/{id}", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.ProductPage(&views.ProductProp{})).ServeHTTP(w, r)
	})

	return r
}

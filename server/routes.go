package server

import (
	"net/http"

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

	r.Post("/signup", views.SignUpHandler)
	r.Get("/signin", views.SignInHandler)
	r.Get("/search", views.SearchHandler)
	r.Get("/cart", views.CartHandler)
	r.Get("/listings", views.IndexPageHandler)
	r.Get("/listings/{id}", views.ProductHandler)

	return r
}

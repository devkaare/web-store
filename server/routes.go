package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(100, time.Minute))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/utils", s.registerUtilsRoutes)
	r.Route("/users", s.registerUserRoutes)
	r.Route("/products", s.registerProductRoutes)
	r.Route("/cart", s.registerCartRoutes)
	r.Route("/sessions", s.registerSessionRoutes)
	r.Route("/auth", s.registerAuthRoutes)
	r.Route("/", s.registerRoutes)

	fileServer := http.StripPrefix("/assets/", http.FileServer(http.Dir("./views/assets")))
	r.Handle("/assets/*", fileServer)
	// fileServer := http.FileServer(http.FS(views.Files))
	// r.Handle("/assets/*", fileServer)

	return r
}

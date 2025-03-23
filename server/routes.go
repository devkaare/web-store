package server

import (
	"net/http"

	"github.com/a-h/templ"
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

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/listings", http.StatusSeeOther)
	// })

	// r.Route("/", s.registerStoreRoutes)
	// r.Route("/utils", s.registerUtilsRoutes)
	// r.Route("/users", s.registerUserRoutes)
	// r.Route("/products", s.registerProductRoutes)
	// r.Route("/carts", s.registerCartRoutes)
	// r.Route("/sessions", s.registerSessionRoutes)

	r.Handle("/", templ.Handler(views.IndexPage()))
	r.Handle("/signup", templ.Handler(views.SignUpPage()))
	r.Handle("/signin", templ.Handler(views.SignInPage()))

	r.Route("/auth", s.registerAuthRoutes)

	return r
}

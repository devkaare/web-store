package server

import (
	"net/http"

	// "github.com/devkaare/todo/handler"
	// "github.com/devkaare/todo/repository/todo"
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
	r.Get("/", templ.Handler(views.Base()).ServeHTTP)
	r.Route("/user", s.RegisterUserRoutes)

	return r
}

func (s *Server) RegisterUserRoutes(r chi.Router) {
	// todoHandler := &handler.Todo{
	// 	Repo: &todo.PostgresRepo{
	// 		Client: s.db,
	// 	},
	// }
	//
	// r.Get("/health", todoHandler.Health)
	// r.Post("/", todoHandler.Create)
	// r.Get("/", todoHandler.List)
	// r.Get("/{ID}", todoHandler.GetByID)
	// r.Get("/edit/{ID}", todoHandler.EditByID)
	// r.Put("/{ID}", todoHandler.UpdateByID)
	// r.Delete("/{ID}", todoHandler.DeleteByID)
}

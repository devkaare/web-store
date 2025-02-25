package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/handler"
	"github.com/devkaare/web-store/repository/query"
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

	r.Route("/utils", s.RegisterUtilsRoutes)
	r.Route("/users", s.RegisterUserRoutes)
	r.Route("/products", s.RegisterProductRoutes)
	r.Route("/cart", s.RegisterCartRoutes)

	return r
}

func (s *Server) RegisterUtilsRoutes(r chi.Router) {
	utilsHandler := &handler.Utils{
		Repo: &query.PostgresRepo{
			Client: s.db,
		},
	}

	r.Get("/health", utilsHandler.Health)
}

func (s *Server) RegisterUserRoutes(r chi.Router) {
	userHandler := &handler.User{
		Repo: &query.PostgresRepo{
			Client: s.db,
		},
	}

	r.Post("/", userHandler.CreateUser)
	r.Get("/", userHandler.GetUsers)
	r.Get("/{ID}", userHandler.GetUserByUserID)
	r.Put("/{ID}", userHandler.UpdateUserByUserID)
	r.Delete("/{ID}", userHandler.DeleteUserByUserID)
}

func (s *Server) RegisterProductRoutes(r chi.Router) {
	productHandler := &handler.Product{
		Repo: &query.PostgresRepo{
			Client: s.db,
		},
	}

	r.Post("/", productHandler.CreateProduct)
	r.Get("/", productHandler.GetProducts)
	r.Get("/{ID}", productHandler.GetProductsByProductID)
	r.Put("/{ID}", productHandler.UpdateProductByProductID)
	r.Delete("/{ID}", productHandler.DeleteProductByProductID)
}

func (s *Server) RegisterCartRoutes(r chi.Router) {
	cartHandler := &handler.CartItem{
		Repo: &query.PostgresRepo{
			Client: s.db,
		},
	}

	r.Post("/", cartHandler.CreateCartItem)
	r.Get("/", cartHandler.GetCartItems)
	r.Get("/{ID}", cartHandler.GetCartItemsByUserID)
	r.Put("/{userID}/{productID}/{size}/{quantity}", cartHandler.UpdateCartItemQuantity)
	r.Delete("/{userID}/{productID}/{size}", cartHandler.DeleteCartItem)
}

package server

import (
	"net/http"

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

	r.Route("/utils", s.registerUtilsRoutes)
	r.Route("/users", s.registerUserRoutes)
	r.Route("/products", s.registerProductRoutes)
	r.Route("/carts", s.registerCartRoutes)
	r.Route("/sessions", s.registerSessionRoutes)

	r.Get("/signup", views.SignUpHandler)
	r.Get("/signin", views.SignInHandler)
	r.Get("/cart", views.CartHandler)
	r.Get("/listings", views.IndexPageHandler)
	r.Get("/listings/{ID}", views.ProductHandler)

	return r
}

func (s *Server) registerUtilsRoutes(r chi.Router) {
	utilsHandler := &handler.Utils{
		Repo: &query.PostgresRepo{
			Client: s.db,
		},
	}

	r.Get("/health", utilsHandler.Health)
}

func (s *Server) registerUserRoutes(r chi.Router) {
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

func (s *Server) registerProductRoutes(r chi.Router) {
	productHandler := &handler.Product{
		Repo: &query.PostgresRepo{
			Client: s.db,
		},
	}

	r.Post("/", productHandler.CreateProduct)
	r.Get("/", productHandler.GetProducts)
	r.Get("/{ID}", productHandler.GetProductsByProductID)
	r.Get("/listings", productHandler.GetProductsByPage)
	r.Put("/{ID}", productHandler.UpdateProductByProductID)
	r.Delete("/{ID}", productHandler.DeleteProductByProductID)
}

func (s *Server) registerCartRoutes(r chi.Router) {
	cartHandler := &handler.CartItem{
		Repo: &query.PostgresRepo{
			Client: s.db,
		},
	}

	r.Post("/", cartHandler.CreateCartItem)
	r.Get("/", cartHandler.GetCartItems)
	r.Get("/{userID}", cartHandler.GetCartItemsByUserID)
	r.Put("/{userID}/{productID}", cartHandler.UpdateCartItemQuantity)
	r.Delete("/{userID}/{productID}", cartHandler.DeleteCartItem)
}

func (s *Server) registerSessionRoutes(r chi.Router) {
	sessionHandler := &handler.Session{
		Repo: &query.PostgresRepo{
			Client: s.db,
		},
	}

	r.Post("/signup", sessionHandler.SignUp)
	r.Post("/signin", sessionHandler.SignIn)
	r.Get("/refresh", sessionHandler.Refresh)
	r.Get("/welcome", sessionHandler.Welcome)
	r.Get("/logout", sessionHandler.LogOut)
	r.Get("/", sessionHandler.GetSessions)
}

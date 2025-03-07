package server

import (
	"net/http"

	"github.com/devkaare/web-store/handler"
	"github.com/devkaare/web-store/repository/cart"
	"github.com/devkaare/web-store/repository/product"
	"github.com/devkaare/web-store/repository/session"
	"github.com/devkaare/web-store/repository/user"
	"github.com/devkaare/web-store/repository/utils"
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
	r.Get("/listings/{id}", views.ProductHandler)

	return r
}

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (s *Server) registerUtilsRoutes(r chi.Router) {
	utilsHandler := &handler.Utils{
		UtilsRepo: &utils.UtilsRepo{
			Client: s.db,
		},
	}

	r.Get("/health", utilsHandler.Health)
}

func (s *Server) registerUserRoutes(r chi.Router) {
	userHandler := &handler.User{
		UserRepo: &user.UserRepo{
			Client: s.db,
		},
	}

	r.Post("/", userHandler.CreateUser)
	r.Get("/", userHandler.GetAllUsers)
	r.Get("/{id}", userHandler.GetUserByUserID)
	r.Put("/{id}", userHandler.UpdateUserByUserID)
	r.Delete("/{id}", userHandler.DeleteUserByUserID)
}

func (s *Server) registerProductRoutes(r chi.Router) {
	productHandler := &handler.Product{
		ProductRepo: &product.ProductRepo{
			Client: s.db,
		},
	}

	r.Post("/", productHandler.CreateProduct)
	r.Get("/", productHandler.GetAllProducts)
	r.Get("/listings", productHandler.GetProductsByPage)
	r.Get("/{id}", productHandler.GetProductsByProductID)
	r.Put("/{id}", productHandler.UpdateProductByProductID)
	r.Delete("/{id}", productHandler.DeleteProductByProductID)
}

func (s *Server) registerCartRoutes(r chi.Router) {
	cartHandler := &handler.CartItem{
		CartRepo: &cart.CartRepo{
			Client: s.db,
		},
	}

	r.Use(cartHandler.CartMiddleware)

	r.Post("/", cartHandler.CreateCartItem)
	r.Get("/", cartHandler.GetAllCartItems)
	r.Get("/{user_id}", cartHandler.GetCartItemsByUserID)
	r.Put("/{user_id}/{product_id}", cartHandler.UpdateCartItemQuantity)
	r.Delete("/{user_id}/{product_id}", cartHandler.DeleteCartItem)
}

func (s *Server) registerSessionRoutes(r chi.Router) {
	sessionHandler := &handler.Session{
		SessionRepo: &session.SessionRepo{
			Client: s.db,
		},
	}

	r.Post("/signup", sessionHandler.SignUp)
	r.Post("/signin", sessionHandler.SignIn)
	r.Get("/refresh", sessionHandler.Refresh)
	r.Get("/welcome", sessionHandler.Welcome)
	r.Get("/logout", sessionHandler.LogOut)
	r.Get("/", sessionHandler.GetAllSessions)
}

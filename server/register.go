package server

import (
	"github.com/devkaare/web-store/handler"
	"github.com/go-chi/chi/v5"
)

func (s *Server) registerUtilsRoutes(r chi.Router) {
	utilsHandler := handler.NewUtilsHandler(s.db)

	r.Get("/health", utilsHandler.Health)
}

func (s *Server) registerUserRoutes(r chi.Router) {
	userHandler := handler.NewUserHandler(s.db)

	r.Post("/", userHandler.CreateUser)
	r.Get("/", userHandler.GetAllUsers)
	r.Get("/{id}", userHandler.GetUserByUserID)
	r.Put("/{id}", userHandler.UpdateUserByUserID)
	r.Delete("/{id}", userHandler.DeleteUserByUserID)
}

func (s *Server) registerProductRoutes(r chi.Router) {
	productHandler := handler.NewProductHandler(s.db)

	r.Post("/", productHandler.CreateProduct)
	r.Post("/search", productHandler.GetProductsBySearch)
	r.Get("/", productHandler.GetAllProducts)
	r.Get("/listings", productHandler.GetProductsByPage)
	r.Get("/{id}", productHandler.GetProductByProductID)
	r.Put("/{id}", productHandler.UpdateProductByProductID)
	r.Delete("/", productHandler.DeleteProductByProductID)
}

func (s *Server) registerCartRoutes(r chi.Router) {
	authHandler := &handler.Authentication{}
	r.Use(authHandler.SessionMiddleware)

	cartHandler := handler.NewCartItemHandler(s.db)

	r.Post("/", cartHandler.CreateCartItem)
	r.Get("/", cartHandler.GetAllCartItems)
	r.Get("/{user_id}", cartHandler.GetCartItemsByUserID)
	r.Put("/{user_id}/{product_id}", cartHandler.UpdateCartItemQuantity)
	r.Delete("/{user_id}/{product_id}", cartHandler.DeleteCartItem)
}

func (s *Server) registerSessionRoutes(r chi.Router) {
	sessionHandler := handler.NewSessionHandler(s.db)

	r.Get("/", sessionHandler.GetAllSessions)
	r.Get("/refresh", sessionHandler.Refresh)
	r.Get("/welcome", sessionHandler.Welcome)
	r.Get("/logout", sessionHandler.LogOut)
}

func (s *Server) registerAuthRoutes(r chi.Router) {
	authHandler := &handler.Authentication{}

	r.Post("/signup", authHandler.SignUp)
	r.Post("/signin", authHandler.SignIn)
}

func (s *Server) registerStoreRoutes(r chi.Router) {
	storeHandler := &handler.Store{}

	r.Get("/signin", storeHandler.SignInPageHandler)
	r.Get("/signup", storeHandler.SignUpPageHandler)
	r.Get("/cart", storeHandler.CartPageHandler)
	r.Get("/listings", storeHandler.HomePageHandler)
	r.Get("/listings/{id}", storeHandler.ProductPageHandler)
}

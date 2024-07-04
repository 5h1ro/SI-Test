package routes

import (
	"customer/config"
	"customer/internal/handler"
	"customer/internal/helper"
	"customer/internal/repository"
	"customer/internal/service"
	"customer/internal/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

func ApiRouter(app *chi.Mux, db *gorm.DB, cfg *config.Config, authToken *jwtauth.JWTAuth) {
	// Repository
	customerRepo := repository.NewCustomerRepository(db)
	// Service
	authService := service.NewAuthService(&types.Server{
		Router:    app,
		DB:        db,
		AuthToken: authToken,
	})
	customerService := service.NewCustomerService(customerRepo)
	// Handler
	authHandler := handler.NewAuthHandler(authService)
	customerHandler := handler.NewCustomerHandler(customerService)
	// Cors
	app.Use()

	// Protected routes
	app.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(authToken))
		r.Use(helper.Authenticator(authToken))

		r.Post("/create", customerHandler.Create)
	})

	// Public routes
	app.Group(func(r chi.Router) {
		r.Get("/login", authHandler.Login)
	})
}

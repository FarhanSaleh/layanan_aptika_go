package api

import (
	"database/sql"

	"github.com/farhansaleh/layanan_aptika_be/internal/auth"
	"github.com/farhansaleh/layanan_aptika_be/internal/instansi"
	"github.com/farhansaleh/layanan_aptika_be/internal/middlewares"
	"github.com/farhansaleh/layanan_aptika_be/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func SetupRoutes(r chi.Router, db *sql.DB) {
	validator := validator.New()

	// Repository
	usersRepository := users.NewRepository()
	instansiRepository := instansi.NewRepository()

	// Service
	usersServices := users.NewService(db, usersRepository, validator)
	authService := auth.NewService(db, usersRepository, validator)
	instansiService := instansi.NewService(db, instansiRepository, validator)
	
	// Handler
	usersHandler := users.NewHandler(usersServices)
	authHandler := auth.NewHandler(authService)
	instansiHandler := instansi.NewHandler(instansiService)
	
	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/users", usersHandler.Create)
		r.Put("/users/{id}", usersHandler.Update)
		r.Delete("/users/{id}", usersHandler.Delete)
		r.Get("/users", usersHandler.FindAll)
		r.Get("/users/{id}", usersHandler.FindById)

		r.Post("/instansi", instansiHandler.Create)
		r.Put("/instansi/{id}", instansiHandler.Update)
		r.Delete("/instansi/{id}", instansiHandler.Delete)
		r.Get("/instansi", instansiHandler.FindAll)
		r.Get("/instansi/{id}", instansiHandler.FindById)
	})

	// Public routes
	r.Post("/login", authHandler.Login)
	r.Delete("/logout", authHandler.Logout)
	r.Put("/change-password", authHandler.ChangePassword)
}
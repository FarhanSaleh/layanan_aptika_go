package api

import (
	"database/sql"

	"github.com/farhansaleh/layanan_aptika_be/internal/auth"
	"github.com/farhansaleh/layanan_aptika_be/internal/middlewares"
	"github.com/farhansaleh/layanan_aptika_be/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func SetupRoutes(r chi.Router, db *sql.DB) {
	validator := validator.New()

	usersRepository := users.NewRepository()
	usersServices := users.NewService(db, usersRepository, validator)
	authService := auth.NewService(db, usersRepository, validator)
	
	usersHandler := users.NewHandler(usersServices)
	authHandler := auth.NewHandler(authService)
	
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/users", usersHandler.Create)
		r.Put("/users/{id}", usersHandler.Update)
		r.Delete("/users/{id}", usersHandler.Delete)
		r.Get("/users", usersHandler.FindAll)
		r.Get("/users/{id}", usersHandler.FindById)
	})

	r.Post("/login", authHandler.Login)
	r.Delete("/logout", authHandler.Logout)
	r.Put("/change-password", authHandler.ChangePassword)
}
package api

import (
	"database/sql"

	"github.com/farhansaleh/layanan_aptika_be/internal/auth"
	"github.com/farhansaleh/layanan_aptika_be/internal/instansi"
	"github.com/farhansaleh/layanan_aptika_be/internal/middlewares"
	rolepengelola "github.com/farhansaleh/layanan_aptika_be/internal/role_pengelola"
	"github.com/farhansaleh/layanan_aptika_be/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func SetupRoutes(r chi.Router, db *sql.DB) {
	validator := validator.New()

	// Repository
	usersRepository := users.NewRepository()
	instansiRepository := instansi.NewRepository()
	rolePengelolaRepository := rolepengelola.NewRepository()

	// Service
	usersServices := users.NewService(db, usersRepository, validator)
	authService := auth.NewService(db, usersRepository, validator)
	instansiService := instansi.NewService(db, instansiRepository, validator)
	rolePengelolaService := rolepengelola.NewService(db, rolePengelolaRepository, validator)
	
	// Handler
	usersHandler := users.NewHandler(usersServices)
	authHandler := auth.NewHandler(authService)
	instansiHandler := instansi.NewHandler(instansiService)
	rolePengelolaHandler := rolepengelola.NewHandler(rolePengelolaService)
	
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

		r.Post("/role-pengelola", rolePengelolaHandler.Create)
		r.Get("/role-pengelola", rolePengelolaHandler.FindAll)
		r.Put("/role-pengelola/{id}", rolePengelolaHandler.Update)
		r.Delete("/role-pengelola/{id}", rolePengelolaHandler.Delete)

		r.Put("/change-password", authHandler.ChangePassword)
		r.Delete("/logout", authHandler.Logout)
	})

	// Public routes
	r.Post("/login", authHandler.Login)
}
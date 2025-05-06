package api

import (
	"database/sql"
	"net/http"
	"path/filepath"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/auth"
	gangguanjip "github.com/farhansaleh/layanan_aptika_be/internal/api/gangguan-jip"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/instansi"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/pengelola"
	rolepengelola "github.com/farhansaleh/layanan_aptika_be/internal/api/role_pengelola"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/users"
	"github.com/farhansaleh/layanan_aptika_be/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func SetupRoutes(r chi.Router, db *sql.DB, config *config.Config) {
	validator := validator.New()

	// Repository
	usersRepository := users.NewRepository()
	instansiRepository := instansi.NewRepository()
	rolePengelolaRepository := rolepengelola.NewRepository()
	pengelolaRepository := pengelola.NewRepository()
	gangguanJIPRepository := gangguanjip.NewRepository()

	// Service
	usersServices := users.NewService(db, usersRepository, validator)
	authService := auth.NewService(db, usersRepository, pengelolaRepository, validator)
	instansiService := instansi.NewService(db, instansiRepository, validator)
	rolePengelolaService := rolepengelola.NewService(db, rolePengelolaRepository, validator)
	pengelolaService := pengelola.NewService(db, pengelolaRepository, validator)
	gangguanJIPService := gangguanjip.NewService(db, gangguanJIPRepository, validator, config) 
	
	// Handler
	usersHandler := users.NewHandler(usersServices)
	authHandler := auth.NewHandler(authService)
	instansiHandler := instansi.NewHandler(instansiService)
	rolePengelolaHandler := rolepengelola.NewHandler(rolePengelolaService)
	pengelolaHandler := pengelola.NewHandler(pengelolaService)
	gangguanJIPHandler := gangguanjip.NewHandler(gangguanJIPService)
	
	// Protected routes user
	r.Group(func(r chi.Router) {
		r.Use(middlewares.UserAuthMiddleware)
		r.Put("/change-password/user", authHandler.UserChangePassword)
		r.Get("/uploads/user/img/*", func(w http.ResponseWriter, r *http.Request) {
			path := chi.URLParam(r, "*")
			filePath := filepath.Join("uploads", "img", path)
			
			http.ServeFile(w, r, filePath)
		})
		r.Get("/uploads/user/docs/*", func(w http.ResponseWriter, r *http.Request) {
			path := chi.URLParam(r, "*")
			filePath := filepath.Join("uploads", "docs", path)
			
			http.ServeFile(w, r, filePath)
		})

		r.Post("/gangguan-jip", gangguanJIPHandler.Create)
		r.Put("/gangguan-jip/{id}", gangguanJIPHandler.Update)
		r.Delete("/gangguan-jip/{id}", gangguanJIPHandler.Delete)
		r.Get("/gangguan-jip/me/{id}", gangguanJIPHandler.FindById)
		r.Get("/gangguan-jip/me", gangguanJIPHandler.FindByUser)

		r.Delete("/logout", authHandler.Logout)
	})
	
	// Protected routes pengelola
	r.Group(func(r chi.Router) {
		r.Use(middlewares.PengelolaAuthMiddleware)
		r.Get("/uploads/pengelola/img/*", func(w http.ResponseWriter, r *http.Request) {
			path := chi.URLParam(r, "*")
			filePath := filepath.Join("uploads", "img", path)
			
			http.ServeFile(w, r, filePath)
		})
		r.Get("/uploads/pengelola/docs/*", func(w http.ResponseWriter, r *http.Request) {
			path := chi.URLParam(r, "*")
			filePath := filepath.Join("uploads", "docs", path)
			
			http.ServeFile(w, r, filePath)
		})
		
		r.Post("/pengelola", pengelolaHandler.Create)
		r.Put("/pengelola/{id}", pengelolaHandler.Update)
		r.Delete("/pengelola/{id}", pengelolaHandler.Delete)
		r.Get("/pengelola", pengelolaHandler.FindAll)
		r.Get("/pengelola/{id}", pengelolaHandler.FindById)

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

		r.Get("/gangguan-jip", gangguanJIPHandler.FindAll)
		r.Get("/gangguan-jip/{id}", gangguanJIPHandler.FindById)
		r.Patch("/gangguan-jip/{id}", gangguanJIPHandler.UpdateStatus)
		
		r.Put("/change-password/pengelola", authHandler.PengelolaChangePassword)
	})

	// Public routes
	r.Post("/login/user", authHandler.Login)
	r.Post("/login/pengelola", authHandler.PengelolaLogin)
}
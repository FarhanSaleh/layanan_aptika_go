package api

import (
	"database/sql"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/auth"
	gangguanjip "github.com/farhansaleh/layanan_aptika_be/internal/api/gangguan-jip"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/instansi"
	pembangunanaplikasi "github.com/farhansaleh/layanan_aptika_be/internal/api/pembangunan_aplikasi"
	pembuatanemail "github.com/farhansaleh/layanan_aptika_be/internal/api/pembuatan_email"
	pembuatansubdomain "github.com/farhansaleh/layanan_aptika_be/internal/api/pembuatan_subdomain"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/pengelola"
	perubahanipserver "github.com/farhansaleh/layanan_aptika_be/internal/api/perubahan_ip_server"
	pusatdatadaerah "github.com/farhansaleh/layanan_aptika_be/internal/api/pusat_data_daerah"
	rolepengelola "github.com/farhansaleh/layanan_aptika_be/internal/api/role_pengelola"
	"github.com/farhansaleh/layanan_aptika_be/internal/api/static"
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
	perubahanIPServerRepository := perubahanipserver.NewRepository()
	pusatDataDaerahRepository := pusatdatadaerah.NewRepository()
	pembangunanaplikasiRepository := pembangunanaplikasi.NewRepository()
	pembuatanSubdomainRepository := pembuatansubdomain.NewRepository()
	pembuatanEmailRepository := pembuatanemail.NewRepository()

	// Service
	usersServices := users.NewService(db, usersRepository, validator)
	authService := auth.NewService(db, usersRepository, pengelolaRepository, validator)
	instansiService := instansi.NewService(db, instansiRepository, validator)
	rolePengelolaService := rolepengelola.NewService(db, rolePengelolaRepository, validator)
	pengelolaService := pengelola.NewService(db, pengelolaRepository, validator)
	gangguanJIPService := gangguanjip.NewService(db, gangguanJIPRepository, validator, config) 
	perubahanIPServerService := perubahanipserver.NewService(db, perubahanIPServerRepository, validator, config)
	pusatDataDaerahService := pusatdatadaerah.NewService(db, pusatDataDaerahRepository, validator, config)
	pembanguananAplikasiService := pembangunanaplikasi.NewService(db, pembangunanaplikasiRepository, validator, config)
	pembuatanSubdomainService := pembuatansubdomain.NewService(db, pembuatanSubdomainRepository, validator, config)
	pembuatanEmailService := pembuatanemail.NewService(db, pembuatanEmailRepository, validator, config)
	
	// Handler
	usersHandler := users.NewHandler(usersServices)
	authHandler := auth.NewHandler(authService)
	instansiHandler := instansi.NewHandler(instansiService)
	rolePengelolaHandler := rolepengelola.NewHandler(rolePengelolaService)
	pengelolaHandler := pengelola.NewHandler(pengelolaService)
	gangguanJIPHandler := gangguanjip.NewHandler(gangguanJIPService)
	perubahanIPServerHandler := perubahanipserver.NewHandler(perubahanIPServerService)
	pusatDataDaerahHandler := pusatdatadaerah.NewHandler(pusatDataDaerahService)
	pembangunanAplikasiHandler := pembangunanaplikasi.NewHandler(pembanguananAplikasiService)
	pembuatanSubdomainHandler := pembuatansubdomain.NewHandler(pembuatanSubdomainService)
	pembuatanEmailHandler := pembuatanemail.NewHandler(pembuatanEmailService)
	staticHandler := static.NewHandler()
	
	// Protected routes user
	r.Group(func(r chi.Router) {
		r.Use(middlewares.UserAuthMiddleware)
		r.Put("/change-password/user", authHandler.UserChangePassword)
		r.Get("/uploads/user/img/{filename}", staticHandler.Image)
		r.Get("/uploads/user/docs/{filename}", staticHandler.Document)

		r.Post("/gangguan-jip", gangguanJIPHandler.Create)
		r.Put("/gangguan-jip/{id}", gangguanJIPHandler.Update)
		r.Delete("/gangguan-jip/{id}", gangguanJIPHandler.Delete)
		r.Get("/gangguan-jip/me/{id}", gangguanJIPHandler.FindById)
		r.Get("/gangguan-jip/me", gangguanJIPHandler.FindByUser)
		
		r.Post("/perubahan-ip-server", perubahanIPServerHandler.Create)
		r.Put("/perubahan-ip-server/{id}", perubahanIPServerHandler.Update)
		r.Delete("/perubahan-ip-server/{id}", perubahanIPServerHandler.Delete)
		r.Get("/perubahan-ip-server/me/{id}", perubahanIPServerHandler.FindById)
		r.Get("/perubahan-ip-server/me", perubahanIPServerHandler.FindByUser)
		
		r.Post("/pusat-data-daerah", pusatDataDaerahHandler.Create)
		r.Put("/pusat-data-daerah/{id}", pusatDataDaerahHandler.Update)
		r.Delete("/pusat-data-daerah/{id}", pusatDataDaerahHandler.Delete)
		r.Get("/pusat-data-daerah/me/{id}", pusatDataDaerahHandler.FindById)
		r.Get("/pusat-data-daerah/me", pusatDataDaerahHandler.FindByUser)
		
		r.Post("/pembangunan-aplikasi", pembangunanAplikasiHandler.Create)
		r.Put("/pembangunan-aplikasi/{id}", pembangunanAplikasiHandler.Update)
		r.Delete("/pembangunan-aplikasi/{id}", pembangunanAplikasiHandler.Delete)
		r.Get("/pembangunan-aplikasi/me/{id}", pembangunanAplikasiHandler.FindById)
		r.Get("/pembangunan-aplikasi/me", pembangunanAplikasiHandler.FindByUser)
		
		r.Post("/pembuatan-subdomain", pembuatanSubdomainHandler.Create)
		r.Put("/pembuatan-subdomain/{id}", pembuatanSubdomainHandler.Update)
		r.Delete("/pembuatan-subdomain/{id}", pembuatanSubdomainHandler.Delete)
		r.Get("/pembuatan-subdomain/me/{id}", pembuatanSubdomainHandler.FindById)
		r.Get("/pembuatan-subdomain/me", pembuatanSubdomainHandler.FindByUser)
		
		r.Post("/pembuatan-email", pembuatanEmailHandler.Create)
		r.Put("/pembuatan-email/{id}", pembuatanEmailHandler.Update)
		r.Delete("/pembuatan-email/{id}", pembuatanEmailHandler.Delete)
		r.Get("/pembuatan-email/me/{id}", pembuatanEmailHandler.FindById)
		r.Get("/pembuatan-email/me", pembuatanEmailHandler.FindByUser)

		r.Delete("/logout", authHandler.Logout)
	})
	
	// Protected routes pengelola
	r.Group(func(r chi.Router) {
		r.Use(middlewares.PengelolaAuthMiddleware)
		r.Get("/uploads/pengelola/img/{filename}", staticHandler.Image)
		r.Get("/uploads/pengelola/docs/{filename}", staticHandler.Document)
		
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
		r.Get("/instansi/{id}", instansiHandler.FindById)

		r.Post("/role-pengelola", rolePengelolaHandler.Create)
		r.Get("/role-pengelola", rolePengelolaHandler.FindAll)
		r.Put("/role-pengelola/{id}", rolePengelolaHandler.Update)
		r.Delete("/role-pengelola/{id}", rolePengelolaHandler.Delete)

		r.Get("/gangguan-jip", gangguanJIPHandler.FindAll)
		r.Get("/gangguan-jip/{id}", gangguanJIPHandler.FindById)
		r.Patch("/gangguan-jip/{id}", gangguanJIPHandler.UpdateStatus)

		r.Get("/perubahan-ip-server", perubahanIPServerHandler.FindAll)
		r.Get("/perubahan-ip-server/{id}", perubahanIPServerHandler.FindById)
		r.Patch("/perubahan-ip-server/{id}", perubahanIPServerHandler.UpdateStatus)
		
		r.Get("/pusat-data-daerah", pusatDataDaerahHandler.FindAll)
		r.Get("/pusat-data-daerah/{id}", pusatDataDaerahHandler.FindById)
		r.Patch("/pusat-data-daerah/{id}", pusatDataDaerahHandler.UpdateStatus)
		
		r.Get("/pembangunan-aplikasi", pembangunanAplikasiHandler.FindAll)
		r.Get("/pembangunan-aplikasi/{id}", pembangunanAplikasiHandler.FindById)
		r.Patch("/pembangunan-aplikasi/{id}", pembangunanAplikasiHandler.UpdateStatus)
		
		r.Get("/pembuatan-subdomain", pembuatanSubdomainHandler.FindAll)
		r.Get("/pembuatan-subdomain/{id}", pembuatanSubdomainHandler.FindById)
		r.Patch("/pembuatan-subdomain/{id}", pembuatanSubdomainHandler.UpdateStatus)
		
		r.Get("/pembuatan-email", pembuatanEmailHandler.FindAll)
		r.Get("/pembuatan-email/{id}", pembuatanEmailHandler.FindById)
		r.Patch("/pembuatan-email/{id}", pembuatanEmailHandler.UpdateStatus)

		r.Put("/change-password/pengelola", authHandler.PengelolaChangePassword)
	})

	// Public routes
	r.Post("/login/user", authHandler.Login)
	r.Post("/login/pengelola", authHandler.PengelolaLogin)
	r.Get("/instansi", instansiHandler.FindAll)
}
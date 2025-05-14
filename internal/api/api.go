package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type APIServer struct {
	addr string
	config *config.Config
}

func NewAPIServer(addr string, config *config.Config) *APIServer {
	return &APIServer{
		addr: addr,
		config: config,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: s.config.AllowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Heartbeat("/ping"))

	db, err := config.NewDB()
	if err != nil {
		return err
	}
	initStorage(db)
	
	apiRoutes := chi.NewRouter()
	SetupRoutes(apiRoutes, db, s.config)
	r.Mount("/api/v1", apiRoutes)
	log.Printf("Server running at port localhost%s", s.addr)
	return http.ListenAndServe(s.addr, r)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("DB: successfuly connected")
}
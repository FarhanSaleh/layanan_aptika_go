package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db, err := config.NewDB()
	if err != nil {
		return err
	}
	initStorage(db)

	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	apiRoutes := chi.NewRouter()
	SetupRoutes(apiRoutes, db)
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
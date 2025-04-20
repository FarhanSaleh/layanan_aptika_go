package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/farhansaleh/layanan_aptika_be/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
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

	validator := validator.New()
	db, err := config.NewDB()
	if err != nil {
		return err
	}
	initStorage(db)

	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	apiRoute := chi.NewRouter()
	
	usersRepository := users.NewRepository()
	usersServices := users.NewService(db, usersRepository, validator)
	usersHandler := users.NewHandler(usersServices)

	users.SetupRoutes(usersHandler, apiRoute)
	r.Mount("/api/v1", apiRoute)

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
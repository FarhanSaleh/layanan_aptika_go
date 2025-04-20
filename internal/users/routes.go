package users

import "github.com/go-chi/chi/v5"

func SetupRoutes(h Handler, r chi.Router) {
	r.Post("/users", h.Create)
	r.Put("/users/{id}", h.Update)
	r.Delete("/users/{id}", h.Delete)
	r.Get("/users", h.FindAll)
	r.Get("/users/{id}", h.FindById)
}
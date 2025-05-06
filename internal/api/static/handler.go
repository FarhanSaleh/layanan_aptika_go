package static

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	Image(w http.ResponseWriter, r *http.Request)
	Document(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct{}

func NewHandler() Handler {
	return &HandlerImpl{}
}

func (h *HandlerImpl) Image(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "filename")
	
	filePath := filepath.Join("uploads", "img", path)
	http.ServeFile(w, r, filePath)
}

func (h *HandlerImpl) Document(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "filename")
	filePath := filepath.Join("uploads", "docs", path)
	
	http.ServeFile(w, r, filePath)
}
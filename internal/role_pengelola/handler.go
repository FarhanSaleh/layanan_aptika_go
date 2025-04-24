package rolepengelola

import (
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
	"github.com/go-chi/chi/v5"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return &HandlerImpl{
		Service: service,
	}
}

func (h *HandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var request domain.RolePengelolaMutationRequest
	helper.ParseBody(r, &request)

	result, err := h.Service.Create(r.Context(), request)
	if err != nil {
		log.Println("Error Service: ", err)

		if validationErr, ok := helper.IsValidationError(err); ok {
			helper.WriteResponseBody(w, http.StatusBadRequest, domain.ErrorValidationResponse{
				Message: validationErr.Error(),
				Errors: validationErr.Errors,
			})
			return
		}

		helper.WriteResponseBody(w, http.StatusBadRequest, domain.DefaultResponse{
			Message: err.Error(),
		})
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: "Success Insert Data",
		Data: result,
	})
}

func (h *HandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var request domain.RolePengelolaMutationRequest
	helper.ParseBody(r, &request)
	
	result, err := h.Service.Update(r.Context(), request, id)
	if err != nil {
		log.Println("ERROR SERVICE: ", err)
		if validationErr, ok := helper.IsValidationError(err); ok {
			helper.WriteResponseBody(w, http.StatusBadRequest, domain.ErrorValidationResponse{
				Message: validationErr.Error(),
				Errors: validationErr.Errors,
			})
			return
		}

		helper.WriteResponseBody(w, http.StatusBadRequest, domain.DefaultResponse{
			Message: err.Error(),
		})
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: "Success Update",
		Data: result,
	})
}

func (h *HandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	err := h.Service.Delete(r.Context(), id)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteResponseBody(w, http.StatusBadRequest, domain.DefaultResponse{
			Message: err.Error(),
		})
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: "Success Delete",
	})
}

func (h *HandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.FindAll(r.Context())
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteResponseBody(w, http.StatusBadRequest, domain.DefaultResponse{
			Message: err.Error(),
		})
		return
	}
	if result == nil {
		helper.WriteResponseBody(w, http.StatusNotFound, domain.DefaultResponse{
			Message: "Instansi Empty",
		})
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: "Success",
		Data: result,
	})
}
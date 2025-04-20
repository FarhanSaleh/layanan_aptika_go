package users

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
	FindById(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct{
	Service Service
}

func NewHandler(service Service) Handler {
	return &HandlerImpl{
		Service: service,
	}
}

func (h *HandlerImpl) Create(w http.ResponseWriter, r *http.Request){
	request := domain.UserMutationRequest{}
	helper.ParseBody(r, &request)

	result, err := h.Service.Create(r.Context(), request)
	if err != nil {
		log.Println("Error Service: ", err)
		if validationErr, ok := helper.IsValidationError(err); ok {
			response := domain.ErrorValidationResponse{
				Message: validationErr.Error(),
				Errors: validationErr.Errors,
			}
			helper.WriteResponseBody(w, http.StatusBadRequest, response)
			return
		}

		response := domain.DefaultResponse{
			Message: err.Error(),
		}
		helper.WriteResponseBody(w, http.StatusBadRequest, response)
		return
	}

	response := domain.DefaultResponse{
		Message: "success insert",
		Data: result,
	}

	helper.WriteResponseBody(w, http.StatusOK, response)
}

func (h *HandlerImpl) Update(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	request := domain.UserMutationRequest{}
	helper.ParseBody(r, &request)

	result, err := h.Service.Update(r.Context(), request, id)
	if err != nil {
		log.Println("Error Service: ", err)
		if validationErr, ok := helper.IsValidationError(err); ok {
			response := domain.ErrorValidationResponse{
				Message: validationErr.Error(),
				Errors: validationErr.Errors,
			}
			helper.WriteResponseBody(w, http.StatusBadRequest, response)
			return
		}

		response := domain.DefaultResponse{
			Message: err.Error(),
		}
		helper.WriteResponseBody(w, http.StatusBadRequest, response)
		return
	}

	response := domain.DefaultResponse{
		Message: "success update",
		Data: result,
	}

	helper.WriteResponseBody(w, http.StatusOK, response)
}

func (h *HandlerImpl) Delete(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	
	err := h.Service.Delete(r.Context(), id)
	if err != nil {
		log.Println("Error Service:", err)
		response := domain.DefaultResponse{
			Message: err.Error(),
		}
		helper.WriteResponseBody(w, http.StatusBadRequest, response)
		return
	}

	response := domain.DefaultResponse{
		Message: "Success Delete",
	}

	helper.WriteResponseBody(w, http.StatusOK, response)
}

func (h *HandlerImpl) FindAll(w http.ResponseWriter, r *http.Request){
	result, err := h.Service.FindAll(r.Context())
	if err != nil {
		log.Println("Error Service:", err)
		response := domain.DefaultResponse{
			Message: err.Error(),
		}
		helper.WriteResponseBody(w, http.StatusBadRequest, response)
		return
	}
	if result == nil {
		response := domain.DefaultResponse{
			Message: "Users Empty",
		}
		helper.WriteResponseBody(w, http.StatusNotFound, response)
		return
	}

	response := domain.DefaultResponse{
		Message: "Success",
		Data: result,
	}
	helper.WriteResponseBody(w, http.StatusOK, response)
}

func (h *HandlerImpl) FindById(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	result, err := h.Service.FindById(r.Context(), id)
	if err != nil {
		log.Println("Error Service:", err)
		response := domain.DefaultResponse{
			Message: err.Error(),
		}
		helper.WriteResponseBody(w, http.StatusBadRequest, response)
		return
	}

	response := domain.DefaultResponse{
		Message: "Success",
		Data: result,
	}
	helper.WriteResponseBody(w, http.StatusOK, response)
}
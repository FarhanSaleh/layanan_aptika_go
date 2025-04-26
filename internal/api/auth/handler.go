package auth

import (
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
)

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
	PengelolaLogin(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return &HandlerImpl{
		Service: service,
	}
} 

func (h *HandlerImpl) Login(w http.ResponseWriter, r *http.Request){
	request := domain.LoginRequest{}
	helper.ParseBody(r, &request)

	result, err := h.Service.Login(r.Context(), request)
	if err != nil {
		log.Println("Error Service: ", err)

		validationErr, ok := helper.IsValidationError(err)
		if ok {
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

	response := domain.DefaultResponse{
		Message: "Success Login",
		Data: result,
	}
	helper.WriteResponseBody(w, http.StatusOK, response)
}

func (h *HandlerImpl) PengelolaLogin(w http.ResponseWriter, r *http.Request){
	request := domain.LoginRequest{}
	helper.ParseBody(r, &request)

	result, err := h.Service.PengelolaLogin(r.Context(), request)
	if err != nil {
		log.Println("Error Service: ", err)

		validationErr, ok := helper.IsValidationError(err)
		if ok {
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

	response := domain.DefaultResponse{
		Message: "Success Login",
		Data: result,
	}
	helper.WriteResponseBody(w, http.StatusOK, response)
}

func (h *HandlerImpl) Logout(w http.ResponseWriter, r *http.Request){
	request := map[string]any{}
	helper.ParseBody(r, &request)

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: "Endpoint belum digunakan",
		Data: request,
	})
}
func (h *HandlerImpl) ChangePassword(w http.ResponseWriter, r *http.Request){
	
}
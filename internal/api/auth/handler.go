package auth

import (
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/constants"
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
		log.Println("ERROR SERVICE: ", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessLogin,
		Data: result,
	})
}

func (h *HandlerImpl) PengelolaLogin(w http.ResponseWriter, r *http.Request){
	request := domain.LoginRequest{}
	helper.ParseBody(r, &request)

	result, err := h.Service.PengelolaLogin(r.Context(), request)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessLogin,
		Data: result,
	})
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
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
	PengelolaChangePassword(w http.ResponseWriter, r *http.Request)
	UserChangePassword(w http.ResponseWriter, r *http.Request)
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
	err := h.Service.Logout(r.Context())
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessLogout,
	})
}

func (h *HandlerImpl) PengelolaChangePassword(w http.ResponseWriter, r *http.Request){
	request := domain.ChangePasswordRequest{}
	helper.ParseBody(r, &request)

	err := h.Service.PengelolaChangePassword(r.Context(), request)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessUpdate,
	})
}

func (h *HandlerImpl) UserChangePassword(w http.ResponseWriter, r *http.Request){
	request := domain.ChangePasswordRequest{}
	helper.ParseBody(r, &request)

	err := h.Service.UserChangePassword(r.Context(), request)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessUpdate,
	})
}
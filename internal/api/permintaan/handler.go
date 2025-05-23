package permintaan

import (
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/constants"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
)

type Handler interface {
	CountAll(w http.ResponseWriter, r *http.Request)
	CountGangguanJIP(w http.ResponseWriter, r *http.Request)
	CountPembuatanEmail(w http.ResponseWriter, r *http.Request)
	CountPembuatanSubdomain(w http.ResponseWriter, r *http.Request)
	CountPembangunanAplikasi(w http.ResponseWriter, r *http.Request)
	CountPusatDataDaerah(w http.ResponseWriter, r *http.Request)
	CountPerubahanIPServer(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return &HandlerImpl{Service: service}
}

func (h *HandlerImpl) CountAll(w http.ResponseWriter, r *http.Request) {
	groupBy := r.URL.Query().Get("group_by")
	year := r.URL.Query().Get("year")

	var result any
	var err error

	if(groupBy == "bulan") {
		result, err = h.Service.CountAllPerMonth(r.Context(), year)
		if err != nil {
			log.Println("ERROR SERVICE:", err)
			helper.WriteErrorResponse(w, err)
			return
		}
	}else{
		result, err = h.Service.CountAll(r.Context())
		if err != nil {
			log.Println("ERROR SERVICE:", err)
			helper.WriteErrorResponse(w, err)
			return
		}
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessGetData,
		Data:    result,
	})
}

func (h *HandlerImpl) CountGangguanJIP(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.CountGangguanJIP(r.Context())
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}
	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessGetData,
		Data:    result,
	})
}

func (h *HandlerImpl) CountPembuatanEmail(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.CountPembuatanEmail(r.Context())
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}
	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessGetData,
		Data:    result,
	})
}

func (h *HandlerImpl) CountPembuatanSubdomain(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.CountPembuatanSubdomain(r.Context())
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}
	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessGetData,
		Data:    result,
	})
}

func (h *HandlerImpl) CountPembangunanAplikasi(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.CountPembangunanAplikasi(r.Context())
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}
	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessGetData,
		Data:    result,
	})
}

func (h *HandlerImpl) CountPusatDataDaerah(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.CountPusatDataDaerah(r.Context())
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}
	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessGetData,
		Data:    result,
	})
}

func (h *HandlerImpl) CountPerubahanIPServer(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.CountPerubahanIPServer(r.Context())
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}
	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessGetData,
		Data:    result,
	})
}

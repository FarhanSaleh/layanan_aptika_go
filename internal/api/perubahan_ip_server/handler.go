package perubahanipserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/constants"
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
	FindByUser(w http.ResponseWriter, r *http.Request)
	UpdateStatus(w http.ResponseWriter, r *http.Request)
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
	r.Body = http.MaxBytesReader(w, r.Body, constants.MaxBytesReader)
	err := r.ParseMultipartForm(constants.MaxUploadSize)
	if err != nil {
		log.Println("ERROR PARSING MULTIPARTFORM:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	err = helper.CheckFormFile(r, "surat_permohonan")
	if err != nil {
		log.Println("ERROR CHECK FORM FILE:", err)
		helper.WriteErrorResponse(w, fmt.Errorf("surat permohonan (PDF) is required"))
		return
	}
	
	suratFileName, err := helper.HandleUploadPdf(w, r, "surat_permohonan")
	if err != nil {
		log.Println("ERROR UPLOAD SURAT:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	request := domain.PerubahanIPServerMutationRequest{
		NamaLengkap:       r.FormValue("nama_lengkap"),
		Jabatan:           r.FormValue("jabatan"),
		NomorHP:           r.FormValue("nomor_hp"),
		NamaSubdomain:     r.FormValue("nama_subdomain"),
		IPLama:            r.FormValue("ip_lama"),
		IPBaru: 		   r.FormValue("ip_baru"),
		SuratPermohonan:   suratFileName,
		InstansiId:        r.FormValue("instansi_id"),
	}

	response, err := h.Service.Create(r.Context(), request)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessInsert,
		Data: response,
	})
}

func (h *HandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	r.Body = http.MaxBytesReader(w, r.Body, constants.MaxBytesReader)
	err := r.ParseMultipartForm(constants.MaxUploadSize)
	if err != nil {
		log.Println("ERROR PARSING MULTIPARTFORM:", err)
		helper.WriteErrorResponse(w, err)
		return
	}
	
	suratFileName, err := helper.HandleUploadPdf(w, r, "surat_permohonan")
	if err != nil {
		log.Println("ERROR UPLOAD SURAT:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	request := domain.PerubahanIPServerMutationRequest{
		NamaLengkap:       r.FormValue("nama_lengkap"),
		Jabatan:           r.FormValue("jabatan"),
		NomorHP:           r.FormValue("nomor_hp"),
		NamaSubdomain:     r.FormValue("nama_subdomain"),
		IPLama:            r.FormValue("ip_lama"),
		IPBaru: 		   r.FormValue("ip_baru"),
		SuratPermohonan:   suratFileName,
		InstansiId:        r.FormValue("instansi_id"),
	}

	response, err := h.Service.Update(r.Context(), request, id)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessInsert,
		Data: response,
	})
}

func (h *HandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.Service.Delete(r.Context(), id)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessDelete,
	})
}

func (h *HandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.FindAll(r.Context())
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

func (h *HandlerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := h.Service.FindById(r.Context(), id)
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

func (h *HandlerImpl) FindByUser(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.FindAllByUser(r.Context())
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

func (h *HandlerImpl) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var request domain.UpdateStatusLayananRequest
	helper.ParseBody(r, &request)

	err := h.Service.UpdateStatus(r.Context(), request, id)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		helper.WriteErrorResponse(w, err)
		return
	}

	helper.WriteResponseBody(w, http.StatusOK, domain.DefaultResponse{
		Message: constants.SuccessUpdate,
	})
}

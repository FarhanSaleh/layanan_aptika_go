package domain

import (
	"database/sql"
	"time"
)

type PerubahanIPServer struct {
	Id                string
	NamaLengkap       string
	Jabatan           string
	NomorHP           string
	NamaSubdomain     string
	IPLama  	 	  string
	IPBaru            string
	SuratPermohonan   string
	Status            string
	IsDeleted 		  string
	CreatedAt         time.Time
	UpdatedAt         sql.NullTime
	InstansiId        string
	NamaInstansi 	  string
	UserId            string
	NamaUser          string
}

type PerubahanIPServerResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	NamaSubdomain     string `json:"nama_subdomain"`
	IPLama  	 	  string `json:"ip_lama"`
	IPBaru            string `json:"ip_baru"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
}

type PerubahanIPServerDetailResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	NamaSubdomain     string `json:"nama_subdomain"`
	IPLama  	 	  string `json:"ip_lama"`
	IPBaru            string `json:"ip_baru"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt		  string `json:"updated_at"`
}

type PerubahanIPServerMutationResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	NamaSubdomain     string `json:"nama_subdomain"`
	IPLama  	 	  string `json:"ip_lama"`
	IPBaru            string `json:"ip_baru"`
	SuratPermohonan   string `json:"surat_permohonan"`
	InstansiId        string `json:"instansi_id"`
}

type PerubahanIPServerMutationRequest struct {
	NamaLengkap       string `validate:"required,ascii,max=255,min=3"`
	Jabatan           string `validate:"required,ascii,max=255,min=3"`
	NomorHP           string `validate:"required,numeric,max=15,min=3"`
	NamaSubdomain     string `validate:"required,ascii,max=255,min=3"`
	IPLama 			  string `validate:"required,ip"`
	IPBaru 			  string `validate:"required,ip"`
	SuratPermohonan   string 
	InstansiId        string `validate:"required,uuid"`
}
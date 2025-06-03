package domain

import (
	"database/sql"
	"time"
)

type PembuatanSubdomain struct {
	Id                string
	NamaLengkap       string
	Jabatan        	  string
	NomorHP           string
	NamaSubdomain     string
	IPPublik     	  string
	Deskripsi	      string
	SuratPermohonan   string
	Status            string
	IsDeleted 		  string
	CreatedAt         time.Time
	UpdatedAt         sql.NullTime
	InstansiId        string
	NamaInstansi 	  string
	UserId            string
	NamaUser          string
	NotificationToken sql.NullString
}

type PembuatanSubdomainResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan        	  string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	NamaSubdomain     string `json:"nama_subdomain"`
	IPPublik     	  string `json:"ip_publik"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
}

type PembuatanSubdomainDetailResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan        	  string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	NamaSubdomain     string `json:"nama_subdomain"`
	IPPublik     	  string `json:"ip_publik"`
	Deskripsi     	  string `json:"deskripsi"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type PembuatanSubdomainMutationResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan        	  string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	NamaSubdomain     string `json:"nama_subdomain"`
	IPPublik     	  string `json:"ip_publik"`
	Deskripsi     	  string `json:"deskripsi"`
	SuratPermohonan   string `json:"surat_permohonan"`
	InstansiId        string `json:"instansi_id"`
}

type PembuatanSubdomainMutationRequest struct {
	NamaLengkap       string `validate:"required,ascii,max=255,min=3"`
	Jabatan           string `validate:"required,ascii,max=255,min=3"`
	NomorHP           string `validate:"required,numeric,max=15,min=3"`
	NamaSubdomain     string `validate:"required,ascii"`
	IPPublik          string `validate:"required,ip"`
	Deskripsi         string `validate:"required,ascii"`
	SuratPermohonan   string 
	InstansiId        string `validate:"required,uuid"`
}
package domain

import (
	"database/sql"
	"time"
)

type PembangunanAplikasi struct {
	Id                string
	NamaPimpinan      string
	NomorHP           string
	EmailDinas        string
	RiwayatPimpinan   string
	JenisAplikasi     string
	TujuanAplikasi    string
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

type PembangunanAplikasiResponse struct {
	Id                string `json:"id"`
	NamaPimpinan      string `json:"nama_pimpinan"`
	NomorHP           string `json:"nomor_hp"`
	EmailDinas	      string `json:"email_dinas"`
	JenisAplikasi     string `json:"jenis_aplikasi"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
}

type PembangunanAplikasiDetailResponse struct {
	Id                string `json:"id"`
	NamaPimpinan      string `json:"nama_pimpinan"`
	NomorHP           string `json:"nomor_hp"`
	EmailDinas	      string `json:"email_dinas"`
	RiwayatPimpinan   string `json:"riwayat_pimpinan"`
	JenisAplikasi     string `json:"jenis_aplikasi"`
	TujuanAplikasi    string `json:"tujuan_aplikasi"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type PembangunanAplikasiMutationResponse struct {
	Id                string `json:"id"`
	NamaPimpinan      string `json:"nama_pimpinan"`
	NomorHP           string `json:"nomor_hp"`
	EmailDinas	      string `json:"email_dinas"`
	RiwayatPimpinan   string `json:"riwayat_pimpinan"`
	JenisAplikasi     string `json:"jenis_aplikasi"`
	TujuanAplikasi    string `json:"tujuan_aplikasi"`
	SuratPermohonan   string `json:"surat_permohonan"`
	InstansiId        string `json:"instansi_id"`
}

type PembangunanAplikasiMutationRequest struct {
	NamaPimpinan      string `validate:"required,ascii,max=255,min=3"`
	NomorHP           string `validate:"required,numeric,max=15,min=3"`
	EmailDinas        string `validate:"required,email,max=255,min=3"`
	RiwayatPimpinan   string `validate:"required,ascii"`
	JenisAplikasi     string `validate:"required,ascii"`
	TujuanAplikasi    string `validate:"required,ascii"`
	SuratPermohonan   string 
	InstansiId        string `validate:"required,uuid"`
}
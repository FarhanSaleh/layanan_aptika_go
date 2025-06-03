package domain

import (
	"database/sql"
	"time"
)

type PusatDataDaerah struct {
	Id                string
	NamaLengkap       string
	Jabatan           string
	NomorHP           string
	JenisLayanan      string
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

type PusatDataDaerahResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	JenisLayanan      string `json:"jenis_layanan"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
}

type PusatDataDaerahDetailResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	JenisLayanan      string `json:"jenis_layanan"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt		  string `json:"updated_at"`
}

type PusatDataDaerahMutationResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	JenisLayanan      string `json:"jenis_layanan"`
	SuratPermohonan   string `json:"surat_permohonan"`
	InstansiId        string `json:"instansi_id"`
}

type PusatDataDaerahMutationRequest struct {
	NamaLengkap       string `validate:"required,ascii,max=255,min=3"`
	Jabatan           string `validate:"required,ascii,max=255,min=3"`
	NomorHP           string `validate:"required,numeric,max=15,min=3"`
	JenisLayanan      string `validate:"required,ascii"`
	SuratPermohonan   string 
	InstansiId        string `validate:"required,uuid"`
}
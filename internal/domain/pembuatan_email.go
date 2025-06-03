package domain

import (
	"database/sql"
	"time"
)

type PembuatanEmail struct {
	Id                string
	NamaLengkap       string
	NIP        	  	  string
	Jabatan        	  string
	NomorHP           string
	BerkasSK          string
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

type PembuatanEmailResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	NIP        	  	  string `json:"nip"`
	Jabatan        	  string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	BerkasSK	      string `json:"berkas_sk"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
}

type PembuatanEmailDetailResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	NIP        	  	  string `json:"nip"`
	Jabatan        	  string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	BerkasSK	      string `json:"berkas_sk"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type PembuatanEmailMutationResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	NIP        	  	  string `json:"nip"`
	Jabatan        	  string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	BerkasSK	      string `json:"berkas_sk"`
	SuratPermohonan   string `json:"surat_permohonan"`
	InstansiId        string `json:"instansi_id"`
}

type PembuatanEmailMutationRequest struct {
	NamaLengkap       string `validate:"required,ascii,max=255,min=3"`
	NIP               string `validate:"required,numeric,max=18,min=18"`
	Jabatan           string `validate:"required,ascii,max=255,min=3"`
	NomorHP           string `validate:"required,numeric,max=15,min=3"`
	BerkasSK          string 
	SuratPermohonan   string 
	InstansiId        string `validate:"required,uuid"`
}
package domain

import (
	"database/sql"
	"time"
)

type GangguanJIP struct {
	Id                string
	NamaLengkap       string
	Jabatan           string
	NomorHP           string
	LokasiGangguan    string
	DeskripsiGangguan string
	Foto              string
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

type GangguanJIPResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	LokasiGangguan    string `json:"lokasi_gangguan"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
}

type GangguanJIPDetailResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	LokasiGangguan    string `json:"lokasi_gangguan"`
	DeskripsiGangguan string `json:"deskripsi_gangguan"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Foto              string `json:"foto"`
	Status            string `json:"status"`
	InstansiId        string `json:"instansi_id"`
	NamaInstansi 	  string `json:"nama_instansi"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type GangguanJIPMutationResponse struct {
	Id                string `json:"id"`
	NamaLengkap       string `json:"nama_lengkap"`
	Jabatan           string `json:"jabatan"`
	NomorHP           string `json:"nomor_hp"`
	LokasiGangguan    string `json:"lokasi_gangguan"`
	DeskripsiGangguan string `json:"deskripsi_gangguan"`
	SuratPermohonan   string `json:"surat_permohonan"`
	Foto              string `json:"foto"`
	InstansiId        string `json:"instansi_id"`
}

type GangguanJIPMutationRequest struct {
	NamaLengkap       string `validate:"required,ascii,max=255,min=3"`
	Jabatan           string `validate:"required,ascii,max=255,min=3"`
	NomorHP           string `validate:"required,numeric,max=255,min=3"`
	LokasiGangguan    string `validate:"required,ascii"`
	DeskripsiGangguan string `validate:"required,ascii"`
	SuratPermohonan   string 
	Foto              string 
	InstansiId        string `validate:"required,uuid"`
}
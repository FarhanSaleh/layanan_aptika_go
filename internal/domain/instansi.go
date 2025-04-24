package domain

import (
	"database/sql"
	"time"
)

type Instansi struct {
	Id         string
	Nama       string
	Alamat     string
	Keterangan sql.NullString
	IsDeleted  bool
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type InstansiResponse struct {
	Id         string `json:"id"`
	Nama       string `json:"nama"`
	Alamat     string `json:"alamat"`
	Keterangan string `json:"keterangan"`
}

type InstansiMutationRequest struct {
	Nama       string `json:"nama" validate:"required"`
	Alamat     string `json:"alamat" validate:"required"`
	Keterangan string `json:"keterangan" validate:"alphanum"`
}
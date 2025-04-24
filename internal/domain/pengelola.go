package domain

import (
	"database/sql"
	"time"
)

type Pengelola struct {
	Id           string
	Nama         string
	Email        string
	Password     string
	RefreshToken sql.NullString
	IsDeleted    bool
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
	RoleId		 string
	NamaRole 	 string
}

type PengelolaMutateResponse struct {
	Id    	string `json:"id"`
	Nama  	string `json:"nama"`
	Email	string `json:"email"`
	RoleId 	string `json:"role_id"`
}

type PengelolaResponse struct {
	Id    	 string `json:"id"`
	Nama  	 string `json:"nama"`
	Email	 string `json:"email"`
	NamaRole string `json:"nama_role"`
}

type PengelolaDetailResponse struct {
	Id        string `json:"id"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	NamaRole  string `json:"nama_role"`
	CreatedAt string `json:"created_at"`
}

type PengelolaMutationRequest struct {
	Nama   string `json:"nama" validate:"required,alpha,max=255,min=3"`
	Email  string `json:"email" validate:"required,email"`
	RoleId string `json:"role_id" validate:"required,uuid"`
}
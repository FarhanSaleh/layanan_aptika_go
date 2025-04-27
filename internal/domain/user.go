package domain

import (
	"database/sql"
	"time"
)

type User struct {
	Id           string
	Nama         string
	Email        string
	Password     string
	RefreshToken sql.NullString
	IsDeleted    bool
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
}

type UserResponse struct {
	Id        string `json:"id"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
}

type UserDetailResponse struct {
	Id        string `json:"id"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type UserMutationRequest struct {
	Nama      string `json:"nama" validate:"required,ascii,max=255,min=3"`
	Email     string `json:"email" validate:"required,email"`
}
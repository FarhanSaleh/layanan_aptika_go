package domain

import (
	"database/sql"
	"time"
)

type RolePengelola struct {
	Id        string
	Nama      string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type RolePengelolaResponse struct {
	Id 		string `json:"id"`
	Nama 	string `json:"nama"`
}

type RolePengelolaMutationRequest struct {
	Nama string `json:"nama" validate:"required,alphanum"`
}
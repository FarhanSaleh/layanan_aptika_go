package domain

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UID   string  `json:"uid"`
	Email string  `json:"email"`
	Nama  string  `json:"nama"`
	RoleId string `json:"role_id,omitempty"`
	RoleName string `json:"nama_role,omitempty"`
	jwt.RegisteredClaims
}
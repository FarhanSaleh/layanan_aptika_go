package domain

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Email string `json:"email"`
	Nama  string `json:"nama"`
	jwt.RegisteredClaims
}
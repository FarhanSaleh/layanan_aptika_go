package middlewares

import (
	"log"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.InitEnvs()

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			helper.WriteResponseBody(w, http.StatusUnauthorized, domain.DefaultResponse{
				Message: "Token tidak ditemukan",
			})
            return
        }
		
		tokenStr := authHeader[7:]
		
		claims := &domain.JWTClaims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return []byte(conf.JWTSecret), nil
        })
		
		log.Println(token.Valid)

        if err != nil || !token.Valid {
            helper.WriteResponseBody(w, http.StatusUnauthorized, domain.DefaultResponse{
				Message: "Token tidak valid",
			})
            return
        }

		next.ServeHTTP(w, r)
	})
}
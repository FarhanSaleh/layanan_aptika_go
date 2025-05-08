package middlewares

import (
	"context"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/config"
	contextkey "github.com/farhansaleh/layanan_aptika_be/internal/context_key"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/farhansaleh/layanan_aptika_be/pkg/helper"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuthMiddleware(next http.Handler) http.Handler {
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
		
        if err != nil || !token.Valid {
            helper.WriteResponseBody(w, http.StatusUnauthorized, domain.DefaultResponse{
				Message: "Token tidak valid",
			})
            return
        }

		tokenClaims := token.Claims.(*domain.JWTClaims)
		ctx := context.WithValue(r.Context(), contextkey.UserKey, tokenClaims)
		ctx = context.WithValue(ctx, contextkey.TypeAccountKey, "user")
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func PengelolaAuthMiddleware(next http.Handler) http.Handler {
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
			return []byte(conf.JWTPengelolaSecret), nil
        })
		
        if err != nil || !token.Valid {
            helper.WriteResponseBody(w, http.StatusUnauthorized, domain.DefaultResponse{
				Message: "Token tidak valid",
			})
            return
        }

		tokenClaims := token.Claims.(*domain.JWTClaims)
		ctx := context.WithValue(r.Context(), contextkey.PengelolaKey, tokenClaims.Email)
		ctx = context.WithValue(ctx, contextkey.TypeAccountKey, "pengelola")
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
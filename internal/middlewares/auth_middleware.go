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
				Message: "Unauthorized",
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
				Message: "Unauthorized",
			})
            return
        }

		tokenClaims := token.Claims.(*domain.JWTClaims)
		ctx := context.WithValue(r.Context(), contextkey.UserKey, tokenClaims)
		ctx = context.WithValue(ctx, contextkey.TypeAccountKey, "user")
		ctx = context.WithValue(ctx, contextkey.RoleKey, tokenClaims.RoleId)
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
				Message: "Unauthorized",
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
				Message: "Unauthorized",
			})
            return
        }

		tokenClaims := token.Claims.(*domain.JWTClaims)
		ctx := context.WithValue(r.Context(), contextkey.PengelolaKey, tokenClaims.Email)
		ctx = context.WithValue(ctx, contextkey.TypeAccountKey, "pengelola")
		ctx = context.WithValue(ctx, contextkey.RoleKey, tokenClaims.RoleId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roleId := r.Context().Value(contextkey.RoleKey)

			if roleId == nil{
				helper.WriteResponseBody(w, http.StatusUnauthorized, domain.DefaultResponse{
					Message: "Unauthorized",
				})
				return
			}

			for _, role := range allowedRoles {
				if roleId == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			helper.WriteResponseBody(w, http.StatusUnauthorized, domain.DefaultResponse{
				Message: "Unauthorized",
			})
		})
	}
}
package helper

import (
	"log"
	"time"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user domain.User) (tokenString string, err error) {
	conf := config.InitEnvs()

	expTime := time.Now().Add(time.Hour)
	claims := domain.JWTClaims{
		Email: user.Email,
		Nama:  user.Nama,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(conf.JWTSecret))
	
	if err != nil {
		log.Println("Error generate token: ", err)
		return 
	}
	return
}
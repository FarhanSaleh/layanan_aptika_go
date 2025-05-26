package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseDSN 			string
	Port        			string
	JWTPublicKey 			string
	JWTPrivateKey			string
	JWTSecret 				string
	JWTPengelolaSecret 		string
	AllowedOrigins 			[]string
	StaticDocsOriginUser 	string
	StaticImgOriginUser 	string
	StaticDocsOriginPengelola string
	StaticImgOriginPengelola string
}

func InitEnvs() Config {
	return Config{
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
		Port: os.Getenv("PORT"),
		JWTPublicKey: os.Getenv("JWT_PUBLIC_KEY"),
		JWTPrivateKey: os.Getenv("JWT_PRIVATE_KEY"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTPengelolaSecret: os.Getenv("JWT_PENGELOLA_SECRET"),
		AllowedOrigins: []string{os.Getenv("DEV_ORIGIN"), os.Getenv("PROD_ORIGIN")},
		StaticDocsOriginUser: fmt.Sprintf("%s%s/%s", os.Getenv("HOST_ORIGIN"), os.Getenv("PORT"), os.Getenv("STATIC_DOCS_ORIGIN_USER")),
		StaticImgOriginUser: fmt.Sprintf("%s%s/%s", os.Getenv("HOST_ORIGIN"), os.Getenv("PORT"), os.Getenv("STATIC_IMG_ORIGIN_USER")),
		StaticDocsOriginPengelola: fmt.Sprintf("%s%s/%s", os.Getenv("HOST_ORIGIN"), os.Getenv("PORT"), os.Getenv("STATIC_DOCS_ORIGIN_PENGELOLA")),
		StaticImgOriginPengelola: fmt.Sprintf("%s%s/%s", os.Getenv("HOST_ORIGIN"), os.Getenv("PORT"), os.Getenv("STATIC_IMG_ORIGIN_PENGELOLA")),
	}
}

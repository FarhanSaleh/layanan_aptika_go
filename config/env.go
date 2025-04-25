package config

import (
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
	}
}

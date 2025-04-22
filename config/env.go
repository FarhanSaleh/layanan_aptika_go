package config

import (
	"os"
)

type Config struct {
	DatabaseDSN 	string
	Port        	string
	JWTPublicKey 	string
	JWTPrivateKey	string
	JWTSecret 		string
}

func InitEnvs() Config {
	return Config{
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
		Port: os.Getenv("PORT"),
		JWTPublicKey: os.Getenv("JWT_PUBLIC_KEY"),
		JWTPrivateKey: os.Getenv("JWT_PRIVATE_KEY"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

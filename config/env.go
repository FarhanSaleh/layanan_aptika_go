package config

import (
	"os"
)

type Config struct {
	DatabaseDSN string
	Port        string
}

func InitEnvs() Config {
	return Config{
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
		Port: os.Getenv("PORT"),
	}
}

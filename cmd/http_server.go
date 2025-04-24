package main

import (
	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/farhansaleh/layanan_aptika_be/internal/api"
	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	config := config.InitEnvs()

	server := api.NewAPIServer(config.Port, &config)
	
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
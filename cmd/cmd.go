package cmd

import (
	"fmt"
	"log"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "layanan-aptika-be",
	Short: "API layanan APTIKA DISKOMINFOSANTIK Provinsi Sulawesi Tengah",
	Long:  `API layanan APTIKA DISKOMINFOSANTIK Provinsi Sulawesi Tengah`,
}

var testEnvCmd = &cobra.Command{
	Use:   "test-env",
	Short: "Run test environment",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.InitEnvs()
		
		fmt.Println(cfg.StaticDocsOriginUser)
		fmt.Println(cfg.StaticImgOriginUser)
		fmt.Println(cfg.StaticDocsOriginPengelola)
		fmt.Println(cfg.StaticImgOriginPengelola)
	},
}

func Execute() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(serveHttpCmd, migrateCreateCmd, migrateDownCmd, migrateUpCmd, createSeederCmd, runSeederCmd, runAllSeederCmd, testEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("error executing root command", err)
	}
}
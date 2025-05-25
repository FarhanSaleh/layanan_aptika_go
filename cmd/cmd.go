package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "layanan-aptika-be",
	Short: "API layanan APTIKA DISKOMINFOSANTIK Provinsi Sulawesi Tengah",
	Long:  `API layanan APTIKA DISKOMINFOSANTIK Provinsi Sulawesi Tengah`,
}

func Execute() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(serveHttpCmd, migrateCreateCmd, migrateDownCmd, migrateUpCmd, createSeederCmd, runSeederCmd, runAllSeederCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("error executing root command", err)
	}
}
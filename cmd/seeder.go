package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/spf13/cobra"
)

var createSeederCmd = &cobra.Command{
	Use:   "seeder:create [name]",
	Short: "Run seeder",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirName := "database/seeders"
		name := strings.ToLower(args[0])
		filename := fmt.Sprintf("%s/%s.sql", dirName, name)
		content := "-- +seeder"

		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating seeders folder:", err)
			return
		}

		err = os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error creating seeder file:", err)
			return
		}

		fmt.Println("Seeder created:", filename)
	},
}

var runAllSeederCmd = &cobra.Command{
	Use:   "seeder:all",
	Short: "Run all seeders",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := config.NewDB()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close()

		files, _ := filepath.Glob("database/seeders/*.sql")
		for _, file := range files {
			content, _ := os.ReadFile(file)
			if _, err := conn.ExecContext(context.Background(), string(content)); err != nil {
				fmt.Println("Failed to run seeder:", file, err)
				return
			}

			fmt.Println("Seeded:", file)
		}
	},
}

var runSeederCmd = &cobra.Command{
	Use:   "seeder:run [name]",
	Short: "Run seeder",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := config.NewDB()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close()

		file := fmt.Sprintf("database/seeders/%s.sql", strings.ToLower(args[0]))
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading seeder file:", err)
			return
		}
		if _, err := conn.ExecContext(context.Background(), string(content)); err != nil {
			fmt.Println("Failed to run seeder:", file, err)
			return
		}

		fmt.Println("Seeded:", file)
	},
}
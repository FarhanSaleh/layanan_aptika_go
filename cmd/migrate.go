package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/spf13/cobra"
)

var migrateCreateCmd = &cobra.Command{
	Use: "migrate:create [name]",
	Short: "Create migration file",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirName := "database/migrations"
		name := strings.ToLower(args[0])
		timestamp := time.Now().Format("20060102150405")
		filename := fmt.Sprintf("%s/%s_%s.sql", dirName, timestamp, name)
		content := "-- +migrate Up\n\n-- +migrate Down\n"

		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating migrations folder:", err)
			return
		}

		err = os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error creating migration file:", err)
			return
		}

		fmt.Println("Migration created:", filename)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "Revert all down migrations",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := config.NewDB()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close()

		files, _ := filepath.Glob("database/migrations/*.sql")
		for i := len(files) - 1; i >= 0; i-- {
			file := files[i]
			content, _ := os.ReadFile(file)
			parts := strings.Split(string(content), "-- +migrate Down")
			if len(parts) < 2 {
				continue
			}
			downSQL := parts[1]

			if _, err := conn.ExecContext(context.Background(), downSQL); err != nil {
				fmt.Println("Failed to rollback migration:", file, err)
				return
			}

			fmt.Println("Rolled back:", file)
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "Apply all up migrations",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := config.NewDB()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close()
		
		files, err := filepath.Glob("database/migrations/*.sql")
		if err != nil {
			fmt.Println("Error getting migration files:", err)
			return
		}

		for _, file := range files {
			content, _ := os.ReadFile(file)
			parts := strings.Split(string(content), "-- +migrate Down")
			upSQL := strings.Replace(parts[0], "-- +migrate Up", "", 1)

			if _, err := conn.ExecContext(context.Background(), upSQL); err != nil {
				fmt.Println("Failed to run migration:", file, err)
				return
			}

			fmt.Println("Migrated:", file)
		}
	},
}
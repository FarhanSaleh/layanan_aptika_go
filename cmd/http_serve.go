package cmd

import (
	"github.com/farhansaleh/layanan_aptika_be/config"
	"github.com/farhansaleh/layanan_aptika_be/internal/api"
	"github.com/spf13/cobra"
)

var serveHttpCmd = &cobra.Command{
	Use: "http",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.InitEnvs()

		server := api.NewAPIServer(config.Port, &config)
		
		err := server.Run()
		if err != nil {
			panic(err)
		}
	},
}
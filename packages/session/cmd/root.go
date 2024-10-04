package cmd

import (
	"catworks/luna/session/internal/config"
	"catworks/luna/session/internal/di"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "session",
	Short: "Session service for Luna app",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := cmd.Flag("config").Value.String()
		cfg := config.Require(configPath)

		container, _ := di.NewContainer(cfg)
		server := container.Server

		server.Register()

		err := server.Start()
		if err != nil {
			container.Logger.WithError(err).Fatal("Failed to start server")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "config file (default is config.yaml)")
}

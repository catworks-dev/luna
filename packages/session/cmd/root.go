package cmd

import (
	"catworks/luna/session/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "session",
	Short: "Session service for Luna app",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := cmd.Flag("config").Value.String()
		cfg := config.Require(configPath)

		container, _ := config.NewContainer(cfg)
		container.Logger.WithField("config", container.Config).Info("Config loaded")
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

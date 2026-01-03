package cmd

import (
	"github.com/spf13/cobra"
)

// caddyCmd represents the caddy command group
var caddyCmd = &cobra.Command{
	Use:   "caddy",
	Short: "Manage Caddy configuration",
	Long:  `Actions related to Caddyfile generation and management.`,
}

func init() {
	rootCmd.AddCommand(caddyCmd)
}
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "konf",
	Short: "A CLI to manage configuration files",
	Long:  `konf is a tool to generate and manage configuration files for various systems (Caddy, Nginx, etc).`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags can go here
}
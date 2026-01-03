package cmd

import (
	"github.com/phillip-england/konf/internal/caddy"
	"github.com/spf13/cobra"
)

var forceInit bool

// caddyInitCmd represents the init command
var caddyInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Caddyfile with defaults",
	RunE: func(cmd *cobra.Command, args []string) error {
		return caddy.GenerateDefault(forceInit)
	},
}

func init() {
	// Add 'init' to the 'caddy' command
	caddyCmd.AddCommand(caddyInitCmd)

	// Add a local flag for forcing overwrite
	caddyInitCmd.Flags().BoolVarP(&forceInit, "force", "f", false, "Overwrite existing Caddyfile")
}

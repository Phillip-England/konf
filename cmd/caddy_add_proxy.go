package cmd

import (
	"github.com/phillip-england/konf/internal/caddy"
	"github.com/spf13/cobra"
)

// caddyAddProxyCmd adds a reverse proxy site.
var caddyAddProxyCmd = &cobra.Command{
	Use:   "proxy <host> <upstream>",
	Short: "Add a reverse proxy site block",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return caddy.AppendReverseProxy(args[0], args[1])
	},
}

func init() {
	caddyAddCmd.AddCommand(caddyAddProxyCmd)
}

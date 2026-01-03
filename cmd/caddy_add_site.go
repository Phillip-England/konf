package cmd

import (
	"github.com/phillip-england/konf/internal/caddy"
	"github.com/spf13/cobra"
)

var siteRoot string

// caddyAddSiteCmd adds a basic file server site.
var caddyAddSiteCmd = &cobra.Command{
	Use:   "site <host>",
	Short: "Add a basic file server site block",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return caddy.AppendSite(args[0], siteRoot)
	},
}

func init() {
	caddyAddCmd.AddCommand(caddyAddSiteCmd)
	caddyAddSiteCmd.Flags().StringVar(&siteRoot, "root", "", "Optional site root path")
}

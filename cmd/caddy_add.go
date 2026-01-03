package cmd

import (
	"github.com/spf13/cobra"
)

// caddyAddCmd represents the add command group
var caddyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add entries to the Caddyfile",
}

func init() {
	caddyCmd.AddCommand(caddyAddCmd)
}

package cmd

import (
	"fmt"

	"github.com/phillip-england/konf/internal/caddy"
	"github.com/spf13/cobra"
)

var (
	uploadLimitMax   string
	uploadLimitMatch string
)

// caddyAddUploadLimitCmd adds a request_body max_size block.
var caddyAddUploadLimitCmd = &cobra.Command{
	Use:   "upload-limit <host>",
	Short: "Add a request_body max_size limit",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if uploadLimitMax == "" {
			return fmt.Errorf("--max is required (e.g., 10MB)")
		}
		return caddy.AppendUploadLimit(args[0], uploadLimitMatch, uploadLimitMax)
	},
}

func init() {
	caddyAddCmd.AddCommand(caddyAddUploadLimitCmd)
	caddyAddUploadLimitCmd.Flags().StringVar(&uploadLimitMax, "max", "", "Max upload size (e.g., 10MB)")
	caddyAddUploadLimitCmd.Flags().StringVar(&uploadLimitMatch, "match", "", "Optional path matcher (e.g., /upload/*)")
}

package cmd

import (
	"fmt"

	"github.com/phillip-england/konf/internal/caddy"
	"github.com/spf13/cobra"
)

var (
	rateLimitZone   string
	rateLimitEvents int
	rateLimitWindow string
	rateLimitBurst  int
	rateLimitMatch  string
)

// caddyAddRateLimitCmd adds a rate_limit block.
var caddyAddRateLimitCmd = &cobra.Command{
	Use:   "rate-limit <host>",
	Short: "Add a rate_limit block (plugin required)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if rateLimitEvents <= 0 {
			return fmt.Errorf("--events must be greater than 0")
		}
		if rateLimitWindow == "" {
			return fmt.Errorf("--window is required")
		}
		return caddy.AppendRateLimit(args[0], rateLimitMatch, rateLimitZone, rateLimitEvents, rateLimitWindow, rateLimitBurst)
	},
}

func init() {
	caddyAddCmd.AddCommand(caddyAddRateLimitCmd)
	caddyAddRateLimitCmd.Flags().StringVar(&rateLimitZone, "zone", "default", "Rate limit zone name")
	caddyAddRateLimitCmd.Flags().IntVar(&rateLimitEvents, "events", 100, "Events allowed per window")
	caddyAddRateLimitCmd.Flags().StringVar(&rateLimitWindow, "window", "1m", "Window duration (e.g., 1m, 10s)")
	caddyAddRateLimitCmd.Flags().IntVar(&rateLimitBurst, "burst", 0, "Burst allowance")
	caddyAddRateLimitCmd.Flags().StringVar(&rateLimitMatch, "match", "", "Optional path matcher (e.g., /api/*)")
}

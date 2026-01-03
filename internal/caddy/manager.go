package caddy

import (
	"fmt"
	"os"
)

// GenerateDefault creates a Caddyfile in the current directory
func GenerateDefault(force bool) error {
	filename := "Caddyfile"

	// Check if file already exists
	if _, err := os.Stat(filename); err == nil {
		if !force {
			return fmt.Errorf("%s already exists. Use --force to overwrite", filename)
		}
	}

	// Write the file
	err := os.WriteFile(filename, []byte(DefaultCaddyfile), 0644)
	if err != nil {
		return fmt.Errorf("failed to write Caddyfile: %w", err)
	}

	fmt.Printf("✓ Successfully created %s\n", filename)
	return nil
}
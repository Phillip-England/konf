package caddy

import (
	"fmt"
	"os"
	"strings"
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

// AppendSite adds a basic site block with optional root and a file server.
func AppendSite(host, root string) error {
	block := siteBlock(host, root)
	return appendBlock(block)
}

// AppendReverseProxy adds a site block with a reverse proxy.
func AppendReverseProxy(host, upstream string) error {
	block := reverseProxyBlock(host, upstream)
	return appendBlock(block)
}

// AppendRateLimit adds a rate_limit block, optionally scoped to a path matcher.
func AppendRateLimit(host, match, zone string, events int, window string, burst int) error {
	block := rateLimitBlock(host, match, zone, events, window, burst)
	fmt.Println("NOTE: rate_limit requires a Caddy plugin (e.g., caddy-ratelimit).")
	return appendBlock(block)
}

// AppendUploadLimit adds a request_body max_size block, optionally scoped to a path matcher.
func AppendUploadLimit(host, match, maxSize string) error {
	block := uploadLimitBlock(host, match, maxSize)
	return appendBlock(block)
}

func siteBlock(host, root string) string {
	if root == "" {
		return fmt.Sprintf(`%s {
    file_server
}`, host)
	}

	return fmt.Sprintf(`%s {
    root * %s
    file_server
}`, host, root)
}

func reverseProxyBlock(host, upstream string) string {
	return fmt.Sprintf(`%s {
    reverse_proxy %s
}`, host, upstream)
}

func rateLimitBlock(host, match, zone string, events int, window string, burst int) string {
	rule := fmt.Sprintf(`rate_limit {
        zone %s {
            key {remote_host}
            events %d
            window %s
            burst %d
        }
    }`, zone, events, window, burst)

	if match == "" {
		return fmt.Sprintf(`%s {
    # NOTE: rate_limit requires a Caddy plugin (e.g., caddy-ratelimit).
    %s
}`, host, rule)
	}

	return fmt.Sprintf(`%s {
    # NOTE: rate_limit requires a Caddy plugin (e.g., caddy-ratelimit).
    @rate_limit path %s
    handle @rate_limit {
        %s
    }
}`, host, match, rule)
}

func uploadLimitBlock(host, match, maxSize string) string {
	rule := fmt.Sprintf(`request_body {
        max_size %s
    }`, maxSize)

	if match == "" {
		return fmt.Sprintf(`%s {
    %s
}`, host, rule)
	}

	return fmt.Sprintf(`%s {
    @upload_limit path %s
    handle @upload_limit {
        %s
    }
}`, host, match, rule)
}

func appendBlock(block string) error {
	filename := "Caddyfile"
	contents, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s not found. Run `konf caddy init` first", filename)
		}
		return fmt.Errorf("failed to read %s: %w", filename, err)
	}

	separator := ""
	if len(contents) > 0 {
		separator = "\n"
		if !strings.HasSuffix(string(contents), "\n") {
			separator = "\n\n"
		}
	}

	updated := append(contents, []byte(separator+block+"\n")...)
	if err := os.WriteFile(filename, updated, 0644); err != nil {
		return fmt.Errorf("failed to update %s: %w", filename, err)
	}

	fmt.Printf("✓ Updated %s\n", filename)
	return nil
}

package caddy

// DefaultCaddyfile provides a standard starting point
// It sets up localhost, compression, and a static file server.
const DefaultCaddyfile = `{
    # Global options
    email your-email@example.com
}

localhost:8080 {
    # Set up a file server
    file_server

    # Enable Gzip compression
    encode gzip

    # Basic logging
    log {
        output file ./caddy.log
    }

    # Example reverse proxy (commented out)
    # reverse_proxy 127.0.0.1:3000
}
`
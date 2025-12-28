package configs

// Default returns a default Config instance.
func Default() Config {
	return Config{
		Logger: LoggerConfig{
			Level: "info",
			JSON:  false,
		},
		TLS: TLSConfig{
			Enable:   true,
			KeyPath:  "/etc/flap/tls/tls.key",
			CertPath: "/etc/flap/tls/tls.crt",
		},
	}
}

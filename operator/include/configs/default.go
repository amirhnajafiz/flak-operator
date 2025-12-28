package configs

// Default returns a default viper instance.
func Default() Config {
	return Config{
		LogLevel: "info",
		JSONLog:  false,
		TLS: struct {
			Enable   bool   "koanf:\"enable\""
			CertPath string "koanf:\"cert_path\""
			KeyPath  string "koanf:\"key_path\""
		}{
			Enable: false,
		},
	}
}

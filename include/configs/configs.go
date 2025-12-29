package configs

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/tidwall/pretty"
)

const (
	// Prefix indicates environment variables prefix.
	Prefix = "FLAK_"
)

// LoggerConfig holds the logger tune parameters.
type LoggerConfig struct {
	Level string `koanf:"level"`
	JSON  bool   `koanf:"json"`
}

// TLSConfig holds the tls parameters.
type TLSConfig struct {
	Enable   bool   `koanf:"enable"`
	CertPath string `koanf:"cert_path"`
	KeyPath  string `koanf:"key_path"`
}

// Config hold the operator tune parameters.
type Config struct {
	Logger LoggerConfig `koanf:"logger"`
	TLS    TLSConfig    `koanf:"tls"`
}

// LoadConfigs reads the config parameters from config.yml and env variables.
func LoadConfigs() *Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from file
	err := k.Load(structs.Provider(Default(), "koanf"), nil)
	if err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	err = k.Load(file.Provider("config.yml"), yaml.Parser())
	if err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	// load environment variables
	err = k.Load(env.Provider(Prefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, Prefix)), "__", ".")
	}), nil)
	if err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	err = k.Unmarshal("", &instance)
	if err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	indent, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		log.Fatalf("error marshaling config to json: %s", err)
	}

	indent = pretty.Color(indent, nil)
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	return &instance
}

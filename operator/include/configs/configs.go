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

const Prefix = "flap_"

// Config hold the operator tune parameters.
// all configuration parameters must be set as environment variables or as a `.env` file.
type Config struct {
	LogLevel string `koanf:"log_level"`
	JSONLog  bool   `koanf:"json_log"`
	TLS      struct {
		Enable   bool   `koanf:"enable"`
		CertPath string `koanf:"cert_path"`
		KeyPath  string `koanf:"key_path"`
	} `koanf:"tls"`
}

// LoadConfigs reads the env variables into a Config struct.
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

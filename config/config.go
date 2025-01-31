package config

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type Config struct {
	DBPath string `yaml:"db_path"`
	Port   string `yaml:"port"`
}

func LoadConfig() *Config {
	cfg := &Config{}

	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Failed to read config file %v", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		log.Fatalf("Failed to parse config file %v", err)
	}

	return cfg
}

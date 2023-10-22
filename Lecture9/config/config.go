package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DatabaseURL string `yaml:"database"`
}

func LoadConfig(configPath string) (*Config, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

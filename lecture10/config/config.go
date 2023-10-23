package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type HttpServerConfig struct {
	Port            int           `mapstructure:"port"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown-timeout"`
}

type AppConfig struct {
	Database   DatabaseConfig `yaml:"database"`
	HttpServer struct {
		Port            int           `yaml:"port"`
		ShutdownTimeout time.Duration `yaml:"shutdown-timeout"`
	} `yaml:"http-server"`
}

func LoadConfig(configPath string) (*AppConfig, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &AppConfig{}
	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

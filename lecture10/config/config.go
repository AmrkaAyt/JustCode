// config/config.go
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	HttpServer struct {
		Port int `yaml:"Port"`
	} `yaml:"HttpServer"`
}

func LoadConfig(configPath string) (*AppConfig, error) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}

	return &appConfig, nil
}

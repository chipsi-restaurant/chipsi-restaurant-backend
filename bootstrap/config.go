package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	SslMode  bool
}

func LoadConfig() (*Config, error) {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		return nil, fmt.Errorf("CONFIG_PATH не задан")
	}

	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

package config

import "github.com/prashant-kumar-src/goms-boilerplate/utils/logger"

type Config struct {
	LoggerClient   logger.LoggerClient
	Name           string
	Version        string
	AllowedOrigins []string
}

func GetConfig() *Config {
	return &Config{
		LoggerClient:   logger.Zerolog,
		Name:           "Users",
		Version:        "1.0.0",
		AllowedOrigins: []string{},
	}
}

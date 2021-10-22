package services

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// NewConfig - Формируем конфиг
func NewConfig(path string) (config *Config) {
	if _, err := toml.DecodeFile(fmt.Sprintf("%sconfig.toml", path), &config); err != nil {
		log.Fatalln(err)
	}
	config.Application.Path = path
	return config
}

// Config - Конфигурационный файл
type Config struct {
	Application struct {
		CurrentVersion string
		FirstMessage   string
		Status         string
		Port           string
		Path           string
	}

	Logger struct {
		File string
	}

	Mailer struct {
		Host string
	}

	Redis struct {
		Host, Port, Pass string
		Size             int
	}

	Session struct {
		Host, Key, Path string
		Secure          bool
	}

	Storage struct {
		Driver       string
		DSN          string
		MaxOpenConns int
	}

	CORS struct {
		AllowedOrigins   []string
		AllowedMethods   []string
		AllowedHeaders   []string
		ExposedHeaders   []string
		AllowCredentials bool
		MaxAge           int
	}
}

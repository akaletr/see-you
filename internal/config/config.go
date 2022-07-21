package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

// Config описывает конфиг приложения
type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
	DatabaseDSN   string `env:"DATABASE_DSN"`
	SecretKey     string `env:"SECRET_KEY"`
}

// GetConfig возвращает ссылку на объект конфига
func GetConfig() (*Config, error) {
	cfg := Config{
		ServerAddress: "localhost:8080",
		BaseURL:       "http://localhost:8080",
		DatabaseDSN:   "",
		SecretKey:     "akaletr",
	}

	// парсим окружение в конфиг
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	// если конфиг задается флагами - перезаписываем поля конфига
	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "base URL")
	flag.StringVar(&cfg.DatabaseDSN, "s", cfg.DatabaseDSN, "database DSN")
	flag.StringVar(&cfg.SecretKey, "d", cfg.SecretKey, "secret key")

	return &Config{}, nil
}

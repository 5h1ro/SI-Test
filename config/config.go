package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		DB   `yaml:"db"`
	}

	// App -.
	App struct {
		Name       string `env-required:"true" env:"APP_NAME"`
		Version    string `env-required:"true" env:"APP_VERSION"`
		JwtSignKey string `env-required:"true" env:"JWT_SIGN_KEY"`
		LogLevel   string `env-required:"true" env:"LOG_LEVEL"`
	}

	// HTTP -.
	HTTP struct {
		Port    string `env-required:"true" env:"HTTP_PORT"`
		BaseUrl string `env-required:"true" env:"BASE_URL"`
	}

	// DB -.
	DB struct {
		PoolMax  int    `env-required:"true" env:"DB_POOL_MAX"`
		User     string `env-required:"true" env:"DB_USER"`
		Password string `env-required:"true" env:"DB_PASSWORD"`
		Host     string `env-required:"true" env:"DB_HOST"`
		Port     int    `env-required:"true" env:"DB_PORT"`
		Database string `env-required:"true" env:"DB_DATABASE"`
	}
)

// NewConfig returns app config.
func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

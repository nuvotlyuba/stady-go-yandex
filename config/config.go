package config

import (
	"log"

	"github.com/caarlos0/env"
)

var BaseURL string = `http://localhost:8080`
var ServerAddress string = `:8080`

const AppLevel = DEVELOPMENT

type Config struct {
	BaseURL       string `env:"BASE_URL"`
	ServerAddress string `env:"SERVER_ADDRESS"`
	AppLevel      Level  `env:"APP_LEVEL"`
}

func New() *Config {
	return &Config{
		BaseURL:       BaseURL,
		ServerAddress: ServerAddress,
		AppLevel:      AppLevel,
	}
}

func (c Config) LoadConfig() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

}

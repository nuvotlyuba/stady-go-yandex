package config

import (
	"log"

	"github.com/caarlos0/env"
)

var BaseURL string = `http://localhost:8080`
var ServerAddress string = `:8080`
var FileStoragePath string = `tmp/file.json`

const AppLevel = DEVELOPMENT

type Config struct {
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	AppLevel        Level  `env:"APP_LEVEL" envDefault:"development"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"tmp/file.json"`
}

func New() *Config {
	return &Config{
		BaseURL:         BaseURL,
		ServerAddress:   ServerAddress,
		AppLevel:        AppLevel,
		FileStoragePath: FileStoragePath,
	}
}

func (c Config) LoadConfig() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

}

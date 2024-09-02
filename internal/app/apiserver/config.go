package apiserver

import "github.com/nuvotlyuba/study-go-yandex/config"

type APIConfig struct {
	ServerAddress string
	AppLevel      config.Level
}

func NewConfig() *APIConfig {
	return &APIConfig{
		ServerAddress: config.ServerAddress,
		AppLevel:      config.AppLevel,
	}
}

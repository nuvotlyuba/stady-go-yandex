package apiserver

import "github.com/nuvotlyuba/study-go-yandex/config"

type APIConfig struct {
	ServerAddress string
}

func NewConfig() *APIConfig {
	return &APIConfig{
		ServerAddress: config.ServerAddress,
	}
}

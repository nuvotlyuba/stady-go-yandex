package main

import (
	"flag"

	"github.com/nuvotlyuba/study-go-yandex/config"
	"github.com/nuvotlyuba/study-go-yandex/internal/utils"
)

func parseFlags() {
	flag.StringVar(&config.ServerAddress, "a", config.ServerAddress, "Server address host:port")
	flag.StringVar(&config.BaseURL, "b", config.BaseURL, "Base URL host:port")
	flag.StringVar(&config.FileStoragePath, "f", config.FileStoragePath, "File storage path")
	flag.Parse()

	if config.FileStoragePath != "" {
		utils.MakeDir(config.FileStoragePath)
	}
}

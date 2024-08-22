package main

import (
	"flag"

	"github.com/nuvotlyuba/study-go-yandex/config"
)

func parseFlags() {
	flag.StringVar(&config.ServerAddress, "a", config.ServerAddress, "Server address host:port")
	flag.StringVar(&config.BaseURL, "b", config.BaseURL, "Base URL host:port")
	flag.Parse()
}

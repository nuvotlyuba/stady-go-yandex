package main

import (
	"flag"

	"github.com/nuvotlyuba/study-go-yandex/config"
)

func parseFlags() {
	flag.StringVar(&config.ServerAddress, "a", "localhost:8080", "Server address host:port")
	flag.StringVar(&config.BaseURL, "b", "http://localhost:8080", "Base URL host:port")
	flag.Parse()
}

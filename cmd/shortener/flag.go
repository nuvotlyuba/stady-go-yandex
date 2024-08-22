package main

import (
	"flag"

	"github.com/nuvotlyuba/study-go-yandex/config"
)

func parseFlags() {
	flag.StringVar(&config.ServerAddress, "a", "", "Server address host:port")
	flag.StringVar(&config.BaseURL, "b", "", "Base URL host:port")
	flag.Parse()
}

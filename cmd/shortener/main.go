package main

import (
	"github.com/nuvotlyuba/study-go-yandex/config"
	"github.com/nuvotlyuba/study-go-yandex/internal/app/apiserver"
)

func main() {
	parseFlags()

	c := config.New()
	c.LoadConfig()

	cfg := apiserver.NewConfig()
	s := apiserver.New(cfg)
	if err := s.Start(); err != nil {
		panic(err)
	}
}

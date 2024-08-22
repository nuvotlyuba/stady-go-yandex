package config

var BaseURL string = `http://localhost:8080`
var ServerAddress string = `:8080`

type Config struct {
	BaseURL       string
	ServerAddress string
}

func (c Config) New() *Config {
	return &Config{}
}

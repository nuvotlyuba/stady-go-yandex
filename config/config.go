package config

var BaseURL string = `http://localhost:8000`
var ServerAddress string = `:8000`

type Config struct {
	BaseURL       string
	ServerAddress string
}

func (c Config) New() *Config {
	return &Config{}
}

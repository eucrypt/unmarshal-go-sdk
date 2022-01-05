package pkg

import "net/http"

type Config struct {
	AuthKey    string
	HttpClient *http.Client
}

type ConfigOptions func(config *Config)

func NewConfig(options ...ConfigOptions) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

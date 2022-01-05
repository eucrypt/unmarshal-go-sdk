package pkg

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"net/http"
)

type Config struct {
	AuthKey     string
	HttpClient  *http.Client
	Environment constants.Environment
}

type ConfigOptions func(config *Config)

func NewConfig(authKey string, options ...ConfigOptions) Config {
	config := Config{AuthKey: authKey}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithHttpClient(httpClient http.Client) ConfigOptions {
	return func(cfg *Config) {
		cfg.HttpClient = &httpClient
	}
}

func WithEnvironment(env constants.Environment) ConfigOptions {
	return func(cfg *Config) {
		cfg.Environment = env
	}
}

func setDefaults(config *Config) {
	if len(config.Environment) == 0 {
		config.Environment = constants.Prod
	}
}

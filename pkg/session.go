package pkg

import (
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_details"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_price"
)

type Session struct {
	Config Config
	Client httpclient.Request
}

type Unmarshal struct {
	token_details.TokenStore
	token_price.PriceStore
}

func NewWithConfig(config Config) Unmarshal {
	setDefaults(&config)
	httpClient := httpclient.NewHttpJSONClient(config.Environment.GetEndpoint())
	if config.HttpClient != nil {
		httpClient.HttpClient = config.HttpClient
	}
	httpClient.DefaultQuery = map[string]string{"auth_key": config.AuthKey}
	sess := Session{Config: config, Client: httpClient}
	return Unmarshal{
		TokenStore: token_details.New(sess),
		PriceStore: token_price.New(sess),
	}
}

func NewWithOptions(authKey string, options ...ConfigOptions) Unmarshal {
	config := NewConfig(authKey, options...)
	sess := Session{Config: config}
	return Unmarshal{
		TokenStore: token_details.New(sess),
		PriceStore: token_price.New(sess),
	}
}

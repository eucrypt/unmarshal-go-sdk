package pkg

import (
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_details"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_price"
)

const (
	productionEndpoint = "https://api.unmarshal.com/"
	stagingEndpoint    = "https://stg-api.unmarshal.com/"
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
	sess := Session{Config: config}
	return Unmarshal{
		TokenStore: token_details.New(sess),
		PriceStore: token_price.New(sess),
	}
}

func NewWithOptions(options ...ConfigOptions) Unmarshal {
	config := NewConfig(options...)
	sess := Session{Config: config}
	return Unmarshal{
		TokenStore: token_details.New(sess),
		PriceStore: token_price.New(sess),
	}
}

package pkg

import (
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	tokenDetails "github.com/eucrypt/unmarshal-go-sdk/pkg/token_details"
)

type Session struct {
	config Config
	client httpclient.Request
}

type Unmarshal struct {
	tokenDetails.TokenStore
	PriceStore
}

func New(config Config) Unmarshal {
	sess := Session{}
	return Unmarshal{
		TokenStore: tokenDetails.New(sess),
	}
}

type PriceStore interface {
	GetPriceAtInstant(contractAddress string, timestamp int64)
}

package pkg

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/assets"
	conf "github.com/eucrypt/unmarshal-go-sdk/pkg/config"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/nft_details"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/notifications"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/protocol_details"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_details"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_price"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details"
)

type Unmarshal struct {
	token_details.TokenDetails
	token_price.PriceStore
	assets.Assets
	nft_details.NftDetails
	transaction_details.TransactionDetails
	protocol_details.ProtocolDetails
	notifications.Notification
}

func NewWithConfig(config conf.Config) Unmarshal {
	conf.SetDefaults(&config)
	httpClient := httpclient.NewHttpJSONClient(config.Environment.GetEndpoint())
	if config.HttpClient != nil {
		httpClient.HttpClient = config.HttpClient
	}
	httpClient.DefaultQuery = map[string]string{"auth_key": config.AuthKey}
	sess := session.Session{Config: config, Client: httpClient}
	return Unmarshal{
		TokenDetails:       token_details.New(sess),
		PriceStore:         token_price.New(sess),
		Assets:             assets.New(sess),
		NftDetails:         nft_details.New(sess),
		TransactionDetails: transaction_details.New(sess),
		ProtocolDetails:    protocol_details.New(sess),
	}
}

func NewWithOptions(authKey string, options ...conf.ConfigOptions) Unmarshal {
	config := conf.NewConfig(authKey, options...)
	sess := session.Session{Config: config}
	return Unmarshal{
		TokenDetails:       token_details.New(sess),
		PriceStore:         token_price.New(sess),
		Assets:             assets.New(sess),
		NftDetails:         nft_details.New(sess),
		TransactionDetails: transaction_details.New(sess),
		ProtocolDetails:    protocol_details.New(sess),
	}
}

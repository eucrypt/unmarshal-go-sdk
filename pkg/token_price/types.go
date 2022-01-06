package token_price

import "github.com/eucrypt/unmarshal-go-sdk/pkg/constants"

type PriceStore interface {
	GetPriceAtInstant(contractAddress string, chain constants.Chain, timestamp int64) (TokenPrice, error)
	GetCurrentPrice(contractAddress string, chain constants.Chain) (resp TokenPrice, err error)
	GetGainers(chain constants.Chain) (resp TokenPriceList, err error)
	GetLosers(chain constants.Chain) (resp TokenPriceList, err error)
}

type TokenPrice struct {
	TokenId   string `json:"tokenId"`
	Timestamp string `json:"timestamp"`
	Price     string `json:"price"`
}
type TokenDetailsWithPrice struct {
	Name          string `json:"name"`
	Contract      string `json:"contract"`
	Decimal       int    `json:"decimal"`
	Logo          string `json:"logo"`
	Symbol        string `json:"symbol"`
	CurrentPrice  string `json:"current_price"`
	PercentChange string `json:"percent_change"`
	DayLow        string `json:"day_low"`
	DayHigh       string `json:"day_high"`
}

type TokenPriceList []TokenDetailsWithPrice

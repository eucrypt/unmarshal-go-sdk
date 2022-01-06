package token_price

import "github.com/eucrypt/unmarshal-go-sdk/pkg/constants"

type PriceStore interface {
	GetPriceAtInstant(chain constants.Chain, contractAddress string, timestamp int64) (TokenPrice, error)
	GetCurrentPrice(chain constants.Chain, contractAddress string) (resp TokenPrice, err error)
	GetGainers(chain constants.Chain) (resp TokenDetailsResp, err error)
	GetLosers(chain constants.Chain) (resp TokenDetailsResp, err error)
	GetLPTokens(chain constants.Chain, lptoken string) (resp TokenListWithPrice, err error)
	GetTokensPrice(chain constants.Chain, tokenList []string) (resp TokenListWithPrice, err error)
	GetPriceWithSymbol(symbol string) (resp PriceWithSymbolResp, err error)
}

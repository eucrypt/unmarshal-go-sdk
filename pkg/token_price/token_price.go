package token_price

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_price/types"
)

type PriceStore interface {
	GetPriceAtInstant(chain constants.Chain, contractAddress string, timestamp int64) (types.TokenPrice, error)
	GetCurrentPrice(chain constants.Chain, contractAddress string) (resp types.TokenPrice, err error)
	GetGainers(chain constants.Chain) (resp types.TokenDetailsResp, err error)
	GetLosers(chain constants.Chain) (resp types.TokenDetailsResp, err error)
	GetLPTokens(chain constants.Chain, lptoken string) (resp types.TokenListWithPrice, err error)
	GetTokensPrice(chain constants.Chain, tokenList []string) (resp types.TokenListWithPrice, err error)
	GetPriceWithSymbol(symbol string) (resp types.PriceWithSymbolResp, err error)
}

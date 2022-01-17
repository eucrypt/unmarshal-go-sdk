package token_price

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_price/types"
)

type PriceStore interface {
	GetTokenPriceAtInstant(chain constants.Chain, contractAddress string, timestamp int64) (types.TokenPrice, error)
	GetTokenCurrentPrice(chain constants.Chain, contractAddress string) (resp types.TokenPrice, err error)
	GetTopGainers(chain constants.Chain) (resp types.TokenDetailsResp, err error)
	GetTopLosers(chain constants.Chain) (resp types.TokenDetailsResp, err error)
	GetLPTokens(chain constants.Chain, lptoken string) (resp types.TokenListWithPrice, err error)
	GetMultipleTokenPrice(chain constants.Chain, tokenList []string) (resp types.TokenListWithPrice, err error)
	GetTokenPriceBySymbol(symbol string) (resp types.PriceWithSymbolResp, err error)
}

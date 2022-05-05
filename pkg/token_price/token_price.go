package token_price

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_price/types"
)

type PriceStore interface {
	GetTokenPrice(chain constants.Chain, contractAddress string, options *types.GetPriceOptions) (types.TokenPrice, error)
	GetTokenPriceBySymbol(symbol string, options *types.GetPriceWithSymbolOptions) (resp types.PriceWithSymbolResp, err error)
	GetTopGainers(chain constants.Chain, options *types.GetTopGainersOptions) (resp types.TokenDetailsResp, err error)
	GetTopLosers(chain constants.Chain, options *types.GetTopLosersOptions) (resp types.TokenDetailsResp, err error)
	GetLPTokens(chain constants.Chain, lptoken string) (resp types.TokenListWithPrice, err error)
	GetMultipleTokenPrice(chain constants.Chain, tokenList []string) (resp types.TokenListWithPrice, err error)
}

package constants

import "errors"

type Chain string
type APIName string

const (
	ETH     Chain = "ethereum"
	BSC     Chain = "bsc"
	MATIC   Chain = "matic"
	XDC     Chain = "xinfin"
	SOL     Chain = "solana"
	ZILLIQA Chain = "zilliqa"
	HUOBI   Chain = "heco"
)

const (
	PS_GetPriceWithAddress    APIName = "v1/pricestore/chain/:chain/:address"
	PS_GetTokensPrice         APIName = "v1/pricestore/chain/:chain/tokens"
	PS_GetLpTokenPrice        APIName = "v1/pricestore/chain/:chain/lptokens"
	PS_GetLosers              APIName = "v1/pricestore/chain/:chain/losers"
	PS_GetGainers             APIName = "v1/pricestore/chain/:chain/gainers"
	PS_GetPriceWithSymbol     APIName = "v1/pricestore/:symbol"
	TS_GetDetailsWithContract APIName = "v1/tokenstore/token/address/:address"
	TS_GetTokenList           APIName = "v1/tokenstore/token/all"
	TS_GetTokenWithSymbol     APIName = "v1/tokenstore/token/symbol/:symbol"
	Assets_GetAssets          APIName = "v1/:chain/address/:address/assets"
)

//This should be manually changed when a new chain starts being supported
var allChains = map[Chain]bool{ETH: true, BSC: true, MATIC: true, XDC: true, SOL: true, ZILLIQA: true, HUOBI: true}

var priceStoreSupported = map[Chain]bool{ETH: true, BSC: true, MATIC: true}

var allowedCallersByAPI = map[APIName]map[Chain]bool{
	PS_GetPriceWithAddress: priceStoreSupported,
	PS_GetTokensPrice:      priceStoreSupported,
	PS_GetLpTokenPrice:     priceStoreSupported,
	PS_GetLosers:           priceStoreSupported,
	PS_GetGainers:          priceStoreSupported,
	Assets_GetAssets:       allChains,
}

//String returns the string specific version of the chain
func (c Chain) String() string {
	return string(c)
}

//SupportsChain Allows a caller to know if a chain specific API supports a passed valid chain
func (api APIName) SupportsChain(chain Chain) bool {
	if allowedCallersByAPI[api] == nil {
		return false
	}

	return allowedCallersByAPI[api][chain]
}

func (api APIName) GetURI() string {
	return string(api)
}

var UnsupportedChainError = errors.New("unsupported Chain for API")

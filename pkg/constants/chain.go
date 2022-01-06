package constants

import "errors"

type Chain string
type APIName string

const (
	ETH    Chain = "eth"
	BSC          = "bsc"
	MATIC        = "matic"
	XDC          = "xinfin"
	SOLANA       = "solana"
)

const (
	GetPriceWithAddress APIName = "GetPriceWithAddress"
	GetTokensPrice              = "GetTokensPrice"
	GetLpTokenPrice             = "GetLpTokenPrice"
	GetLosers                   = "GetLosers"
	GetGainers                  = "GetGainers"
)

var allChainsTrue = map[Chain]bool{ETH: true, BSC: true, MATIC: true, XDC: true, SOLANA: true}
var allowedCallersByAPI = map[APIName]map[Chain]bool{
	GetPriceWithAddress: allChainsTrue,
	GetTokensPrice:      allChainsTrue,
	GetLpTokenPrice:     allChainsTrue,
	GetLosers:           allChainsTrue,
	GetGainers:          allChainsTrue,
}

func (c Chain) String() string {
	return string(c)
}

//IsAllowedToCallOnChain Allows a caller to know if a chain specific API supports a passed valid chain
func (api APIName) IsAllowedToCallOnChain(chain Chain) bool {
	if allowedCallersByAPI[api] == nil {
		return false
	}

	return allowedCallersByAPI[api][chain]
}

var UnsupportedChainError = errors.New("unsupported Chain for API")

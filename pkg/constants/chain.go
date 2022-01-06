package constants

import "errors"

type Chain string
type APIName string

const (
	ETH    Chain = "ethereum"
	BSC    Chain = "bsc"
	MATIC  Chain = "matic"
	XDC    Chain = "xinfin"
	SOLANA Chain = "solana"
)

const (
	GetPriceWithAddress APIName = "GetPriceWithAddress"
	GetTokensPrice      APIName = "GetTokensPrice"
	GetLpTokenPrice     APIName = "GetLpTokenPrice"
	GetLosers           APIName = "GetLosers"
	GetGainers          APIName = "GetGainers"
)

//This should be manually changed when a new chain starts being supported
var allChainsTrue = map[Chain]bool{ETH: true, BSC: true, MATIC: true, XDC: true, SOLANA: true}

var allowedCallersByAPI = map[APIName]map[Chain]bool{
	GetPriceWithAddress: allChainsTrue,
	GetTokensPrice:      allChainsTrue,
	GetLpTokenPrice:     allChainsTrue,
	GetLosers:           allChainsTrue,
	GetGainers:          allChainsTrue,
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

var UnsupportedChainError = errors.New("unsupported Chain for API")

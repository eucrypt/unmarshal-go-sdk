package constants

import "errors"

type Chain string

const (
	ETH       Chain = "ethereum"
	BSC       Chain = "bsc"
	MATIC     Chain = "matic"
	XDC       Chain = "xinfin"
	SOL       Chain = "solana"
	ZILLIQA   Chain = "zilliqa"
	HUOBI     Chain = "heco"
	AVALANCHE Chain = "avalanche"
)

//This should be manually changed when a new chain starts being supported
var allChains = map[Chain]bool{
	ETH:       true,
	BSC:       true,
	MATIC:     true,
	XDC:       true,
	SOL:       true,
	ZILLIQA:   true,
	HUOBI:     true,
	AVALANCHE: true,
}

var priceStoreSupported = map[Chain]bool{ETH: true, BSC: true, MATIC: true}
var priceStoreSupportedWithAvax = map[Chain]bool{ETH: true, BSC: true, MATIC: true, AVALANCHE: true}

var nftEVMSupport = map[Chain]bool{ETH: true, BSC: true, MATIC: true, AVALANCHE: true}

//String returns the string specific version of the chain
func (c Chain) String() string {
	return string(c)
}

var UnsupportedChainError = errors.New("unsupported chain for API")

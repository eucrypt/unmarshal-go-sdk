package constants

import (
	"errors"
)

type Chain string

//goland:noinspection GoSnakeCaseUsage
const (
	ETH           Chain = "ethereum"
	ETH_RINKEBY   Chain = "rinkeby-testnet"
	BSC           Chain = "bsc"
	BSC_TESTNET   Chain = "bsc-testnet"
	MATIC         Chain = "matic"
	MATIC_TESTNET Chain = "matic-testnet"
	XDC           Chain = "xinfin"
	SOL           Chain = "solana"
	ZILLIQA       Chain = "zilliqa"
	HUOBI         Chain = "heco"
	AVALANCHE     Chain = "avalanche"
	OPTIMISM      Chain = "optimism"
	ARBITRUM      Chain = "arbitrum"
	CELO          Chain = "celo"
	FANTOM        Chain = "fantom"
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
	OPTIMISM:  true,
	ARBITRUM:  true,
	CELO:      true,
	FANTOM:    true,
}

var priceStoreSupported = map[Chain]bool{ETH: true, BSC: true, MATIC: true}

var rawTxnSupported = map[Chain]bool{ETH: true, ETH_RINKEBY: true, BSC: true, BSC_TESTNET: true, MATIC: true,
	MATIC_TESTNET: true, OPTIMISM: true, AVALANCHE: true, ARBITRUM: true, CELO: true, FANTOM: true}

var nftEVMSupport = map[Chain]bool{ETH: true, BSC: true, MATIC: true, AVALANCHE: true}

//String returns the string specific version of the chain
func (c Chain) String() string {
	return string(c)
}

var UnsupportedChainError = errors.New("unsupported chain for API")

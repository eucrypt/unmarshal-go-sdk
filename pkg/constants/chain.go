package constants

import (
	"errors"
)

type Chain string

//goland:noinspection GoSnakeCaseUsage
const (
	ARBITRUM      Chain = "arbitrum"
	AVALANCHE     Chain = "avalanche"
	BSC           Chain = "bsc"
	BSC_TESTNET   Chain = "bsc-testnet"
	CELO          Chain = "celo"
	CRONOS        Chain = "cronos"
	ETH           Chain = "ethereum"
	ETH_RINKEBY   Chain = "rinkeby-testnet"
	FANTOM        Chain = "fantom"
	FUSE          Chain = "fuse"
	HUOBI         Chain = "heco"
	KLAYTN        Chain = "klaytn"
	MATIC         Chain = "matic"
	MATIC_TESTNET Chain = "matic-testnet"
	OPTIMISM      Chain = "optimism"
	SOL           Chain = "solana"
	VELAS         Chain = "velas"
	XDC           Chain = "xinfin"
	ZILLIQA       Chain = "zilliqa"
	MOONBEAM      Chain = "moonbeam"
	METIS         Chain = "metis"
	AURORA        Chain = "aurora"
	MaticSupernet Chain = "matic-supernet"
	ZKEVM         Chain = "zkevm"
)

// This should be manually changed when a new chain starts being supported
var allChains = map[Chain]bool{
	ARBITRUM:      true,
	AVALANCHE:     true,
	BSC:           true,
	CELO:          true,
	CRONOS:        true,
	ETH:           true,
	FANTOM:        true,
	FUSE:          true,
	HUOBI:         true,
	KLAYTN:        true,
	MATIC:         true,
	OPTIMISM:      true,
	SOL:           true,
	VELAS:         true,
	XDC:           true,
	ZILLIQA:       true,
	MOONBEAM:      true,
	METIS:         true,
	AURORA:        true,
	MaticSupernet: true,
	ZKEVM:         true,
}

var priceStoreSupported = map[Chain]bool{
	ARBITRUM: true,
	BSC:      true,
	ETH:      true,
	MATIC:    true,
}

var rawTxnSupported = map[Chain]bool{ETH: true, ETH_RINKEBY: true, BSC: true, BSC_TESTNET: true, MATIC: true,
	MATIC_TESTNET: true, OPTIMISM: true, AVALANCHE: true, ARBITRUM: true, CELO: true, FANTOM: true, KLAYTN: true,
	FUSE: true, CRONOS: true, VELAS: true, MOONBEAM: true, METIS: true, AURORA: true, MaticSupernet: true, ZKEVM: true}

var nftEVMSupport = map[Chain]bool{ETH: true, BSC: true, MATIC: true, AVALANCHE: true, VELAS: true, KLAYTN: true, FUSE: true, CRONOS: true}

// String returns the string specific version of the chain
func (c Chain) String() string {
	return string(c)
}

var UnsupportedChainError = errors.New("unsupported chain for API")

package constants

import "errors"

type Chain string

const (
	ETH     Chain = "ethereum"
	BSC     Chain = "bsc"
	MATIC   Chain = "matic"
	XDC     Chain = "xinfin"
	SOL     Chain = "solana"
	ZILLIQA Chain = "zilliqa"
	HUOBI   Chain = "heco"
)

//This should be manually changed when a new chain starts being supported
var allChains = map[Chain]bool{ETH: true, BSC: true, MATIC: true, XDC: true, SOL: true, ZILLIQA: true, HUOBI: true}

var priceStoreSupported = map[Chain]bool{ETH: true, BSC: true, MATIC: true}

var nftEVMSupport = map[Chain]bool{ETH: true, BSC: true, MATIC: true}

var allowedCallersByChain = map[APIName]map[Chain]bool{
	PS_GetPriceWithAddress: priceStoreSupported,
	PS_GetTokensPrice:      priceStoreSupported,
	PS_GetLpTokenPrice:     priceStoreSupported,
	PS_GetLosers:           priceStoreSupported,
	PS_GetGainers:          priceStoreSupported,
	ASSETS_GetAssets:       allChains,
	NFT_GetAssets:          {ETH: true, BSC: true, MATIC: true, SOL: true},
	NFT_GetTxns:            nftEVMSupport,
	NFT_GetDetailsWithID:   nftEVMSupport,
	NFT_GetHoldersByID:     nftEVMSupport,
	TXN_GetTokenTxns:       {ETH: true, BSC: true, MATIC: true, SOL: true, ZILLIQA: true, XDC: true},
	TXN_GetTxnDetails:      {ETH: true, BSC: true, MATIC: true, SOL: true, XDC: true},
	TXN_GetTokenTxnsV2:     {ETH: true, BSC: true, XDC: true},
}

//String returns the string specific version of the chain
func (c Chain) String() string {
	return string(c)
}

var UnsupportedChainError = errors.New("unsupported chain for API")

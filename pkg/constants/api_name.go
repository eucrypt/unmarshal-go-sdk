package constants

type APIName string

//goland:noinspection GoSnakeCaseUsage
const (
	PS_GetPriceWithAddress       APIName = "v1/pricestore/chain/:chain/:address"
	PS_GetTokensPrice            APIName = "v1/pricestore/chain/:chain/tokens"
	PS_GetLpTokenPrice           APIName = "v1/pricestore/chain/:chain/lptokens"
	PS_GetLosers                 APIName = "v1/pricestore/chain/:chain/losers"
	PS_GetGainers                APIName = "v1/pricestore/chain/:chain/gainers"
	PS_GetPriceBySymbol          APIName = "v1/pricestore/:symbol"
	TS_GetDetailsByContract      APIName = "v1/tokenstore/token/address/:address"
	TS_GetTokenList              APIName = "v1/tokenstore/token/all"
	TS_GetTokenBySymbol          APIName = "v1/tokenstore/token/symbol/:symbol"
	ASSETS_GetTokenAssets        APIName = "v1/:chain/address/:address/assets"
	ASSETS_GetProfitsAndLoss     APIName = "v2/:chain/address/:address/userData"
	NFT_GetNFTAssets             APIName = "v1/:chain/address/:address/nft-assets"
	NFT_GetTxns                  APIName = "v1/:chain/address/:address/nft-transactions"
	NFT_GetDetailsByID           APIName = "v1/:chain/address/:address/details"
	NFT_GetHoldersByID           APIName = "v1/:chain/address/:address/nftholders"
	TXN_GetTokenTxns             APIName = "v1/:chain/address/:address/transactions"
	TXN_GetTxnDetails            APIName = "v1/:chain/transactions/:txnID"
	TXN_GetTokenTxnsV2           APIName = "v2/:chain/address/:address/transactions"
	TXN_GetRawTransactionDetails APIName = "v3/:chain/address/:address/transactions"
	PROTO_GetPositions           APIName = "v2/protocols/:protocol/address/:address/positions"
	PROTO_GetPairs               APIName = "v2/protocols/:protocol/pairs"
)

var allowedCallersByChain = map[APIName]map[Chain]bool{
	PS_GetPriceWithAddress:       priceStoreSupported,
	PS_GetTokensPrice:            priceStoreSupported,
	PS_GetLpTokenPrice:           priceStoreSupported,
	PS_GetLosers:                 priceStoreSupported,
	PS_GetGainers:                priceStoreSupported,
	ASSETS_GetTokenAssets:        allChains,
	ASSETS_GetProfitsAndLoss:     priceStoreSupported,
	NFT_GetNFTAssets:             {ETH: true, BSC: true, MATIC: true, AVALANCHE: true, SOL: true},
	NFT_GetTxns:                  nftEVMSupport,
	NFT_GetDetailsByID:           nftEVMSupport,
	NFT_GetHoldersByID:           nftEVMSupport,
	TXN_GetTokenTxns:             {ETH: true, BSC: true, MATIC: true, SOL: true, ZILLIQA: true, AVALANCHE: true, XDC: true, OPTIMISM: true, ARBITRUM: true},
	TXN_GetTxnDetails:            {ETH: true, BSC: true, MATIC: true, SOL: true, AVALANCHE: true, XDC: true, OPTIMISM: true, ARBITRUM: true},
	TXN_GetTokenTxnsV2:           {ETH: true, BSC: true, MATIC: true, AVALANCHE: true, XDC: true, OPTIMISM: true, ARBITRUM: true},
	TXN_GetRawTransactionDetails: rawTxnSupported,
}

//SupportsChain Allows a caller to know if a chain specific API supports a passed valid chain
func (api APIName) SupportsChain(chain Chain) bool {
	if allowedCallersByChain[api] == nil {
		return false
	}

	return allowedCallersByChain[api][chain]
}

//SupportsProtocol Allows a caller to know if a protocol specific API supports a passed valid chain
func (api APIName) SupportsProtocol(protocol Protocol) bool {
	if allowedCallersByProtocol[api] == nil {
		return false
	}
	return allowedCallersByProtocol[api][protocol]
}

func (api APIName) GetURI() string {
	return string(api)
}

//GetSupportedChains fetches all chains that an API supports as a map of Chain -> bool
func (api APIName) GetSupportedChains() map[Chain]bool {
	return allowedCallersByChain[api]
}

//GetSupportedProtocols fetches all Protocols that an API supports as a map of Protocol -> bool
func (api APIName) GetSupportedProtocols() map[Protocol]bool {
	return allowedCallersByProtocol[api]
}

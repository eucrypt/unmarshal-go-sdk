package constants

type APIName string

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
	ASSETS_GetAssets          APIName = "v1/:chain/address/:address/assets"
	NFT_GetAssets             APIName = "v1/:chain/address/:address/nft-assets"
	NFT_GetTxns               APIName = "v1/:chain/address/:address/nft-transactions"
	NFT_GetDetailsWithID      APIName = "v1/:chain/address/:address/details"
	NFT_GetHoldersByID        APIName = "v1/:chain/address/:address/nftholders"
	TXN_GetTokenTxns          APIName = "v1/:chain/address/:address/transactions"
	TXN_GetTxnDetails         APIName = "v1/:chain/transactions/:txnID"
	TXN_GetTokenTxnsV2        APIName = "v2/:chain/address/:address/transactions"
	PROTO_GetPositions        APIName = "v2/protocols/:protocol/address/:address/positions"
	PROTO_GetPairs            APIName = "v2/protocols/:protocol/pairs"
)

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
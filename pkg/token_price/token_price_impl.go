package token_price

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_price/types"
	"net/url"
	"strings"
)

//PriceStoreImpl is an implementation of the Unmarshal Price Store.
type PriceStoreImpl struct {
	sess session.Session
}

func New(sess session.Session) PriceStoreImpl {
	return PriceStoreImpl{sess}
}

// GetTokenPriceAtInstant GetPriceAtInstant accepts a chain, contract address and a timestamp.
//It fetches the price of a token at a given point in time.
//If the token is not found, expect an empty response with no error.
func (p PriceStoreImpl) GetTokenPrice(chain constants.Chain, contractAddress string, options *types.GetPriceOptions) (resp types.TokenPrice, err error) {
	fmt.Println("Getting token price")
	if !constants.PS_GetTokensPrice.SupportsChain(chain) {
		return types.TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetPriceWithAddress.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", contractAddress, 1)
	var urlVals = p.getTokenPriceParams(options)
	fmt.Println("Got query params")
	err = p.sess.Client.Get(&resp, path, urlVals)
	return
}

//GetTopGainers accepts only the chain and returns a list of top gainers
func (p PriceStoreImpl) GetTopGainers(chain constants.Chain, options *types.GetTopGainersOptions) (resp types.TokenDetailsResp, err error) {
	if !constants.PS_GetGainers.SupportsChain(chain) {
		return types.TokenDetailsResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetGainers.GetURI(), ":chain", chain.String(), 1)
	var urlVals = p.getGainersOptions(options)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetTopLosers accepts only the chain and returns a list of top losers for that chain
func (p PriceStoreImpl) GetTopLosers(chain constants.Chain, options *types.GetTopLosersOptions) (resp types.TokenDetailsResp, err error) {
	if !constants.PS_GetLosers.SupportsChain(chain) {
		return types.TokenDetailsResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetLosers.GetURI(), ":chain", chain.String(), 1)
	var urlVals = p.getLosersOptions(options)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetLPTokens fetches LPToken pairs and the price. It requires a chain.
//If the token is not found, expect an empty response with no error.
func (p PriceStoreImpl) GetLPTokens(chain constants.Chain, lptoken string) (resp types.TokenListWithPrice, err error) {
	if !constants.PS_GetLpTokenPrice.SupportsChain(chain) {
		return types.TokenListWithPrice{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetLpTokenPrice.GetURI(), ":chain", chain.String(), 1)
	var queryParams = map[string]interface{}{
		"lptokens": lptoken,
	}
	var urlVals = httpclient.QueryParamHelper(queryParams)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetMultipleTokenPrice takes in a chain and a list of tokens. It then returns the price for all specified tokens in the list
func (p PriceStoreImpl) GetMultipleTokenPrice(chain constants.Chain, tokenList []string) (resp types.TokenListWithPrice, err error) {
	if !constants.PS_GetTokensPrice.SupportsChain(chain) {
		return types.TokenListWithPrice{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetTokensPrice.GetURI(), ":chain", chain.String(), 1)
	var urlVals = httpclient.QueryParamHelper(map[string]interface{}{
		"tokens": tokenList,
	})
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetTokenPriceBySymbol accepts a Symbol and returns an array of token with their prices that match the symbol
//If the token is not found, expect an empty response with no error.
func (p PriceStoreImpl) GetTokenPriceBySymbol(symbol string, options *types.GetPriceWithSymbolOptions) (resp types.PriceWithSymbolResp, err error) {
	path := strings.Replace(constants.PS_GetPriceBySymbol.GetURI(), ":symbol", symbol, 1)
	var urlVals = p.getTokenPriceBySymbolParams(options)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

func (p PriceStoreImpl) getTokenPriceParams(options *types.GetPriceOptions) url.Values {
	if options == nil {
		return make(url.Values)
	}
	var queryStrings = make(map[string]interface{})
	if options.Timestamp != 0 {
		queryStrings["timestamp"] = options.Timestamp
	}
	queryStrings["24hchange"] = options.TwentyFourHourChange
	queryStrings["alternateChain"] = options.AlternateChain

	return httpclient.QueryParamHelper(queryStrings)
}

func (p PriceStoreImpl) getGainersOptions(options *types.GetTopGainersOptions) url.Values {
	if options == nil {
		return make(url.Values)
	}
	var queryStrings = make(map[string]interface{})

	queryStrings["price"] = options.MinimumPrice
	return httpclient.QueryParamHelper(queryStrings)
}

func (p PriceStoreImpl) getLosersOptions(options *types.GetTopLosersOptions) url.Values {
	if options == nil {
		return make(url.Values)
	}
	var queryStrings = make(map[string]interface{})

	queryStrings["price"] = options.MinimumPrice
	return httpclient.QueryParamHelper(queryStrings)
}

func (p PriceStoreImpl) getTokenPriceBySymbolParams(options *types.GetPriceWithSymbolOptions) url.Values {
	if options == nil {
		return make(url.Values)
	}
	var queryStrings = make(map[string]interface{})
	if options.Timestamp != 0 {
		queryStrings["timestamp"] = options.Timestamp
	}
	return httpclient.QueryParamHelper(queryStrings)
}

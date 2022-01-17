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

//GetPriceAtInstant accepts a chain, contract address and a timestamp.
//It fetches the price of a token at a given point in time.
//If the token is not found, expect an empty response with no error.
func (p PriceStoreImpl) GetTokenPriceAtInstant(chain constants.Chain, contractAddress string, timestamp int64) (resp types.TokenPrice, err error) {
	if !constants.PS_GetTokensPrice.SupportsChain(chain) {
		return types.TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetPriceWithAddress.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", contractAddress, 1)
	var queryParams = map[string]interface{}{
		"timestamp": fmt.Sprint(timestamp),
	}
	var urlVals = httpclient.QueryParamHelper(queryParams)
	err = p.sess.Client.Get(&resp, path, urlVals)
	return
}

//GetCurrentPrice accepts a chain and a contract address. It fetches the token's current price.
//If the token is not found, expect an empty response with no error.
func (p PriceStoreImpl) GetTokenCurrentPrice(chain constants.Chain, contractAddress string) (resp types.TokenPrice, err error) {
	if !constants.PS_GetTokensPrice.SupportsChain(chain) {
		return types.TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetPriceWithAddress.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", contractAddress, 1)
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetGainers accepts only the chain and returns a list of top gainers
func (p PriceStoreImpl) GetTopGainers(chain constants.Chain) (resp types.TokenDetailsResp, err error) {
	if !constants.PS_GetGainers.SupportsChain(chain) {
		return types.TokenDetailsResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetGainers.GetURI(), ":chain", chain.String(), 1)
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetLosers accepts only the chain and returns a list of top losers for that chain
func (p PriceStoreImpl) GetTopLosers(chain constants.Chain) (resp types.TokenDetailsResp, err error) {
	if !constants.PS_GetLosers.SupportsChain(chain) {
		return types.TokenDetailsResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.PS_GetLosers.GetURI(), ":chain", chain.String(), 1)
	var urlVals = make(url.Values)
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

//GetTokensPrice takes in a chain and a list of tokens. It then returns the price for all specified tokens in the list
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

//GetPriceWithSymbol accepts a Symbol and returns an array of token with their prices that match the symbol
//If the token is not found, expect an empty response with no error.
func (p PriceStoreImpl) GetTokenPriceBySymbol(symbol string) (resp types.PriceWithSymbolResp, err error) {
	path := strings.Replace(constants.PS_GetPriceBySymbol.GetURI(), ":symbol", symbol, 1)
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

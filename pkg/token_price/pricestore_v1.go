package token_price

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"net/url"
	"strings"
)

//PriceStoreV1 is the v1 implementation of the Unmarshal Price Store.
type PriceStoreV1 struct {
	sess session.Session
}

func New(sess session.Session) PriceStoreV1 {
	return PriceStoreV1{sess}
}

const PriceStoreV1Path = "v1/pricestore"

//GetPriceAtInstant accepts a chain, contract address and a timestamp.
//It fetches the price of a token at a given point in time.
//If the token is not found, expect an empty response with no error.
func (p PriceStoreV1) GetPriceAtInstant(chain constants.Chain, contractAddress string, timestamp int64) (resp TokenPrice, err error) {
	if !constants.GetTokensPrice.SupportsChain(chain) {
		return TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), contractAddress}, "/")
	var queryParams = map[string]interface{}{
		"timestamp": fmt.Sprint(timestamp),
	}
	var urlVals = httpclient.QueryParamHelper(queryParams)
	err = p.sess.Client.Get(&resp, path, urlVals)
	return
}

//GetCurrentPrice accepts a chain and a contract address. It fetches the token's current price.
//If the token is not found, expect an empty response with no error.
func (p PriceStoreV1) GetCurrentPrice(chain constants.Chain, contractAddress string) (resp TokenPrice, err error) {
	if !constants.GetTokensPrice.SupportsChain(chain) {
		return TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), contractAddress}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetGainers accepts only the chain and returns a list of top gainers
func (p PriceStoreV1) GetGainers(chain constants.Chain) (resp TokenDetailsResp, err error) {
	if !constants.GetGainers.SupportsChain(chain) {
		return TokenDetailsResp{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), "gainers"}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetLosers accepts only the chain and returns a list of top losers for that chain
func (p PriceStoreV1) GetLosers(chain constants.Chain) (resp TokenDetailsResp, err error) {
	if !constants.GetLosers.SupportsChain(chain) {
		return TokenDetailsResp{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), "losers"}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetLPTokens fetches LPToken pairs and the price. It requires a chain.
//If the token is not found, expect an empty response with no error.
func (p PriceStoreV1) GetLPTokens(chain constants.Chain, lptoken string) (resp TokenListWithPrice, err error) {
	if !constants.GetLpTokenPrice.SupportsChain(chain) {
		return TokenListWithPrice{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), "lptokens"}, "/")
	var queryParams = map[string]interface{}{
		"lptokens": lptoken,
	}
	var urlVals = httpclient.QueryParamHelper(queryParams)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetTokensPrice takes in a chain and a list of tokens. It then returns the price for all specified tokens in the list
func (p PriceStoreV1) GetTokensPrice(chain constants.Chain, tokenList []string) (resp TokenListWithPrice, err error) {
	if !constants.GetTokensPrice.SupportsChain(chain) {
		return TokenListWithPrice{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), "tokens"}, "/")
	var urlVals = httpclient.QueryParamHelper(map[string]interface{}{
		"tokens": tokenList,
	})
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetPriceWithSymbol accepts a Symbol and returns an array of token with their prices that match the symbol
//If the token is not found, expect an empty response with no error.
func (p PriceStoreV1) GetPriceWithSymbol(symbol string) (resp PriceWithSymbolResp, err error) {
	path := strings.Join([]string{PriceStoreV1Path, symbol}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

package token_price

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"net/url"
	"strings"
)

type PriceStoreV1 struct {
	sess session.Session
}

func New(sess session.Session) PriceStoreV1 {
	return PriceStoreV1{sess}
}

const PriceStoreV1Path = "v1/pricestore"

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

func (p PriceStoreV1) GetCurrentPrice(chain constants.Chain, contractAddress string) (resp TokenPrice, err error) {
	if !constants.GetTokensPrice.SupportsChain(chain) {
		return TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), contractAddress}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

func (p PriceStoreV1) GetGainers(chain constants.Chain) (resp TokenDetailsResp, err error) {
	if !constants.GetGainers.SupportsChain(chain) {
		return TokenDetailsResp{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), "gainers"}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

func (p PriceStoreV1) GetLosers(chain constants.Chain) (resp TokenDetailsResp, err error) {
	if !constants.GetLosers.SupportsChain(chain) {
		return TokenDetailsResp{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), "losers"}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

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
func (p PriceStoreV1) GetPriceWithSymbol(symbol string) (resp PriceWithSymbolResp, err error) {
	path := strings.Join([]string{PriceStoreV1Path, symbol}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

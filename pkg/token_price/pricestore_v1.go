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

func (p PriceStoreV1) GetPriceAtInstant(contractAddress string, chain constants.Chain, timestamp int64) (resp TokenPrice, err error) {
	if !constants.GetTokensPrice.IsAllowedToCallOnChain(chain) {
		return TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), contractAddress}, "/")
	var urlVals = make(url.Values)
	var queryParams = map[string]string{
		"timestamp": fmt.Sprint(timestamp),
	}
	httpclient.QueryParamHelper(queryParams, &urlVals)
	err = p.sess.Client.Get(&resp, path, urlVals)
	return
}

func (p PriceStoreV1) GetCurrentPrice(contractAddress string, chain constants.Chain) (resp TokenPrice, err error) {
	if !constants.GetTokensPrice.IsAllowedToCallOnChain(chain) {
		return TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), contractAddress}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

func (p PriceStoreV1) GetGainers(chain constants.Chain) (resp TokenPriceList, err error) {
	if !constants.GetGainers.IsAllowedToCallOnChain(chain) {
		return TokenPriceList{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), "gainers"}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}
func (p PriceStoreV1) GetLosers(chain constants.Chain) (resp TokenPriceList, err error) {
	if !constants.GetLosers.IsAllowedToCallOnChain(chain) {
		return TokenPriceList{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), "losers"}, "/")
	var urlVals = make(url.Values)
	err = p.sess.Client.Get(&resp, path, urlVals)

	return
}

package token_price

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"net/url"
	"strings"
)

type PriceStore interface {
	GetPriceAtInstant(contractAddress string, timestamp int64) (TokenPrice, error)
}

type PriceStoreV1 struct {
	sess pkg.Session
}

func New(sess pkg.Session) PriceStoreV1 {
	return PriceStoreV1{sess}
}

const PriceStoreV1Path = "/pricestore"

func (t PriceStoreV1) GetPriceAtInstant(contractAddress string, chain constants.Chain, timestamp int64) (resp TokenPrice, err error) {
	if constants.GetTokensPrice.IsAllowedToCallOnChain(chain) {
		return TokenPrice{}, constants.UnsupportedChainError
	}
	path := strings.Join([]string{PriceStoreV1Path, "chain", chain.String(), contractAddress}, "/")
	var urlVals = new(url.Values)
	var queryParams = map[string]string{
		"timestamp": fmt.Sprint(timestamp),
	}
	httpclient.QueryParamHelper(queryParams, urlVals)
	err = t.sess.Client.Get(&resp, path, *urlVals)
	return
}

func (t PriceStoreV1) GetCurrentPrice(contractAddress string) (resp TokenPrice, err error) {
	path := strings.Join([]string{PriceStoreV1Path, "all"}, "/")
	var urlVals = new(url.Values)
	err = t.sess.Client.Get(&resp, path, *urlVals)
	return
}

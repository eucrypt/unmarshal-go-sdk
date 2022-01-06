package token_details

import (
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"net/url"
	"strings"
)

const TokenStoreV1Path = "v1/tokenstore/token"

type TokenStoreV1 struct {
	sess session.Session
}

func New(sess session.Session) TokenStoreV1 {
	return TokenStoreV1{sess}
}

func (t TokenStoreV1) GetTokenDetailsWithContract(contractAddress string) (resp TokenDetails, err error) {
	path := strings.Join([]string{TokenStoreV1Path, "address", contractAddress}, "/")
	err = t.sess.Client.Get(&resp, path, nil)
	return
}

func (t TokenStoreV1) GetTokenList(queryParams map[string]string) (resp GetTokenListResponse, err error) {
	path := strings.Join([]string{TokenStoreV1Path, "all"}, "/")
	var urlVals = new(url.Values)
	httpclient.QueryParamHelper(queryParams, urlVals)
	err = t.sess.Client.Get(&resp, path, *urlVals)
	return
}

func (t TokenStoreV1) GetTokenWithSymbol(symbol string) (resp []TokenDetails, err error) {
	path := strings.Join([]string{TokenStoreV1Path, "symbol", symbol}, "/")
	err = t.sess.Client.Get(&resp, path, nil)
	return
}

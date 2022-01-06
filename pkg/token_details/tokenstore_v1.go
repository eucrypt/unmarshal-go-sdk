package token_details

import (
	"fmt"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
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

func (t TokenStoreV1) GetTokenList(pageNumber int, pageSize int) (resp GetTokenListResponse, err error) {
	path := strings.Join([]string{TokenStoreV1Path, "all"}, "/")
	vals := map[string]interface{}{
		"pageNumber": fmt.Sprint(pageNumber),
		"pageSize":   fmt.Sprint(pageSize),
	}
	var urlVals = httpclient.QueryParamHelper(vals)
	err = t.sess.Client.Get(&resp, path, urlVals)
	return
}

func (t TokenStoreV1) GetTokenWithSymbol(symbol string) (resp []TokenDetails, err error) {
	path := strings.Join([]string{TokenStoreV1Path, "symbol", symbol}, "/")
	err = t.sess.Client.Get(&resp, path, nil)
	return
}

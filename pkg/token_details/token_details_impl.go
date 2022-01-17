package token_details

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_details/types"
	"strings"
)

type TokenStoreV1 struct {
	sess session.Session
}

func New(sess session.Session) TokenStoreV1 {
	return TokenStoreV1{sess}
}

//GetTokenDetailsByContract returns token data when provided with a valid contract.
//The search happens across every supported chain
func (t TokenStoreV1) GetTokenDetailsByContract(contractAddress string) (resp types.TokenDetails, err error) {
	path := strings.Replace(constants.TS_GetDetailsWithContract.GetURI(), ":address", contractAddress, 1)
	err = t.sess.Client.Get(&resp, path, nil)
	return
}

//GetTokenList Returns a list of tracked tokens from the token store.
//It accepts a page size and page number. If either is 0 the default value is to be assumed on the API end
func (t TokenStoreV1) GetTokenList(pageNumber int, pageSize int) (resp types.GetTokenListResponse, err error) {
	path := constants.TS_GetTokenList.GetURI()
	vals := map[string]interface{}{
		"pageNumber": fmt.Sprint(pageNumber),
		"pageSize":   fmt.Sprint(pageSize),
	}
	var urlVals = httpclient.QueryParamHelper(vals)
	err = t.sess.Client.Get(&resp, path, urlVals)
	return
}

//GetTokenDetailsBySymbol Accepts a symbol and returns token data.
//The search is cross-chain and the result includes the blockchain of the specific token
func (t TokenStoreV1) GetTokenDetailsBySymbol(symbol string) (resp []types.TokenDetails, err error) {
	path := strings.Replace(constants.TS_GetTokenWithSymbol.GetURI(), ":symbol", symbol, 1)
	err = t.sess.Client.Get(&resp, path, nil)
	return
}

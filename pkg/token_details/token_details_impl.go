package token_details

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_details/types"
	"net/url"
	"strings"
)

type TokenStoreV1 struct {
	sess session.Session
}

func New(sess session.Session) TokenStoreV1 {
	return TokenStoreV1{sess}
}

//GetTokenDetailsByContract returns token data when provided with a valid contract.
//The search happens across every supported chain by default, additionally accepting the chain param for a more
//specific search
func (t TokenStoreV1) GetTokenDetailsByContract(contractAddress string, options *TokenDetailsOptions) (
	resp types.TokenDetails, err error) {
	var urlVals url.Values
	if options != nil {
		urlVals = httpclient.QueryParamHelper(map[string]interface{}{
			"chain": options.Chain.String(),
		})
	}
	path := strings.Replace(constants.TS_GetDetailsByContract.GetURI(), ":address", contractAddress, 1)
	err = t.sess.Client.Get(&resp, path, urlVals)
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
	path := strings.Replace(constants.TS_GetTokenBySymbol.GetURI(), ":symbol", symbol, 1)
	err = t.sess.Client.Get(&resp, path, nil)
	return
}

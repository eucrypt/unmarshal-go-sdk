package assets

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/assets/types"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"strings"
)

type V1Store struct {
	sess session.Session
}

//GetProfitAndLoss Accepts chain, address and the token contract. It uses this data to return user data that gives the
//user the profit or loss incurred by the address holder for the particular contract
func (a V1Store) GetProfitAndLoss(chain constants.Chain, address, contract string) (response types.UserContractData, err error) {
	if !constants.ASSETS_GetProfitsAndLoss.SupportsChain(chain) {
		return response, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.ASSETS_GetProfitsAndLoss.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	vals := map[string]interface{}{
		"contract": fmt.Sprint(contract),
	}
	queryParams := httpclient.QueryParamHelper(vals)
	err = a.sess.Client.Get(&response, path, queryParams)
	return
}

func New(sess session.Session) V1Store {
	return V1Store{sess: sess}
}

//GetTokenAssets accepts the chain and address and returns the assets of the address on the chain. It includes,
//in addition to the native token balances, all ERC20 assets (EVM chains) or SPL tokens (Solana)
func (a V1Store) GetTokenAssets(chain constants.Chain, address string) (response types.AssetDetailsV1Resp, err error) {
	if !constants.ASSETS_GetTokenAssets.SupportsChain(chain) {
		return response, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.ASSETS_GetTokenAssets.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	err = a.sess.Client.Get(&response, path, nil)
	return
}

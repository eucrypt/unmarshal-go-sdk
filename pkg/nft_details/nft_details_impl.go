package nft_details

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/nft_details/types"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"strings"
)

type NFTDetailsImpl struct {
	sess session.Session
}

func New(sess session.Session) NFTDetailsImpl {
	return NFTDetailsImpl{sess: sess}
}

func (nft NFTDetailsImpl) GetAssetsByAddress(chain constants.Chain, address string) (resp types.NFTAssetsResp, err error) {
	if !constants.NFT_GetAssets.SupportsChain(chain) {
		return types.NFTAssetsResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.NFT_GetAssets.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	err = nft.sess.Client.Get(&resp, path, nil)

	return
}

func (nft NFTDetailsImpl) GetTransactionsByAddress(chain constants.Chain, address string, pageNumber int, pageSize int) (
	resp types.NFTTxnsResp, err error) {
	if !constants.NFT_GetTxns.SupportsChain(chain) {
		return types.NFTTxnsResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.NFT_GetTxns.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	vals := map[string]interface{}{
		"page":     fmt.Sprint(pageNumber),
		"pageSize": fmt.Sprint(pageSize),
	}
	var urlVals = httpclient.QueryParamHelper(vals)
	err = nft.sess.Client.Get(&resp, path, urlVals)

	return
}

func (nft NFTDetailsImpl) GetDetailsByID(chain constants.Chain, tokenID string, NFTAddress string) (
	resp types.NFTByTokenIDResp, err error) {

	if !constants.NFT_GetDetailsWithID.SupportsChain(chain) {
		return types.NFTByTokenIDResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.NFT_GetDetailsWithID.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", NFTAddress, 1)
	vals := map[string]interface{}{
		"tokenId": tokenID,
	}
	var urlVals = httpclient.QueryParamHelper(vals)
	err = nft.sess.Client.Get(&resp, path, urlVals)
	return
}

func (nft NFTDetailsImpl) GetHolderByID(chain constants.Chain, tokenID string, NFTAddress string) (
	resp types.NFTHolderResponse, err error) {

	if !constants.NFT_GetHoldersByID.SupportsChain(chain) {
		return types.NFTHolderResponse{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.NFT_GetHoldersByID.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", NFTAddress, 1)
	vals := map[string]interface{}{
		"tokenId": tokenID,
	}
	var urlVals = httpclient.QueryParamHelper(vals)
	err = nft.sess.Client.Get(&resp, path, urlVals)

	return
}

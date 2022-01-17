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

//GetNFTAssetsByAddress accepts a chain and address and returns the NFT assets in possession with the address.
//For Solana it works only with the Metaplex NFTs wherein it returns assets whose metadata has a URI
func (nft NFTDetailsImpl) GetNFTAssetsByAddress(chain constants.Chain, address string) (resp types.NFTAssetsResp, err error) {
	if !constants.NFT_GetNFTAssets.SupportsChain(chain) {
		return types.NFTAssetsResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.NFT_GetNFTAssets.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	err = nft.sess.Client.Get(&resp, path, nil)

	return
}

//GetNFTTransactionsByAddress accepts a chain, address and pagination params.
//It returns the list of the address' NFT transactions
func (nft NFTDetailsImpl) GetNFTTransactionsByAddress(chain constants.Chain, address string, pageNumber int, pageSize int) (
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

//GetNFTDetailsByID accepts the NFT's tokenID, chain and address and returns the associated NFT metadata
func (nft NFTDetailsImpl) GetNFTDetailsByID(chain constants.Chain, tokenID string, NFTAddress string) (
	resp types.NFTByTokenIDResp, err error) {

	if !constants.NFT_GetDetailsByID.SupportsChain(chain) {
		return types.NFTByTokenIDResp{}, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.NFT_GetDetailsByID.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", NFTAddress, 1)
	vals := map[string]interface{}{
		"tokenId": tokenID,
	}
	var urlVals = httpclient.QueryParamHelper(vals)
	err = nft.sess.Client.Get(&resp, path, urlVals)
	return
}

func (nft NFTDetailsImpl) GetNFTHolderByID(chain constants.Chain, tokenID string, NFTAddress string) (
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

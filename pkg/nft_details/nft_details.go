package nft_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/nft_details/types"
)

type NftDetails interface {
	GetNFTAssetsByAddress(chain constants.Chain, address string) (response types.NFTAssetsResp, err error)
	GetNFTTransactionsByAddress(chain constants.Chain, address string, pageNumber int, pageSize int) (response types.NFTTxnsResp, err error)
	GetNFTDetailsByID(chain constants.Chain, tokenID string, NFTAddress string) (response types.NFTByTokenIDResp, err error)
	GetNFTHolderByID(chain constants.Chain, tokenID string, NFTAddress string) (response types.NFTHolderResponse, err error)
}

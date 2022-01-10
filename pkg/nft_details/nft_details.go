package nft_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/nft_details/types"
)

type NftDetails interface {
	GetAssetsByAddress(chain constants.Chain, address string) (response types.NFTAssetsResp, err error)
	GetTransactionsByAddress(chain constants.Chain, address string, pageNumber int, pageSize int) (response types.NFTTxnsResp, err error)
	GetDetailsByID(chain constants.Chain, tokenID string, NFTAddress string) (response types.NFTByTokenIDResp, err error)
	GetHolderByID(chain constants.Chain, tokenID string, NFTAddress string) (response types.NFTHolderResponse, err error)
}

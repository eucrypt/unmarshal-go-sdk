package assets

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/assets/types"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
)

type Assets interface {
	GetTokenAssets(chain constants.Chain, address string) (response types.AssetDetailsV1Resp, err error)
	GetProfitAndLoss(chain constants.Chain, address, contract string) (response types.UserContractData, err error)
}

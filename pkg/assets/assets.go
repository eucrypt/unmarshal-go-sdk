package assets

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/assets/types"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
)

type Assets interface {
	GetTokenAssets(chain constants.Chain, address string) (response types.AssetDetailsV1Resp, err error)
}

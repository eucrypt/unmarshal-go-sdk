package assets

import "github.com/eucrypt/unmarshal-go-sdk/pkg/constants"

type Assets interface {
	GetAssets(chain constants.Chain, address string) (response AssetDetailsV1Resp, err error)
}

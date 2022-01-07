package assets

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/assets/types"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"strings"
)

type V1Store struct {
	sess session.Session
}

func New(sess session.Session) V1Store {
	return V1Store{sess: sess}
}

func (a V1Store) GetAssets(chain constants.Chain, address string) (response types.AssetDetailsV1Resp, err error) {
	if !constants.Assets_GetAssets.SupportsChain(chain) {
		return response, constants.UnsupportedChainError
	}
	path := strings.Replace(constants.Assets_GetAssets.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	err = a.sess.Client.Get(&response, path, nil)
	return
}

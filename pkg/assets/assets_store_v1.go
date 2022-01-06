package assets

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"strings"
)

type V1Store struct {
	sess session.Session
}

const V1Path = "v1"

func New(sess session.Session) V1Store {
	return V1Store{sess: sess}
}

func (a V1Store) GetAssets(chain constants.Chain, address string) (response AssetDetailsV1Resp, err error) {
	if !constants.GetAssets.SupportsChain(chain) {
		return response, constants.UnsupportedChainError
	}

	path := strings.Join([]string{V1Path, chain.String(), "address", address, "assets"}, "/")
	err = a.sess.Client.Get(&response, path, nil)
	return
}

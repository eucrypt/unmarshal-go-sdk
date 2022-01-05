package token_details

import "github.com/eucrypt/unmarshal-go-sdk/pkg"

type TokenStore interface {
	GetTokenDetails(contractAddress string)
}

type TokenStoreImpl struct {
	sess pkg.Session
}

func New(sess pkg.Session) TokenStoreImpl {
	return TokenStoreImpl{sess}
}

func (t TokenStoreImpl) GetTokenDetails(contractAddress string) {
	panic("implement me")
}

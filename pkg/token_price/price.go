package token_price

import "github.com/eucrypt/unmarshal-go-sdk/pkg"

type PriceStore interface {
	GetPriceAtInstant(contractAddress string, timestamp int64)
}

type PriceStoreImpl struct {
	sess pkg.Session
}

func New(sess pkg.Session) PriceStoreImpl {
	return PriceStoreImpl{sess}
}

func (t PriceStoreImpl) GetPriceAtInstant(contractAddress string, timestamp int64) {
	panic("implement me")
}

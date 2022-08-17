package token_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_details/types"
)

type TokenDetails interface {
	GetTokenDetailsByContract(contractAddress string, options *TokenDetailsOptions) (types.TokenDetails, error)
	GetTokenList(pageNumber int, pageSize int) (types.GetTokenListResponse, error)
	GetTokenDetailsBySymbol(string) ([]types.TokenDetails, error)
}

type TokenDetailsOptions struct {
	Chain constants.Chain
}

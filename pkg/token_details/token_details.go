package token_details

import "github.com/eucrypt/unmarshal-go-sdk/pkg/token_details/types"

type TokenDetails interface {
	GetTokenDetailsByContract(contractAddress string) (types.TokenDetails, error)
	GetTokenList(pageNumber int, pageSize int) (types.GetTokenListResponse, error)
	GetTokenDetailsBySymbol(string) ([]types.TokenDetails, error)
}

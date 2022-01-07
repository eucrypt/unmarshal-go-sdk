package token_details

import "github.com/eucrypt/unmarshal-go-sdk/pkg/token_details/types"

type TokenDetails interface {
	GetDetailsWithContract(contractAddress string) (types.TokenDetails, error)
	GetTokenList(pageNumber int, pageSize int) (types.GetTokenListResponse, error)
	GetTokenWithSymbol(string) ([]types.TokenDetails, error)
}

package token_details

type TokenStore interface {
	GetTokenDetailsWithContract(contractAddress string) (TokenDetails, error)
	GetTokenList(pageNumber int, pageSize int) (GetTokenListResponse, error)
	GetTokenWithSymbol(string) ([]TokenDetails, error)
}

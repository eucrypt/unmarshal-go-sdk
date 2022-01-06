package token_details

type TokenStore interface {
	GetTokenDetailsWithContract(contractAddress string) (TokenDetails, error)
	GetTokenList(int, int) (GetTokenListResponse, error)
	GetTokenWithSymbol(string) ([]TokenDetails, error)
}

type TokenDetails struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Contract    string `json:"contract"`
	Image       string `json:"image"`
	Decimal     int    `json:"decimal"`
	Blockchain  string `json:"blockchain"`
	TotalSupply string `json:"total_supply"`
	Verified    bool   `json:"verified"`
	Website     string `json:"website,omitempty"`
	Explorer    string `json:"explorer,omitempty"`
}
type GetTokenListResponse struct {
	Page        int            `json:"page"`
	TotalPages  int            `json:"total_pages"`
	ItemsOnPage int            `json:"items_on_page"`
	Data        []TokenDetails `json:"data"`
}

package types

type TokenPrice struct {
	TokenId   string `json:"tokenId"`
	Timestamp string `json:"timestamp"`
	Price     string `json:"price"`
}
type TokenDetailsWithPrice struct {
	Name          string `json:"name"`
	Contract      string `json:"contract"`
	Decimal       int    `json:"decimal"`
	Logo          string `json:"logo"`
	Symbol        string `json:"symbol"`
	CurrentPrice  string `json:"current_price"`
	PercentChange string `json:"percent_change"`
	DayLow        string `json:"day_low"`
	DayHigh       string `json:"day_high"`
}

type TokenDetailsResp []TokenDetailsWithPrice

type MultiTokenResponse struct {
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	TokenId          string `json:"tokenId"`
	Price            string `json:"price"`
	PercentageChange string `json:"percentage_change"`
}

type TokenListWithPrice []MultiTokenResponse

type PriceFromSymbol struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Contract    string `json:"contract"`
	Decimal     int    `json:"decimal"`
	Blockchain  string `json:"blockchain"`
	Price       string `json:"price"`
	PriceChange string `json:"price_change"`
	Timestamp   string `json:"timestamp"`
}

type PriceWithSymbolResp []PriceFromSymbol

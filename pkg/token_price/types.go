package token_price

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

type TokenPriceList []TokenDetailsWithPrice

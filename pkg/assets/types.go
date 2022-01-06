package assets

type AssetDetailsV1 struct {
	ContractName         string  `json:"contract_name"`
	ContractTickerSymbol string  `json:"contract_ticker_symbol"`
	ContractDecimals     int     `json:"contract_decimals"`
	ContractAddress      string  `json:"contract_address"`
	Coin                 int     `json:"coin"`
	Type                 string  `json:"type"`
	Balance              string  `json:"balance"`
	Quote                float64 `json:"quote"`
	QuoteRate            float64 `json:"quote_rate"`
	LogoUrl              string  `json:"logo_url"`
	QuoteRate24H         string  `json:"quote_rate_24h"`
	QuotePctChange24H    float64 `json:"quote_pct_change_24h"`
}

type AssetDetailsV1Resp []AssetDetailsV1

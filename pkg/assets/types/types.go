package types

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

type UserContractData struct {
	QuoteRate              float64 `json:"quote_rate"`
	TotalFeesPaid          float64 `json:"total_fees_paid"`
	TotalFeesPaidUsd       float64 `json:"total_fees_paid_usd"`
	AverageTokenPrice      float64 `json:"average_token_price"`
	OverallProfitLoss      float64 `json:"overall_profit_loss"`
	CurrentHoldingQuantity float64 `json:"current_holding_quantity"`
	PercentageChange24H    float64 `json:"percentage_change_24H"`
	PriceChange24H         float64 `json:"price_change_24H"`
}

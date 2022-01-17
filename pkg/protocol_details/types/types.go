package types

type GetPairsResp struct {
	Total         int `json:"total"`
	ProtocolPairs []struct {
		Token0 struct {
			ContractName         string  `json:"contract_name"`
			ContractTickerSymbol string  `json:"contract_ticker_symbol"`
			ContractDecimals     int     `json:"contract_decimals"`
			ContractAddress      string  `json:"contract_address"`
			Coin                 int     `json:"coin"`
			Quote                int     `json:"quote"`
			QuoteRate            float64 `json:"quote_rate"`
			LogoUrl              string  `json:"logo_url"`
			QuoteRate24H         string  `json:"quote_rate_24h"`
			QuotePctChange24H    int     `json:"quote_pct_change_24h"`
		} `json:"token_0"`
		Token1 struct {
			ContractName         string  `json:"contract_name"`
			ContractTickerSymbol string  `json:"contract_ticker_symbol"`
			ContractDecimals     int     `json:"contract_decimals"`
			ContractAddress      string  `json:"contract_address"`
			Coin                 int     `json:"coin"`
			Quote                int     `json:"quote"`
			QuoteRate            float64 `json:"quote_rate"`
			LogoUrl              string  `json:"logo_url"`
			QuoteRate24H         string  `json:"quote_rate_24h"`
			QuotePctChange24H    int     `json:"quote_pct_change_24h"`
		} `json:"token_1"`
		Pair struct {
			ContractName         string  `json:"contract_name"`
			ContractTickerSymbol string  `json:"contract_ticker_symbol"`
			ContractDecimals     int     `json:"contract_decimals"`
			ContractAddress      string  `json:"contract_address"`
			Coin                 int     `json:"coin"`
			Quote                int     `json:"quote"`
			QuoteRate            float64 `json:"quote_rate"`
			LogoUrl              string  `json:"logo_url"`
			QuoteRate24H         string  `json:"quote_rate_24h"`
			QuotePctChange24H    int     `json:"quote_pct_change_24h"`
		} `json:"pair"`
		TotalLiquidityQuote  string `json:"total_liquidity_quote"`
		Token0Reserve        string `json:"token_0_reserve"`
		Token1Reserve        string `json:"token_1_reserve"`
		CreatedAtBlockNumber string `json:"created_at_block_number"`
	} `json:"protocol_pairs"`
}

type GetPositionsResp struct {
	UserAddress string `json:"user_address"`
	Positions   []struct {
		Token0Reserve string `json:"token_0_reserve"`
		Token1Reserve string `json:"token_1_reserve"`
		Token0        struct {
			ContractName         string `json:"contract_name"`
			ContractTickerSymbol string `json:"contract_ticker_symbol"`
			ContractDecimals     int    `json:"contract_decimals"`
			ContractAddress      string `json:"contract_address"`
			Coin                 int    `json:"coin"`
			Balance              string `json:"balance"`
			Quote                int    `json:"quote"`
			QuoteRate            int    `json:"quote_rate"`
			LogoUrl              string `json:"logo_url"`
			QuoteRate24H         string `json:"quote_rate_24h"`
			QuotePctChange24H    int    `json:"quote_pct_change_24h"`
		} `json:"token_0"`
		Token1 struct {
			ContractName         string `json:"contract_name"`
			ContractTickerSymbol string `json:"contract_ticker_symbol"`
			ContractDecimals     int    `json:"contract_decimals"`
			ContractAddress      string `json:"contract_address"`
			Coin                 int    `json:"coin"`
			Balance              string `json:"balance"`
			Quote                int    `json:"quote"`
			QuoteRate            int    `json:"quote_rate"`
			LogoUrl              string `json:"logo_url"`
			QuoteRate24H         string `json:"quote_rate_24h"`
			QuotePctChange24H    int    `json:"quote_pct_change_24h"`
		} `json:"token_1"`
		PoolToken struct {
			ContractName         string  `json:"contract_name"`
			ContractTickerSymbol string  `json:"contract_ticker_symbol"`
			ContractDecimals     int     `json:"contract_decimals"`
			ContractAddress      string  `json:"contract_address"`
			Coin                 int     `json:"coin"`
			Balance              string  `json:"balance"`
			Quote                float64 `json:"quote"`
			QuoteRate            float64 `json:"quote_rate"`
			LogoUrl              string  `json:"logo_url"`
			QuoteRate24H         string  `json:"quote_rate_24h"`
			QuotePctChange24H    int     `json:"quote_pct_change_24h"`
			TotalSupply          string  `json:"total_supply"`
			PoolSharePercentage  string  `json:"pool_share_percentage"`
			LiquidityValue       float64 `json:"liquidity_value"`
		} `json:"pool_token"`
	} `json:"positions"`
}

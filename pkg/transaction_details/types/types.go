package types

type TokenTxn struct {
	Page         int `json:"page"`
	TotalPages   int `json:"total_pages"`
	ItemsOnPage  int `json:"items_on_page"`
	TotalTxs     int `json:"total_txs"`
	Transactions []struct {
		Id                  string `json:"id"`
		From                string `json:"from"`
		To                  string `json:"to"`
		Fee                 string `json:"fee"`
		Date                int    `json:"date"`
		Status              string `json:"status"`
		Type                string `json:"type"`
		Block               int    `json:"block"`
		Value               string `json:"value"`
		Nonce               int    `json:"nonce"`
		NativeTokenDecimals int    `json:"native_token_decimals"`
		Description         string `json:"description"`
		Received            []struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			TokenId  string `json:"token_id"`
			Decimals int    `json:"decimals"`
			Value    string `json:"value"`
			LogoUrl  string `json:"logo_url"`
			From     string `json:"from"`
			To       string `json:"to"`
		} `json:"received,omitempty"`
		Sent []struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			TokenId  string `json:"token_id"`
			Decimals int    `json:"decimals"`
			Value    string `json:"value"`
			LogoUrl  string `json:"logo_url"`
			From     string `json:"from"`
			To       string `json:"to"`
		} `json:"sent,omitempty"`
	} `json:"transactions"`
}

type TxnByID struct {
	Id                  string `json:"id"`
	From                string `json:"from"`
	To                  string `json:"to"`
	Fee                 string `json:"fee"`
	Date                int    `json:"date"`
	Status              string `json:"status"`
	Type                string `json:"type"`
	Block               int    `json:"block"`
	Value               string `json:"value"`
	Nonce               int    `json:"nonce"`
	NativeTokenDecimals int    `json:"native_token_decimals"`
	Description         string `json:"description"`
	Sent                []struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		TokenId  string `json:"token_id"`
		Decimals int    `json:"decimals"`
		Value    string `json:"value"`
		LogoUrl  string `json:"logo_url"`
		From     string `json:"from"`
		To       string `json:"to"`
	} `json:"sent"`
}

type TokenTxnV2 struct {
	Page         int `json:"page"`
	TotalPages   int `json:"total_pages"`
	ItemsOnPage  int `json:"items_on_page"`
	TotalTxs     int `json:"total_txs"`
	Transactions []struct {
		Id                  string `json:"id"`
		From                string `json:"from"`
		To                  string `json:"to"`
		Fee                 string `json:"fee"`
		Date                int    `json:"date"`
		Status              string `json:"status"`
		Type                string `json:"type"`
		Block               int    `json:"block"`
		Value               string `json:"value"`
		Nonce               int    `json:"nonce"`
		NativeTokenDecimals int    `json:"native_token_decimals"`
		Description         string `json:"description"`
		Sent                []struct {
			Name      string  `json:"name"`
			Symbol    string  `json:"symbol"`
			TokenId   string  `json:"token_id"`
			Decimals  int     `json:"decimals"`
			Value     string  `json:"value"`
			LogoUrl   string  `json:"logo_url"`
			From      string  `json:"from"`
			To        string  `json:"to"`
			QuoteRate float64 `json:"quoteRate,omitempty"`
			Quote     float64 `json:"quote,omitempty"`
		} `json:"sent,omitempty"`
		Received []struct {
			Name      string  `json:"name"`
			Symbol    string  `json:"symbol"`
			TokenId   string  `json:"token_id"`
			Decimals  int     `json:"decimals"`
			Value     string  `json:"value"`
			Quote     float64 `json:"quote,omitempty"`
			QuoteRate float64 `json:"quoteRate,omitempty"`
			LogoUrl   string  `json:"logo_url"`
			From      string  `json:"from"`
			To        string  `json:"to"`
		} `json:"received,omitempty"`
	} `json:"transactions"`
}

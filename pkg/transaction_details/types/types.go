package types

import "math/big"

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
type RawTransactionsResponseV1 struct {
	TotalCount        int              `json:"total_count"`
	NextPage          bool             `json:"next_page"`
	LastVerifiedBlock *big.Int         `json:"last_verified_block"`
	Transactions      []RawTransaction `json:"result"`
}

type RawTokenTransfer struct {
	From         string   `json:"from"`
	To           string   `json:"to"`
	Token        string   `json:"token"`
	Value        *big.Int `json:"value"`
	TokenDecimal int64    `json:"token_decimal,omitempty"`
	TokenName    string   `json:"token_name,omitempty"`
	TokenSymbol  string   `json:"token_symbol,omitempty"`
	TokenImage   string   `json:"token_image,omitempty"`
	LogIndex     string   `json:"log_index"`
}

type RawAdditionalData struct {
	Status   int64    `json:"status"`
	Nonce    *big.Int `json:"nonce"`
	GasLimit *big.Int `json:"gas_limit"`
	GasUsed  *big.Int `json:"gas_used"`
	GasPrice *big.Int `json:"gas_price"`
	Data     string   `json:"data"`
}

type RawTransaction struct {
	TxHash         string             `json:"tx_hash"`
	From           string             `json:"from"`
	To             string             `json:"to"`
	Value          *big.Int           `json:"value"`
	BlockHash      string             `json:"block_hash"`
	BlockNumber    string             `json:"block_number"`
	BlockTime      *big.Int           `json:"block_time"`
	Fees           *big.Int           `json:"fees"`
	AdditionalData RawAdditionalData  `json:"additional_data"`
	TokenTransfers []RawTokenTransfer `json:"token_transfers"`
	TxIndex        uint               `json:"tx_index"`
	ShardId        int                `json:"shard_id"`
	ToShardId      int                `json:"to_shard_id"`
}

type TransactionByCursorRequest struct {
	StartCursor *AddressTxCursor `json:"start_cursor,omitempty"`
	EndCursor   *AddressTxCursor `json:"end_cursor,omitempty"`
}

type AddressTxCursor struct {
	ContractAddress string   `json:"contract_address"`
	BlockNumber     *big.Int `json:"block_number" `
	SeqID           uint64   `json:"seq_id"`
}

type GetTransactionByCursorResponse struct {
	EndCursor         *AddressTxCursor `json:"end_cursor,omitempty"`
	StartCursor       *AddressTxCursor `json:"start_cursor,omitempty"`
	LastVerifiedBlock *big.Int         `json:"last_verified_block"`
	Transactions      []RawTransaction `json:"transactions"`
}

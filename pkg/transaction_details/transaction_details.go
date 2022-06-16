package transaction_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
)

type TokenTxnsOpts struct {
	Contract string
	PaginationOptions
}
type PaginationOptions struct {
	Page     int
	PageSize int
}

type RawTransactionOptions struct {
	PaginationOptions
	FromBlock uint64
	ToBlock   uint64
}

type TransactionDetails interface {
	GetTokenTxns(chain constants.Chain, address string, options *TokenTxnsOpts) (resp types.TokenTxn, err error)
	GetTxnDetails(chain constants.Chain, txnID string) (resp types.TxnByID, err error)
	GetTokenTxnsV2(chain constants.Chain, address string, options *TokenTxnsOpts) (resp types.TokenTxnV2, err error)
	GetRawTransactionsForAddress(chain constants.Chain, address string, options *RawTransactionOptions) (
		resp types.RawTransactionsResponseV1, err error)
}

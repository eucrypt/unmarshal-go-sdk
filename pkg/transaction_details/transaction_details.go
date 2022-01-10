package transaction_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
)

type TokenTxnsOpts struct {
	Contract string
	Page     int
	PageSize int
}

type TransactionDetails interface {
	GetTokenTxns(chain constants.Chain, address string, options *TokenTxnsOpts) (resp types.TokenTxn, err error)
	GetTxnDetails(chain constants.Chain, txnID string) (resp types.TxnByID, err error)
	GetTokenTxnsV2(chain constants.Chain, address string, options *TokenTxnsOpts) (resp types.TokenTxnV2, err error)
}

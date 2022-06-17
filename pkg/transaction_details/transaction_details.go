package transaction_details

import (
	"fmt"
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

type TransactionDetailsOpts struct {
	PaginationOptions
	FromBlock uint64
	ToBlock   uint64
}

const RawFormat = "raw"

func (p PaginationOptions) getPaginationAsQueryParams() map[string]interface{} {
	var queryParams = map[string]interface{}{
		"page":     fmt.Sprint(p.Page),
		"pageSize": fmt.Sprint(p.PageSize),
	}
	return queryParams
}

func (options TokenTxnsOpts) getMappableQueryParams() map[string]interface{} {
	var queryParams = options.getPaginationAsQueryParams()
	queryParams["contract"] = options.Contract
	return queryParams
}

func (options TransactionDetailsOpts) getMappableQueryParamsForRawFormat() map[string]interface{} {
	var queryParams = options.getPaginationAsQueryParams()
	queryParams["fromBlock"] = options.FromBlock
	queryParams["toBlock"] = options.ToBlock
	queryParams["format"] = RawFormat
	return queryParams
}

type TransactionDetails interface {
	GetTokenTxns(chain constants.Chain, address string, options *TokenTxnsOpts) (resp types.TokenTxn, err error)
	GetTxnDetails(chain constants.Chain, txnID string) (resp types.TxnByID, err error)
	GetTokenTxnsV2(chain constants.Chain, address string, options *TokenTxnsOpts) (resp types.TokenTxnV2, err error)
	GetRawTransactionsForAddress(chain constants.Chain, address string, options *TransactionDetailsOpts) (
		resp types.RawTransactionsResponseV1, err error)
}

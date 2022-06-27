package transaction_details

import (
	"fmt"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
)

type TokenTxnsOpts struct {
	Contract string
	PaginationOptions
	BlockLimitsOpts
}

type PaginationOptions struct {
	Page     int
	PageSize int
}

type BlockLimitsOpts struct {
	FromBlock uint64
	ToBlock   uint64
}

type TransactionDetailsOpts struct {
	PaginationOptions
	BlockLimitsOpts
	format transactionFormats
}

type transactionFormats string

const (
	Standard transactionFormats = "standard"
	Raw      transactionFormats = "raw"
)

func (options TokenTxnsOpts) getMappableQueryParams() map[string]interface{} {
	queryParams := make(map[string]interface{})
	options.mustAddPaginationToQueryParams(queryParams)
	options.mustAddBlockLimitsToQueryParams(queryParams)
	queryParams["contract"] = options.Contract
	return queryParams
}

func (options TransactionDetailsOpts) getMappableQueryParams() map[string]interface{} {
	queryParams := make(map[string]interface{})
	options.mustAddPaginationToQueryParams(queryParams)
	options.mustAddBlockLimitsToQueryParams(queryParams)
	queryParams["format"] = options.format
	return queryParams
}

func (p PaginationOptions) mustAddPaginationToQueryParams(queryParams map[string]interface{}) {
	queryParams["page"] = fmt.Sprint(p.Page)
	queryParams["pageSize"] = fmt.Sprint(p.PageSize)
	return
}

func (blk BlockLimitsOpts) mustAddBlockLimitsToQueryParams(queryParams map[string]interface{}) {
	queryParams["fromBlock"] = blk.FromBlock
	queryParams["toBlock"] = blk.ToBlock
}

type TransactionDetails interface {
	GetTokenTxns(chain constants.Chain, address string, options *TokenTxnsOpts) (resp types.TokenTxn, err error)
	GetTxnDetails(chain constants.Chain, txnID string) (resp types.TxnByID, err error)
	GetTokenTxnsV2(chain constants.Chain, address string, options *TokenTxnsOpts) (resp types.TokenTxnV2, err error)
	GetRawTransactionsForAddress(chain constants.Chain, address string, options *TransactionDetailsOpts) (
		resp types.RawTransactionsResponseV1, err error)
}

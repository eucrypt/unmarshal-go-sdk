package transaction_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
	"net/url"
	"strings"
)

type TxnDetailsImpl struct {
	sess session.Session
}

func New(sess session.Session) TxnDetailsImpl {
	return TxnDetailsImpl{sess: sess}
}

//GetTokenTxns Gets all the token transactions for a particular address,
//the options param allows filtering based on contract with options for pagination as well
func (txn TxnDetailsImpl) GetTokenTxns(chain constants.Chain, address string, options *TokenTxnsOpts) (
	resp types.TokenTxn, err error) {

	if !constants.TXN_GetTokenTxns.SupportsChain(chain) {
		return types.TokenTxn{}, constants.UnsupportedChainError
	}
	var urlVals url.Values
	if options != nil {
		urlVals = httpclient.QueryParamHelper(options.getMappableQueryParams())
	}
	path := strings.Replace(constants.TXN_GetTokenTxns.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	err = txn.sess.Client.Get(&resp, path, urlVals)

	return
}

//GetTxnDetails accepts a transaction signature or ID and returns transaction details if available.
func (txn TxnDetailsImpl) GetTxnDetails(chain constants.Chain, txnID string) (resp types.TxnByID, err error) {
	if !constants.TXN_GetTxnDetails.SupportsChain(chain) {
		return types.TxnByID{}, constants.UnsupportedChainError
	}

	path := strings.Replace(constants.TXN_GetTxnDetails.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":txnID", txnID, 1)
	err = txn.sess.Client.Get(&resp, path, nil)
	return
}

//GetBulkTxnDetails accepts a transaction signature or ID and returns transaction details if available.
func (txn TxnDetailsImpl) GetBulkTxnDetails(chain constants.Chain, hashList []string) (resp []types.TxnByID, err error) {
	if len(hashList) == 0 {
		return
	}
	if !constants.TXN_GetBulkTxnDetails.SupportsChain(chain) {
		return []types.TxnByID{}, constants.UnsupportedChainError
	}

	path := strings.Replace(constants.TXN_GetBulkTxnDetails.GetURI(), ":chain", chain.String(), 1)
	txnID := strings.Join(hashList, ",")
	path = strings.Replace(path, ":txnID", txnID, 1)
	err = txn.sess.Client.Get(&resp, path, nil)
	return
}

//GetTokenTxnsV2 Gets all the token transactions for a particular address along with pricing data,
//the options param allows filtering based on contract with options for pagination as well
func (txn TxnDetailsImpl) GetTokenTxnsV2(chain constants.Chain, address string, options *TokenTxnsOpts) (
	resp types.TokenTxnV2, err error) {
	if !constants.TXN_GetTokenTxnsV2.SupportsChain(chain) {
		return types.TokenTxnV2{}, constants.UnsupportedChainError
	}
	var urlVals url.Values
	if options != nil {
		urlVals = httpclient.QueryParamHelper(options.getMappableQueryParams())
	}

	path := strings.Replace(constants.TXN_GetTokenTxnsV2.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	err = txn.sess.Client.Get(&resp, path, urlVals)
	return
}

//GetRawTransactionsForAddress Gets the raw transaction details for a particular address.
//The Details include data passed during the transaction as well as detailed information about gas usage. It does still include token transfers if they were made
func (txn TxnDetailsImpl) GetRawTransactionsForAddress(chain constants.Chain, address string, options *TransactionDetailsOpts) (resp types.RawTransactionsResponseV1, err error) {
	if !constants.TXN_GetRawTransactionDetails.SupportsChain(chain) {
		return types.RawTransactionsResponseV1{}, constants.UnsupportedChainError
	}
	var urlVals url.Values
	if options == nil {
		options = &TransactionDetailsOpts{}
	}

	options.format = Raw
	urlVals = httpclient.QueryParamHelper(options.getMappableQueryParams())
	path := strings.Replace(constants.TXN_GetRawTransactionDetails.GetURI(), ":chain", chain.String(), 1)
	path = strings.Replace(path, ":address", address, 1)
	err = txn.sess.Client.Get(&resp, path, urlVals)
	return
}

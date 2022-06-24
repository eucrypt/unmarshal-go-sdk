package transaction_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net/http"
	"os"
	"testing"
)

func TestTxnDetailsImpl_GetTokenTxns(t *testing.T) {
	txnDetails := getTestTxnDetails()
	ast := assert.New(t)
	validAddr := "demo.eth"
	validChain := constants.ETH
	var validOptions = TokenTxnsOpts{
		Contract: "0xf8C3527CC04340b208C854E985240c02F7B7793f",
		PaginationOptions: PaginationOptions{
			Page:     1,
			PageSize: 10,
		},
	}

	t.Run("Evaluating Get Token Txns with no options", func(t *testing.T) {
		resp, err := txnDetails.GetTokenTxns(validChain, validAddr, nil)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
	})
	t.Run("Evaluating Get Token Txns with options", func(t *testing.T) {
		resp, err := txnDetails.GetTokenTxns(validChain, validAddr, &validOptions)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Len(resp.Transactions, 10, "Exactly 10 objects should be a part of the response")
	})
	t.Run("Evaluating Get Token Txns with options", func(t *testing.T) {
		validOptions.PageSize = 5

		resp, err := txnDetails.GetTokenTxns(validChain, validAddr, &validOptions)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Len(resp.Transactions, 5, "Exactly 5 objects should be a part of the response")
	})
	t.Run("Evaluating GetTokenTxns with incorrect chains", func(t *testing.T) {
		resp, err := txnDetails.GetTokenTxns(constants.HUOBI, "", nil)

		ast.Error(err, "There should be an error for an invalid call")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
		ast.Empty(resp, "The response should be empty")

	})
	t.Run("Evaluating GetTokenTxns with incorrect chains and valid options", func(t *testing.T) {
		resp, err := txnDetails.GetTokenTxns(constants.HUOBI, validAddr, &validOptions)

		ast.Error(err, "There should be an error for an invalid call")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
		ast.Empty(resp, "The response should be empty")

	})
	t.Run("Evaluating Get Token Txns on avalanche", func(t *testing.T) {
		avalancheAddr := "0x59DD353A69e2370ca98C1ff32587131b779a587a"
		avalancheChain := constants.AVALANCHE
		resp, err := txnDetails.GetTokenTxns(avalancheChain, avalancheAddr, &TokenTxnsOpts{
			PaginationOptions: PaginationOptions{
				Page:     2,
				PageSize: 10,
			},
		})

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Len(resp.Transactions, 10, "Exactly 10 objects should be a part of the response")
	})

}

func TestTxnDetailsImpl_GetTokenTxnsV2(t *testing.T) {

	txnDetails := getTestTxnDetails()
	ast := assert.New(t)
	validAddr := "demo.eth"
	validChain := constants.ETH
	var validOptions = TokenTxnsOpts{
		Contract: "0xf8C3527CC04340b208C854E985240c02F7B7793f",
		PaginationOptions: PaginationOptions{
			Page:     1,
			PageSize: 10,
		},
	}

	t.Run("Evaluating Get Token TxnsV2 with no options", func(t *testing.T) {
		resp, err := txnDetails.GetTokenTxnsV2(validChain, validAddr, nil)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
	})
	t.Run("Evaluating Get Token TxnsV2 with options", func(t *testing.T) {
		resp, err := txnDetails.GetTokenTxnsV2(validChain, validAddr, &validOptions)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Len(resp.Transactions, 10, "Exactly 10 objects should be a part of the response")
	})
	t.Run("Evaluating Get Token TxnsV2 with options", func(t *testing.T) {
		validOptions.PageSize = 5

		resp, err := txnDetails.GetTokenTxns(validChain, validAddr, &validOptions)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Len(resp.Transactions, 5, "Exactly 5 objects should be a part of the response")
	})
	t.Run("Evaluating GetTokenTxnsV2 with incorrect chains", func(t *testing.T) {
		resp, err := txnDetails.GetTokenTxnsV2(constants.HUOBI, "", nil)

		ast.Error(err, "There should be an error for an invalid call")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
		ast.Empty(resp, "The response should be empty")

	})
	t.Run("Evaluating GetTokenTxnsV2 with incorrect chains and valid options", func(t *testing.T) {
		resp, err := txnDetails.GetTokenTxnsV2(constants.HUOBI, validAddr, &validOptions)

		ast.Error(err, "There should be an error for an invalid call")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
		ast.Empty(resp, "The response should be empty")

	})

}

func TestTxnDetailsImpl_GetTxnDetails(t *testing.T) {
	txnDetails := getTestTxnDetails()
	ast := assert.New(t)
	validID := "0x4717c3987e8b1cdb831a45f99a0dbc1390ee53b704e8e171fa2154e598fdec1e"

	t.Run("Evaluating Get Transaction Details with valid data", func(t *testing.T) {
		validChain := constants.ETH

		resp, err := txnDetails.GetTxnDetails(validChain, validID)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
	})
	t.Run("Evaluating Get Transaction Details with valid data", func(t *testing.T) {
		validChain := constants.AVALANCHE
		validTxnId := "0xf1b56f250233d0cc1390e33e34223f6eadf4c935fe4c9ad4ded6aad416feec7f"
		resp, err := txnDetails.GetTxnDetails(validChain, validTxnId)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
	})

	t.Run("Evaluating Get Transaction Details with invalid data", func(t *testing.T) {
		resp, err := txnDetails.GetTxnDetails(constants.HUOBI, validID)

		ast.Error(err, "There should be an error for an invalid call")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
		ast.Empty(resp, "The response should be empty")
	})
}

func getTestTxnDetails() TxnDetailsImpl {
	httpClient := httpclient.NewHttpJSONClient(constants.Environment.GetEndpoint("prod"))
	authKey := os.Getenv("API_KEY")
	httpClient.DefaultQuery = map[string]string{"auth_key": authKey}
	txnDetails := New(session.Session{Config: struct {
		AuthKey     string
		HttpClient  *http.Client
		Environment constants.Environment
	}{AuthKey: authKey, HttpClient: nil, Environment: constants.Prod}, Client: httpClient})
	return txnDetails
}

func TestTxnDetailsImpl_GetRawTransactionsForAddress(t *testing.T) {
	txnDetails := getTestTxnDetails()
	ast := assert.New(t)
	validAddr := "0x5a666c7d92E5fA7Edcb6390E4efD6d0CDd69cF37"
	validChain := constants.ETH

	t.Run("Valid params should make a call with no errors", func(t *testing.T) {

		resp, err := txnDetails.GetRawTransactionsForAddress(validChain, validAddr, &TransactionDetailsOpts{
			PaginationOptions: PaginationOptions{
				Page:     1,
				PageSize: 2,
			},
		})

		ast.NoError(err, "there should be no error for a valid call")
		ast.NotEmpty(resp.Transactions, "this call should have valid transactions present")
		ast.Len(resp.Transactions, 2, "There should be only two transactions for this call.")
	})
	t.Run("missing options should fill defaults", func(t *testing.T) {

		resp, err := txnDetails.GetRawTransactionsForAddress(validChain, validAddr, nil)

		ast.NoError(err, "there should be no error for a valid call")
		ast.Empty(resp.Transactions, "Should return empty transaction response for invalid page number")
	})

	t.Run("Invalid chain should cause error", func(t *testing.T) {

		resp, err := txnDetails.GetRawTransactionsForAddress(constants.HUOBI, validAddr, &TransactionDetailsOpts{
			PaginationOptions: PaginationOptions{
				Page:     1,
				PageSize: 2,
			},
		})

		ast.EqualError(err, constants.UnsupportedChainError.Error(), "Call should err for an unsupported chain")
		ast.Empty(resp, "call should return an empty object")
	})

}

//@dev test Pending until gateway pr is merged
func TestTxnDetailsImpl_GetTransactionsByCursor(t *testing.T) {
	txnDetails := getTestTxnDetails()
	ast := assert.New(t)
	validAddr := "0x5a666c7d92E5fA7Edcb6390E4efD6d0CDd69cF37"
	validChain := constants.ETH
	t.Run("Valid params should make a call with no errors", func(t *testing.T) {

		resp, err := txnDetails.GetTransactionsByCursor(validChain, &types.AddressTxCursor{
			ContractAddress: validAddr,
			BlockNumber:     big.NewInt(12141206),
			SeqID:           0,
		}, nil, 10)

		ast.NoError(err, "there should be no error for a valid call")
		ast.NotEmpty(resp.Transactions, "this call should have valid transactions present")
		ast.Len(resp.Transactions, 10, "There should be 10 transactions for this call.")
		ast.NotNil(resp.EndCursor, "there should be an end cursor in the call as the start cursor was passed")
	})

	t.Run("wrong chain should error", func(t *testing.T) {
		invalidChain := constants.HUOBI
		resp, err := txnDetails.GetTransactionsByCursor(invalidChain, &types.AddressTxCursor{
			ContractAddress: validAddr,
			BlockNumber:     big.NewInt(12141206),
			SeqID:           0,
		}, nil, 10)

		ast.EqualError(err, constants.UnsupportedChainError.Error(), "an unsupported chain should present an error")
		ast.Empty(resp, "erroring call should return and empty object")
	})

}

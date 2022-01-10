package transaction_details

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestTxnDetailsImpl_GetTokenTxns(t *testing.T) {
	txnDetails := getTextTxnDetails()
	ast := assert.New(t)
	validAddr := "demo.eth"
	validChain := constants.ETH
	var validOptions = TokenTxnsOpts{
		Contract: "0xf8C3527CC04340b208C854E985240c02F7B7793f",
		Page:     1,
		PageSize: 10,
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

}

func TestTxnDetailsImpl_GetTokenTxnsV2(t *testing.T) {

	txnDetails := getTextTxnDetails()
	ast := assert.New(t)
	validAddr := "demo.eth"
	validChain := constants.ETH
	var validOptions = TokenTxnsOpts{
		Contract: "0xf8C3527CC04340b208C854E985240c02F7B7793f",
		Page:     1,
		PageSize: 10,
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
	txnDetails := getTextTxnDetails()
	ast := assert.New(t)
	validID := "0x4717c3987e8b1cdb831a45f99a0dbc1390ee53b704e8e171fa2154e598fdec1e"

	t.Run("Evaluating Get Transaction Details with valid data", func(t *testing.T) {
		validChain := constants.ETH

		resp, err := txnDetails.GetTxnDetails(validChain, validID)

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

func getTextTxnDetails() TxnDetailsImpl {
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

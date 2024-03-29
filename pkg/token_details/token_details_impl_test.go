package token_details

//@dev all tests require an "API_KEY" env to be specified

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestTokenStoreV1_GetTokenDetailsWithContract(t *testing.T) {
	ts := getTestTokenStore()
	ValidContract := "0x5a666c7d92e5fa7edcb6390e4efd6d0cdd69cf37"
	ast := assert.New(t)
	t.Run("Evaluating get Token Details with valid data", func(t *testing.T) {
		resp, err := ts.GetTokenDetailsByContract(ValidContract, nil)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
	})

	t.Run("Evaluating get Token Details with valid data and specific chain", func(t *testing.T) {
		chain := constants.BSC
		opts := TokenDetailsOptions{Chain: chain}
		resp, err := ts.GetTokenDetailsByContract("0x111111111117dC0aa78b770fA6A738034120C302", &opts)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Equal(resp.Blockchain, chain.String())
		//fmt.Printf("%#v", resp)
	})

	t.Run("Evaluating get Token Details with valid data and specific chain", func(t *testing.T) {
		chain := constants.ETH
		opts := TokenDetailsOptions{Chain: chain}
		resp, err := ts.GetTokenDetailsByContract("0x111111111117dC0aa78b770fA6A738034120C302", &opts)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
		ast.Equal(resp.Blockchain, chain.String())
		//fmt.Printf("%#v", resp)
	})

	t.Run("Evaluating get Token Details with valid data", func(t *testing.T) {
		resp, err := ts.GetTokenDetailsByContract(ValidContract, nil)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
	})

	t.Run("Evaluating get Token Details with invalid data", func(t *testing.T) {
		resp, _ := ts.GetTokenDetailsByContract("", nil)
		ast.Empty(resp, "The response should be empty for invalid data")
	})

}

func getTestTokenStore() TokenStoreV1 {
	httpClient := httpclient.NewHttpJSONClient(constants.Environment.GetEndpoint("prod"))
	authKey := os.Getenv("API_KEY")
	httpClient.DefaultQuery = map[string]string{"auth_key": authKey}
	PsV1 := New(session.Session{Config: struct {
		AuthKey     string
		HttpClient  *http.Client
		Environment constants.Environment
	}{AuthKey: authKey, HttpClient: nil, Environment: constants.Prod}, Client: httpClient})
	return PsV1
}

func TestTokenStoreV1_GetTokenList(t *testing.T) {
	ts := getTestTokenStore()
	pageNumber := 1

	t.Run("Get token List", func(t *testing.T) {
		pageSize := 10
		resp, err := ts.GetTokenList(pageNumber, pageSize)
		assert.NoError(t, err, "There should be no error for a valid call")
		assert.NotEmpty(t, resp, "The response should not be empty")
		assert.Len(t, resp.Data, 10, "Exactly 10 objects should be a part of the response")
	})
	t.Run("Get token List", func(t *testing.T) {
		pageSize := 5
		resp, err := ts.GetTokenList(pageNumber, pageSize)
		assert.NoError(t, err, "There should be no error for a valid call")
		assert.NotEmpty(t, resp, "The response should not be empty")
		assert.Len(t, resp.Data, 5, "Exactly 5 objects should be a part of the response")
	})
}

func TestTokenStoreV1_GetTokenWithSymbol(t *testing.T) {
	ts := getTestTokenStore()
	validSymbol := "marsh"
	ast := assert.New(t)
	t.Run("Evaluating GetTokenDetailsBySymbol with valid data", func(t *testing.T) {
		resp, err := ts.GetTokenDetailsBySymbol(validSymbol)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")
	})

	t.Run("Evaluating GetTokenDetailsBySymbol with invalid data", func(t *testing.T) {
		resp, _ := ts.GetTokenDetailsBySymbol("")
		ast.Empty(resp, "The response should be empty for invalid data")
	})

}

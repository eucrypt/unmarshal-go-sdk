package assets

import (
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	httpclient "github.com/eucrypt/unmarshal-go-sdk/pkg/http"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestV1Store_GetAssets(t *testing.T) {
	testAssetsStore := getTestAssetsStore()
	ast := assert.New(t)
	validAddr := "demo.eth"
	chain := constants.ETH
	t.Run("Evaluate TokenAssets", func(t *testing.T) {
		resp, err := testAssetsStore.GetTokenAssets(chain, validAddr)

		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = testAssetsStore.GetTokenAssets(chain, "invalidAddr")
		ast.Empty(resp, "should have an empty response for an invalid call")
	})
}

func TestV1Store_GetProfitAndLoss(t *testing.T) {
	assetsStore := getTestAssetsStore()
	ast := assert.New(t)
	validAddr := "demo.eth"
	validContract := "0xdac17f958d2ee523a2206206994597c13d831ec7"
	chain := constants.ETH
	t.Run("Evaluate Profit and Loss", func(t *testing.T) {
		resp, err := assetsStore.GetProfitAndLoss(chain, validAddr, validContract)
		ast.NoError(err, "There should be no error for a valid call")
		ast.NotEmpty(resp, "The response should not be empty")

		resp, _ = assetsStore.GetProfitAndLoss(chain, "invalid", validContract)
		ast.Empty(resp, "should have an empty response for an invalid call")
		resp, _ = assetsStore.GetProfitAndLoss(chain, validAddr, "invalid")
		ast.Empty(resp, "should have an empty response for an invalid call")

		//@dev The chain below is currently unsupported. Test will be deprecated if the chain is ever supported
		_, err = assetsStore.GetProfitAndLoss(constants.HUOBI, "demo.eth", "")
		ast.Equal(constants.UnsupportedChainError, err, "Call should result in an unsupported chain error")
	})
}

func getTestAssetsStore() V1Store {
	httpClient := httpclient.NewHttpJSONClient(constants.Environment.GetEndpoint("prod"))
	authKey := os.Getenv("API_KEY")
	httpClient.DefaultQuery = map[string]string{"auth_key": authKey}
	assetsStore := New(session.Session{Config: struct {
		AuthKey     string
		HttpClient  *http.Client
		Environment constants.Environment
	}{AuthKey: authKey, HttpClient: nil, Environment: constants.Prod}, Client: httpClient})
	return assetsStore
}
